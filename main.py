'''discord-stock-ticker'''
from os import getenv
import logging

import asyncio
import discord
from redis import Redis, exceptions

from utils.yahoo import get_stock_price_async
from utils.coin_gecko import get_crypto_price_async

CURRENCY = 'usd'
NAME_CHANGE_DELAY = 3600


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
                    getenv('FREQUENCY', 60)
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
                    getenv('FREQUENCY', 60)
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

                    old_price = price
                    logging.info('name updated')
                except discord.HTTPException as e:
                    logging.warning(f'updating name failed: {e.status}: {e.text}')

            else:
                logging.info('no price change')

            # Only update every hour
            logging.info(f'stock name sleeping for {NAME_CHANGE_DELAY}s')
            await asyncio.sleep(NAME_CHANGE_DELAY)
            logging.info('stock name sleep ended')


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

            logging.info('stock activity update started')
            
            # Grab the current price data w/ day difference
            data = await get_stock_price_async(ticker)
            price_data = data.get('quoteSummary', {}).get('result', []).pop().get('price', {})
            price = price_data.get('regularMarketPrice', {}).get('raw', 0.00)

            # If after hours, get change
            if price_data.get('postMarketChange'):
                raw_diff = price_data.get('postMarketChange', {}).get('raw', 0.00)
                diff = round(raw_diff, 2)
                if diff > 0:
                    diff = '+' + str(diff)

                activity_content = f'After Hours: {diff}'
                logging.info(f'stock activity after hours price retrived: {activity_content}')
            else:
                raw_diff = price_data.get('regularMarketChange', {}).get('raw', 0.00)
                diff = round(raw_diff, 2)
                if diff > 0:
                    diff = '+' + str(diff)

                activity_content = f'${price} / {diff}'
                logging.info(f'stock activity price retrived: {activity_content}')

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
                    
                    # Check what price we are displaying
                    if price_data.get('postMarketChange'):
                        activity_content_header = 'After Hours'
                    else:
                        activity_content_header = 'Day Diff'
                    
                    activity_content = f'{activity_content_header}: {diff}'
                    

                # Change activity
                try:
                    await self.change_presence(
                        activity=discord.Activity(
                            type=discord.ActivityType.watching,
                            name=activity_content
                        )
                    )

                    old_price = price
                    logging.info('activity updated')

                except discord.InvalidArgument as e:
                    logging.error(f'updating activity failed: {e.status}: {e.text}')


            else:
                logging.info('no price change')

            # Only update every min
            logging.info(f'stock activity sleeping for {frequency}s')
            await asyncio.sleep(int(frequency))
            logging.info('stock activity sleep ended')
    

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
            data = await get_crypto_price_async(crypto_name)
            price = data.get('market_data', {}).get('current_price', {}).get(CURRENCY, 0.0)
            logging.info(f'crypto name price retrived {price}')

            # Only update on price change
            if old_price != price:

                try:
                    await self.user.edit(
                        username=f'{ticker} - ${price}'
                    )

                    old_price = price
                    logging.info('crypto name updated')
                except discord.HTTPException as e:
                    logging.warning(f'updating name failed: {e.status}: {e.text}')

            else:
                logging.info('no price change')

            # Only update every hour
            logging.info(f'crypto name sleeping for {NAME_CHANGE_DELAY}s')
            await asyncio.sleep(NAME_CHANGE_DELAY)
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
            data = await get_crypto_price_async(crypto_name)
            price = data.get('market_data', {}).get('current_price', {}).get(CURRENCY, 0.0)
            change = data.get('market_data', {}).get('price_change_24h', 0)
            change_header = ''
            if change > 0:
                change_header = '+'

            logging.info(f'crypto activity price retrived {price}')

            activity_content = f'${price} / {change_header}{change}'

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
                    activity_content = f'24hr Diff: {change_header}{change}'

                # Change activity
                try:
                    await self.change_presence(
                        activity=discord.Activity(
                            type=discord.ActivityType.watching,
                            name=activity_content
                        )
                    )

                    old_price = price
                    logging.info('crypto activity updated')
                except discord.InvalidArgument as e:
                    logging.error(f'updating activity failed: {e.status}: {e.text}')

            else:
                logging.info('no price change')

            # Only update every min
            logging.info(f'crypto sleeping for {frequency}s')
            await asyncio.sleep(int(frequency))
            logging.info('crypto activity sleep ended')


if __name__ == "__main__":

    logging.basicConfig(
        filename=getenv('LOG_FILE'),
        level=logging.INFO,
        datefmt='%Y-%m-%d %H:%M:%S',
        format='%(asctime)s %(levelname)-8s %(message)s',
    )

    token = getenv('DISCORD_BOT_TOKEN')
    if not token:
        print('DISCORD_BOT_TOKEN not set!')

    client = Ticker()
    client.run(token)
