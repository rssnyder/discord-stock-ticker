'''discord-stock-ticker'''
from os import getenv
from sys import stdout
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

        if not getenv("TICKER"):
            logging.error('TICKET not set!')
            return

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

        name = self.user.name.split(' ')
        logging.info(f'{name[0]}: logged in')


    async def stock_update_name(self):
        '''
        Update the bot name based on stock price
        '''

        await self.wait_until_ready()

        while not self.is_closed():            
            
            ticker = getenv("TICKER")

            data = yf.Ticker(ticker)

            await self.user.edit(
                username=f'{ticker} - ${data.info["bid"]}'
            )
            logging.info(f'{ticker}: name update')

            await asyncio.sleep(3605)
    

    async def stock_update_activity(self):
        '''
        Update the bot activity based on stock price
        '''

        await self.wait_until_ready()

        while not self.is_closed():            
            
            ticker = getenv("TICKER")

            data = yf.Ticker(ticker)

            diff = data.info['bid'] - data.info['open']
            diff = round(diff, 2)
            if diff > 0:
                diff = '+' + str(diff)

            await self.change_presence(
                activity=discord.Activity(
                    type=discord.ActivityType.watching,
                    name=f'${data.info["bid"]} / {diff}'
                )
            )
            logging.info(f'{ticker}: activity update')

            await asyncio.sleep(60)
    

    async def crypto_update_name(self, gapi: CoinGeckoAPI):
        '''
        Update the bot name based on crypto price
        '''

        await self.wait_until_ready()

        while not self.is_closed():            
            
            ticker = getenv("TICKER")

            name = getenv('CRYPTO_NAME')
            data = gapi.get_price(ids=name, vs_currencies=CURRENCY)
            price = data.get(name, {}).get(CURRENCY)

            await self.user.edit(
                username=f'{ticker} - ${price}'
            )
            logging.info(f'{name}: name update')

            await asyncio.sleep(3605)
    

    async def crypto_update_activity(self, gapi: CoinGeckoAPI):
        '''
        Update the bot activity based on crypto price
        '''

        await self.wait_until_ready()

        while not self.is_closed():            

            name = getenv('CRYPTO_NAME')
            data = gapi.get_price(ids=name, vs_currencies=CURRENCY)
            price = data.get(name, {}).get(CURRENCY)

            await self.change_presence(
                activity=discord.Activity(
                    type=discord.ActivityType.watching,
                    name=f'${price}'
                )
            )
            logging.info(f'{name}: activity update')

            await asyncio.sleep(60)


if __name__ == "__main__":

    logging.basicConfig(
        filename='discord-stock-ticker.log',
        level=logging.INFO,
        datefmt='%Y-%m-%d %H:%M:%S',
        format='%(asctime)s %(levelname)-8s %(message)s',
    )

    client = Ticker()
    client.run(getenv('DISCORD_BOT_TOKEN'))
