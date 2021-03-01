'''Yahoo Finance Helpers'''
import aiohttp
import asyncio


YAHOO_URL = 'https://query1.finance.yahoo.com/v10/finance/'
HEADERS = {
    'User-Agent': 'Mozilla/5.0'
}


async def get_stock_price_async(ticker: str) -> dict:
    '''
    Get a live stock price from YF API
    '''

    try:
        timeout = aiohttp.ClientTimeout(total=5)

        async with aiohttp.ClientSession(timeout=timeout) as session:

            async with session.get(YAHOO_URL + f'quoteSummary/{ticker}?modules=price') as response:

                assert 200 == response.status, response.reason
                return await response.json()
    except asyncio.TimeoutError as e:
        print(f'Unable to get yahoo prices: {e}')
        return {}
    except aiohttp.client_exceptions.ClientConnectorError:
        print(f'Unable to get yahoo prices: {e}')
        return {}
