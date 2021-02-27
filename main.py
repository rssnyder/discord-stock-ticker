'''discord-stock-ticker'''
from os import getenv
from datetime import datetime
from random import choice
import logging
import asyncio
import discord
from redis import Redis, exceptions

from utils.yahoo import get_stock_price_async
from utils.coin_gecko import get_crypto_price_async

CURRENCY = 'usd'
WEEKEND = [
    5,
    6
]
ALERTS = [
    'discord.gg/CQqnCYEtG7',
    'markets be closed',
    'gme to the moon',
    'what about second breakfast',
    'subscribe for real time prices!'
]


class Ticker(discord.Client):
    '''
    Discord client for watching stock/crypto prices
    '''

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)

        ticker = getenv("TICKER")
        crypto_name = getenv('CRYPTO_NAME')
        stock_name = getenv("STOCK_NAME", ticker)


        # Check that at least a ticker is set
        if not ticker:
            logging.error('TICKER not set!')
            return

        # Use different updates based on security type
        if crypto_name:
            logging.info('crypo ticker')
            self.sm_task = self.loop.create_task(
                self.crypto_update_name(
                    ticker.upper(),
                    crypto_name
                )
            )
            self.bg_task = self.loop.create_task(
                self.crypto_update_activity(
                    ticker.upper(),
                    crypto_name,
                    getenv('SET_NICKNAME'),
                    getenv('FREQUENCY')
                )
            )
        else:
            logging.info('stock ticker')
            self.sm_task = self.loop.create_task(
                self.stock_update_name(
                    ticker.upper(),
                    stock_name.upper()
                )
            )
            self.bg_task = self.loop.create_task(
                self.stock_update_activity(
                    ticker.upper(),
                    stock_name.upper(),
                    getenv('SET_NICKNAME'),
                    getenv('FREQUENCY')
                )
            )


    async def on_ready(self):
        '''
        Log that we have successfully connected
        '''

        logging.info('logged in')

        # Use redis to store stats
        r = Redis()

        # We want to know some stats
        servers = [x.name for x in list(self.guilds)]

        try:
            for server in servers:
                r.incr(server)
        except exceptions.ConnectionError:
            logging.info('No redis server found, not storing stats')

        logging.info('servers: ' + str(servers))


    async def stock_update_name(self, ticker: str, name: str):
        '''
        Update the bot name based on stock price
        ticker = stock symbol
        name = override for symbol as shown on bot
        '''

        old_price = ''

        await self.wait_until_ready()
        logging.info(f'stock name update ready: {name}')

        # Loop as long as the bot is running
        while not self.is_closed():

            # Dont bother updating if markets closed
            if (datetime.now().hour >= 17) or (datetime.now().hour < 8) or (datetime.today().weekday() in WEEKEND):
                logging.info('markets are closed')
                await asyncio.sleep(3600)
                continue

            logging.info('stock name update started')
            
            # Grab the current price data
            data = await get_stock_price_async(ticker)
            price_data = data.get('quoteSummary', {}).get('result', []).pop().get('price', {})
            price = price_data.get('regularMarketPrice', {}).get('raw', 0.00)
            logging.info(f'stock name price retrived {price}')

            # Only update on price change
            if old_price != price:

                try:
                    await self.user.edit(
                        username=f'{name} - ${price}'
                    )
                except discord.HTTPException as e:
                    logging.error(f'updating name failed: {e.status}: {e.text}')

                old_price = price
                logging.info('name updated')

            else:
                logging.info('no price change')

            # Only update every hour
            await asyncio.sleep(3600)

            logging.info('name sleep ended')


    async def stock_update_activity(self, ticker: str, name: str, change_nick: bool = False, frequency: int = 60):
        '''
        Update the bot activity based on stock price
        ticker = stock symbol
        name = override for symbol as shown on bot
        change_nick = flag for changing nickname
        frequency = how often to update in seconds
        '''

        # Get config
        old_price = ''

        await self.wait_until_ready()
        logging.info(f'stock activity update ready: {name}')

        # Loop as long as the bot is running
        while not self.is_closed():

            # If markets are closed, utilize activity for other messages
            hour = datetime.now().hour
            day = datetime.today().weekday()
            if (hour >= 17) or (hour < 8) or (day in WEEKEND):

                logging.info('markets are closed')

                try:
                    await self.change_presence(
                        activity=discord.Activity(
                            type=discord.ActivityType.watching,
                            name=choice(ALERTS)
                        )
                    )
                except discord.InvalidArgument as e:
                    logging.error(f'updating activity failed: {e.status}: {e.text}')

                await asyncio.sleep(600)
                continue

            logging.info('stock activity update started')
            
            # Grab the current price data w/ day difference
            data = await get_stock_price_async(ticker)
            price_data = data.get('quoteSummary', {}).get('result', []).pop().get('price', {})
            price = price_data.get('regularMarketPrice', {}).get('raw', 0.00)
            raw_diff = price - data.get('regularMarketPreviousClose', {}).get('raw', 0.00)
            diff = round(raw_diff, 2)
            if diff > 0:
                diff = '+' + str(diff)

            logging.info(f'stock activity price retrived {price}')

            activity_content = f'${price} / {diff}'

            # Only update on price change
            if old_price != price:

                # Change name via nickname if set
                if change_nick:

                    for server in self.guilds:

                        try:
                            await server.me.edit(
                                nick=f'{name} - ${price}'
                            )
                        except discord.HTTPException as e:
                            logging.error(f'updating nick failed: {e.status}: {e.text}')
                        except discord.Forbidden as f:
                            logging.error(f'lacking perms for chaning nick: {f.status}: {f.text}')

                        logging.info(f'stock updated nick in {server.name}')
                    
                    # Do not include price in activity now
                    activity_content = f'Day Change: ${diff}'
                    

                # Change activity
                try:
                    await self.change_presence(
                        activity=discord.Activity(
                            type=discord.ActivityType.watching,
                            name=activity_content
                        )
                    )
                except discord.InvalidArgument as e:
                    logging.error(f'updating activity failed: {e.status}: {e.text}')

                logging.info('activity updated')

                old_price = price

            else:
                logging.info('no price change')

            # Only update every min
            await asyncio.sleep(frequency)
            logging.info('activity sleep ended')
    

    async def crypto_update_name(self, ticker: str, crypto_name: str):
        '''
        Update the bot name based on crypto price
        ticker = symbol to display on bot
        name = crypto name for CG api
        '''

        old_price = ''

        await self.wait_until_ready()
        logging.info(f'crypto name update ready: {crypto_name}')

        # Loop as long as the bot is running
        while not self.is_closed():

            logging.info('crypto name started')

            # Grab the current price data
            data = get_crypto_price_async(name)
            price = data.get('market_data', {}).get('current_price', {}).get(CURRENCY, 0.0)
            logging.info(f'crypto name price retrived {price}')

            # Only update on price change
            if old_price != price:

                try:
                    await self.user.edit(
                        username=f'{ticker} - ${price}'
                    )
                except discord.HTTPException as e:
                    logging.error(f'updating name failed: {e.status}: {e.text}')
                
                old_price = price
                logging.info('crypto name updated')

            else:
                logging.info('no price change')

            # Only update every hour
            await asyncio.sleep(3600)
            logging.info('crypto name sleep ended')
    

    async def crypto_update_activity(self, ticker: str, crypto_name: str, change_nick: bool = False, frequency: int = 60):
        '''
        Update the bot activity based on crypto price
        ticker = symbol to display on bot
        name = crypto name for CG api
        change_nick = flag for changing nickname
        frequency = how often to update in seconds
        '''

        old_price = ''

        await self.wait_until_ready()
        logging.info(f'crypto activity update ready: {crypto_name}')

        # Loop as long as the bot is running
        while not self.is_closed():

            logging.info('crypto activity started')       

            # Grab the current price data
            data = get_crypto_price_async(name)
            price = data.get('market_data', {}).get('current_price', {}).get(CURRENCY, 0.0)
            change = data.get('price_change_24h', 0)

            logging.info(f'crypto activity price retrived {price}')

            activity_content = f'${price} / {change}'

            # Only update on price change
            if old_price != price:

                # Change name via nickname if set
                if change_nick:
                    
                    for server in self.guilds:

                        try:
                            await server.me.edit(
                                nick=f'{ticker} - ${price}'
                            )
                        except discord.HTTPException as e:
                            logging.error(f'updating nick failed: {e.status}: {e.text}')
                        except discord.Forbidden as f:
                            logging.error(f'lacking perms for chaning nick: {f.status}: {f.text}')

                        logging.info(f'updated nick in {server.name}')
                    
                    # Use activity for other fun stuff
                    activity_content = 'price changes'

                # Change activity
                try:
                    await self.change_presence(
                        activity=discord.Activity(
                            type=discord.ActivityType.watching,
                            name=activity_content
                        )
                    )
                except discord.InvalidArgument as e:
                    logging.error(f'updating activity failed: {e.status}: {e.text}')

                logging.info('crypto activity updated')

                old_price = price

            else:
                logging.info('no price change')

            # Only update every min
            await asyncio.sleep(frequency)
            logging.info('crypto activity sleep ended')


if __name__ == "__main__":

    logging.basicConfig(
        filename='discord-stock-ticker.log',
        level=logging.INFO,
        datefmt='%Y-%m-%d %H:%M:%S',
        format='%(asctime)s %(levelname)-8s %(message)s',
    )

    token = getenv('DISCORD_BOT_TOKEN')
    if not token:
        print('DISCORD_BOT_TOKEN not set!')

    client = Ticker()
    client.run(token)
