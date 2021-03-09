'''Coin Gecko Helpers'''
from logging import error

from requests import get


COINGECKO_URL = 'https://api.coingecko.com/api/v3/'
HEADERS = {
    'User-Agent': 'Mozilla/5.0',
    'accept': 'application/json'
}


def get_crypto_price(ticker: str) -> dict:
    '''
    Get a live crypto price from CoinGecko API
    '''

    resp = get(
        COINGECKO_URL + f'coins/{ticker}',
        headers=HEADERS
    )

    try:
        resp.raise_for_status()
    except:
        error('Error reaching yahoo')
        return {}

    return resp.json()
