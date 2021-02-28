'''Yahoo Finance Helpers'''
import aiohttp
import asyncio


COINGECKO_URL = 'https://api.coingecko.com/api/v3/'
HEADERS = {
    'User-Agent': 'Mozilla/5.0',
    'accept': 'application/json'
}


async def get_crypto_price_async(ticker: str) -> dict:
    '''
    Get a live stock price from YF API
    '''

    try:
        timeout = aiohttp.ClientTimeout(total=5)

        async with aiohttp.ClientSession(timeout=timeout) as session:

            async with session.get(COINGECKO_URL + f'coins/{ticker}') as response:

                assert 200 == response.status, response.reason
                return await response.json()
    except asyncio.TimeoutError as e:
        print(f'Unable to get coingecko prices: {e}')
        return {}
