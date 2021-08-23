#!/usr/bin/python3
from sqlite3 import connect
import logging
from json import dumps
from time import sleep
import sys
import os

from requests import get, post

os.environ['http_proxy'] = ''
os.environ['https_proxy'] = ''
os.environ['HTTP_PROXY'] = ''
os.environ['HTTPS_PROXY'] = ''

if len(sys.argv) < 3:
    print('pass db file and url')
    sys.exit(1)

db_client = connect(sys.argv[1])

get_cur = db_client.cursor()
get_cur.execute(
    'SELECT token, ticker, type, bitcoin, activity, currency, extend_activity, decimals, decorator, currency_symbol FROM tickers'
)

item = get_cur.fetchone()
while item:
    if not item[1]:
        item = get_cur.fetchone()
        continue

    token = item[0]
    ticker = item[1]
    typ = item[2]
    bitcoin = item[3]
    activity = item[4]
    currency = item[5]
    extend_activity = item[6]
    decimals = item[7]
    decorator = item[8]
    currency_symbol = item[9]

    print(f'{typ}: {ticker}')

    # GAS
    if typ == "gas":
        data = {
            "network": ticker,
            "frequency": 10,
            "set_nickname": True,
            "discord_bot_token": token
        }

        if activity:
            data['activity'] = activity

        resp = post(
            f'http://{sys.argv[2]}/gas',
            data=dumps(data)
        )
    
    # TOKEN
    elif typ.startswith('token'):
        data = {
            "network": ticker.rsplit("-", 1)[0],
            "contract": ticker.rsplit("-", 1)[1],
            "name": typ.split('token-', 1).pop(),
            "frequency": 10,
            "set_nickname": True,
            "discord_bot_token": token
        }

        if activity:
            data['activity'] = activity

        resp = post(
            f'http://{sys.argv[2]}/token',
            data=dumps(data)
        )

    # HODLER
    elif typ.startswith('holders'):
        data = {
            "network": ticker.rsplit("-", 1)[0],
            "address": ticker.rsplit("-", 1)[1],
            "name": typ.split('holders-', 1).pop(),
            "frequency": 10,
            "set_nickname": True,
            "discord_bot_token": token
        }

        if activity:
            data['activity'] = activity

        resp = post(
            f'http://{sys.argv[2]}/holders',
            data=dumps(data)
        )

    # TICKER
    else:
        data = {
            "ticker": ticker,
            "frequency": 60,
            "bitcoin": bitcoin == 1,
            "discord_bot_token": token
        }
        
        if currency:
            data['currency'] = currency
        if extend_activity:
            data['extend_activity'] = True
        if decimals:
            data['decimals'] = int(decimals)
        if decorator:
            data['decorator'] = decorator
        if currency_symbol:
            data['currency_symbol'] = currency_symbol

        if typ.startswith('crypto'):
            data['crypto'] = True
            data['name'] = ticker
            if typ.startswith('crypto-'):
                data['ticker'] = typ.split('crypto-', 1).pop()
            else:
                data['ticker'] = ''
        else:
            data['crypto'] = False
            if typ != 'stock':
                data['name'] = typ

        if len(sys.argv) == 4:
            data['frequency'] = 10
            data['set_nickname'] = True
            data['set_color'] = True
            if activity:
                data['activity'] = activity

        resp = post(
            f'http://{sys.argv[2]}/ticker',
            data=dumps(data)
        )

    print(resp.status_code)
    print(resp.text)

    sleep(2)

    item = get_cur.fetchone()
