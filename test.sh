#!/bin/bash

# curl -X POST localhost:8080/ticker --data '{
#   "crypto": true,
#   "ticker": "TICKER: ",
#   "set_color": true,
#   "decorator": "@",
#   "currency": "aud",
#   "pair": "binancecoin",
#   "pair_flip": true,
#   "activity": "Test Message!",
#   "decimals": 4,
#   "set_nickname": true,
#   "frequency": 3,
#   "discord_bot_token": "ODc4NzgxMDg4NzU4MTEyMzE3.YSGKqg.C0F30gkR3CXrZ5YzEtXd0fKMvw8"
# }'

curl -X POST localhost:8080/ticker --data '{
  "ticker": "PFG",
  "set_color": true,
  "decorator": "@",
  "currency": "aud",
  "decimals": 4,
  "set_nickname": true,
  "frequency": 3,
  "twelve_data_key": "97c79030bbf8423a8f15f2655c942f6e",
  "discord_bot_token": "ODc4NzgxMDg4NzU4MTEyMzE3.YSGKqg.C0F30gkR3CXrZ5YzEtXd0fKMvw8"
}'

curl -X POST localhost:8080/gas --data '{
  "network": "ethereum",
  "set_nickname": true,
  "frequency": 3,
  "discord_bot_token": "ODc4Nzg5NDA5MTE2NDE4MDg4.YSGSag.6FYuRH7l0ctREFsfxp3yd8x2rz0"
}'

curl -X POST localhost:8080/holders --data '{
  "network": "ethereum",
  "address": "0x1f9840a85d5af5bf1d1762f925bdaddc4201f984",
  "activity": "Test Message!",
  "set_nickname": true,
  "frequency": 3,
  "discord_bot_token": "ODc4NzkwNjA1NTMzMjMzMTUy.YSGThw.bjrdKySyhS4io9dtlxiu85MTsAY"
}'

curl -X POST localhost:8080/token --data '{
  "network": "solana",
  "name": "test ",
  "contract": "AtNnsY1AyRERWJ8xCskfz38YdvruWVJQUVXgScC1iPb",
  "set_nickname": true,
  "set_color": true,
  "decorator": "@",
  "activity": "Test Message!",
  "source": "dexlab",
  "frequency": 3,
  "discord_bot_token": "ODc4NzkxODYyOTMxMDU4NzA4.YSGUsw.HYHgoOuR-9c_gEc2dZCEfNEpBMQ"
}'

curl -X POST localhost:8080/tickerboard --data '{
  "name": "Cryptos",
  "crypto": true,
  "items": ["bitcoin", "ethereum", "dogecoin"],
  "header": "BOARD: ",
  "set_color": true,
  "arrows": true,
  "set_nickname": true,
  "frequency": 10,
  "discord_bot_token": "ODc4NzkyMDYyMTEwMTQyNDk0.YSGU4g.py8vZ_QXPsLTMEc6UnKeDQP7KtA"
}'
