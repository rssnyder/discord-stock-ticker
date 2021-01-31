import discord
from os import getenv
import yfinance as yf
import asyncio


class Ticker(discord.Client):


    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)

        self.bg_task = self.loop.create_task(self.update_price())


    async def on_ready(self):

        print('Logged in as', self.user.name)


    async def update_price(self):

        await self.wait_until_ready()

        while not self.is_closed():            
            
            print(f'Updating price for {getenv("TICKER")}')

            ticker = yf.Ticker(getenv('TICKER'))
            await self.user.edit(
                username=f'{getenv("TICKER")} - ${ticker.info["regularMarketPrice"]}'
            )

            diff = ticker.info["regularMarketPrice"] - ticker.info['open']
            if diff > 0:
                diff = '+' + str(diff)

            await self.change_presence(
                activity=discord.Activity(
                    type=discord.ActivityType.watching,
                    name=diff
                )
            )

            await asyncio.sleep(600)


client = Ticker()
client.run(getenv('DISCORD_BOT_TOKEN'))
