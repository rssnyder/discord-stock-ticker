'''discord-stock-ticker'''
from os import getenv
import time
import logging
import asyncio
import discord
import yfinance as yf
from pycoingecko import CoinGeckoAPI


CURRENCY = 'usd'


class Ticker(discord.Client):
    '''
    Discord client for watching stock/crypto prices
    '''

    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)

        # Check that at least a ticker is set
        if not getenv("TICKER"):
            logging.error('TICKET not set!')
            return

        # Use different updates based on security type
        if getenv('CRYPTO_NAME'):
            logging.info('crypo ticket')
            api = CoinGeckoAPI()
            self.sm_task = self.loop.create_task(self.crypto_update_name(api))
            self.bg_task = self.loop.create_task(self.crypto_update_activity(api))
        else:
            logging.info('stock ticket')
            self.sm_task = self.loop.create_task(self.stock_update_name())
            self.bg_task = self.loop.create_task(self.stock_update_activity())


    async def on_ready(self):
        '''
        Log that we have successfully connected
        '''

        # Connect to discord
        name = self.user.name.split(' ')
        logging.info('logged in')

        # We want to know where we are running
        servers = [x.name for x in list(self.guilds)]
        logging.info('installed: ' + servers)


    async def stock_update_name(self):
        '''
        Update the bot name based on stock price
        '''

        ticker = getenv("TICKER")
        old_price = ''

        await self.wait_until_ready()
        logging.info('name ready')

        while not self.is_closed():

            start = time.time()
            logging.info('name started')
            
            # Grab the current price data
            data = yf.Ticker(ticker)
            price = data.info['bid']
            logging.info(f'name price retrived {price}')

            # Only update on price change
            if old_price != price:
                await self.user.edit(
                    username=f'{ticker} - ${price}'
                )
                logging.info('name updated')
                old_price = price
            else:
                logging.info('no price change')

            # Only update every hour
            end = time.time()
            await asyncio.sleep(3600 - int(end - start))
            logging.info('name sleep ended')
    

    async def stock_update_activity(self):
        '''
        Update the bot activity based on stock price
        '''

        ticker = getenv("TICKER")
        old_price = ''

        await self.wait_until_ready()
        logging.info('activity ready')

        while not self.is_closed():

            start = time.time()
            logging.info('activity started')
            
            # Grab the current price data w/ day difference
            data = yf.Ticker(ticker)
            price = data.info['bid']
            diff = price - data.info['open']
            diff = round(diff, 2)
            if diff > 0:
                diff = '+' + str(diff)
            logging.info(f'activity price retrived {price}')

            # Only update on price change
            if old_price != price:
                await self.change_presence(
                    activity=discord.Activity(
                        type=discord.ActivityType.watching,
                        name=f'${price} / {diff}'
                    )
                )
                logging.info('activity updated')
                old_price = price
            else:
                logging.info('no price change')

            # Only update every min
            end = time.time()
            await asyncio.sleep(60 - int(end - start))
            logging.info('activity sleep ended')
    

    async def crypto_update_name(self, gapi: CoinGeckoAPI):
        '''
        Update the bot name based on crypto price
        '''

        name = getenv('CRYPTO_NAME')
        ticker = getenv("TICKER")
        old_price = ''

        await self.wait_until_ready()
        logging.info('name ready')

        while not self.is_closed():

            start = time.time()
            logging.info('name started')

            # Grab the current price data
            data = gapi.get_price(ids=name, vs_currencies=CURRENCY)
            price = data.get(name, {}).get(CURRENCY)
            logging.info(f'name price retrived {price}')

            # Only update on price change
            if old_price != price:
                await self.user.edit(
                    username=f'{ticker} - ${price}'
                )
                logging.info('name updated')
                old_price = price
            else:
                logging.info('no price change')

            # Only update every hour
            end = time.time()
            await asyncio.sleep(3600 - int(end - start))
            logging.info('name sleep ended')
    

    async def crypto_update_activity(self, gapi: CoinGeckoAPI):
        '''
        Update the bot activity based on crypto price
        '''

        name = getenv('CRYPTO_NAME')
        old_price = ''

        await self.wait_until_ready()
        logging.info('activity ready')

        while not self.is_closed():

            start = time.time()
            logging.info('activity started')       

            # Grab the current price data
            data = gapi.get_price(ids=name, vs_currencies=CURRENCY)
            price = data.get(name, {}).get(CURRENCY)
            logging.info(f'activity price retrived {price}')

            # Only update on price change
            if old_price != price:
                await self.change_presence(
                    activity=discord.Activity(
                        type=discord.ActivityType.watching,
                        name=f'${price}'
                    )
                )
                logging.info('activity updated')
                old_price = price
            else:
                logging.info('no price change')

            # Only update every min
            end = time.time()
            await asyncio.sleep(60 - int(end - start))
            logging.info('activity sleep ended')

if __name__ == "__main__":

    logging.basicConfig(
        filename='discord-stock-ticker.log',
        level=logging.INFO,
        datefmt='%Y-%m-%d %H:%M:%S',
        format='%(asctime)s %(levelname)-8s %(message)s',
    )

    client = Ticker()
    client.run(getenv('DISCORD_BOT_TOKEN'))
