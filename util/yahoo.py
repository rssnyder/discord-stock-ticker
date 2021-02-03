'''Yahoo Finance Helpers'''
from requests import get
import aiohttp


YAHOO_URL = 'https://query1.finance.yahoo.com/v8/finance/'
HEADERS = {
    'User-Agent': 'Mozilla/5.0'
}


def get_stock_price(ticker: str) -> dict:
    '''
    Get a live stock price from YF API
    '''

    resp = get(YAHOO_URL + f'chart/{ticker}', headers=HEADERS)

    return resp.json()


async def get_stock_price_async(ticker: str):
    '''
    Get a live stock price from YF API
    '''

    async with aiohttp.ClientSession() as session:

        async with session.get(YAHOO_URL + f'chart/{ticker}') as response:

            assert 200 == response.status, response.reason
            return await response.json()