# discord-stock-ticker

Live stock tickers for your discord server.

![Discord Sidebar w/ Bots](/assets/sidebar.png)

Due to discord limitations the prices in the usernames will be updated every hour while the prices in the activity section update every 60 seconds.

Use the free bots listed below or see options to host some yourself!

## Add tickers to your servers

Stock | Crypto
------------ | -------------
[![GameStop](https://logo.clearbit.com/gamestop.com)](https://discord.com/api/oauth2/authorize?client_id=805268557994262529&permissions=0&scope=bot) | [![Bitcoin](https://logo.clearbit.com/bitcoin.org)](https://discord.com/api/oauth2/authorize?client_id=805599050871210014&permissions=0&scope=bot)
[![Blackberry](https://logo.clearbit.com/blackberry.com)](https://discord.com/api/oauth2/authorize?client_id=805289769272999986&permissions=0&scope=bot) | [![Bitcoin Cash](https://logo.clearbit.com/bitcoin.com)](https://discord.com/api/oauth2/authorize?client_id=805604560013230170&permissions=0&scope=bot)
[![AMC Theatres](https://logo.clearbit.com/amctheatres.com)](https://discord.com/api/oauth2/authorize?client_id=805294017441038357&permissions=0&scope=bot) | [![Ethereum](https://logo.clearbit.com/ethereum.org)](https://discord.com/api/oauth2/authorize?client_id=805605209522962452&permissions=0&scope=bot)
[![Nokia](https://logo.clearbit.com/nokia.com)](https://discord.com/api/oauth2/authorize?client_id=805294107962245120&permissions=0&scope=bot) | [![Dogecoin](https://logo.clearbit.com/dogecoin.com)](https://discord.com/api/oauth2/authorize?client_id=805605888387186699&permissions=0&scope=bot)
[![Principal Financial Group](https://logo.clearbit.com/principal.com)](https://discord.com/api/oauth2/authorize?client_id=805466470930055189&permissions=0&scope=bot) | [![Monero](https://logo.clearbit.com/getmonero.org)](https://discord.com/api/oauth2/authorize?client_id=806282848045629451&permissions=0&scope=bot)
[![Apple](https://logo.clearbit.com/apple.com)](https://discord.com/api/oauth2/authorize?client_id=806569145184550922&permissions=0&scope=bot) | [![Litecoin](https://logo.clearbit.com/litecoin.com)](https://discord.com/api/oauth2/authorize?client_id=806635240482668574&permissions=0&scope=bot)
[![Amazon](https://logo.clearbit.com/amazon.com)](https://discord.com/api/oauth2/authorize?client_id=806570287042002945&permissions=0&scope=bot) | [![Ripple](https://logo.clearbit.com/ripple.com)](https://discord.com/api/oauth2/authorize?client_id=806634757168693258&permissions=0&scope=bot)
[![Alphabet](https://logo.clearbit.com/google.com)](https://discord.com/api/oauth2/authorize?client_id=806570628156882945&permissions=0&scope=bot) | [![Polkadot](https://logo.clearbit.com/polkadot.network)](https://discord.com/api/oauth2/authorize?client_id=806633568787890217&permissions=0&scope=bot)
[![](https://logo.clearbit.com/)](https://discord.com/api/oauth2/authorize?client_id=&permissions=0&scope=bot) | [![Cardano](https://logo.clearbit.com/cardano.com)](https://discord.com/api/oauth2/authorize?client_id=807311315055542272&permissions=0&scope=bot)


## Hosting

### Hosted by rssnyder

The bots above are hosted using [piku](https://github.com/piku/piku) on a Ubuntu 20.04 server. They are free to use and should have little to no downtime. There is a full logging stack that includes loki & promtail with grafana for visualization.

If you encounter any issues with the bots listed above please see the support options at the bottom of this page.

![Really cool grafana dashboard](/assets/grafana.png)

### Self-Hosting

To run for youself, simply set DISCORD_BOT_TOKEN and TICKER in your environment, and run `main.py`.

You will need one bot for every ticker you want to add to your server.

#### Stocks

```
git clone git@github.com:rssnyder/discord-stock-ticker.git && cd discord-stock-ticker

pip install -r requirements.txt

export DISCORD_BOT_TOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
export TICKER=AAPL

python main.py
```

#### Crypto

If you want to watch a crypto, you must also set CRYPTO_NAME, where CRYPTO_NAME is the full name (eg. bitcoin) and TICKER is how you want the coin to appear (eg. BTC).

```
git clone git@github.com:rssnyder/discord-stock-ticker.git && cd discord-stock-ticker

pip install -r requirements.txt

export DISCORD_BOT_TOKEN=xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
export TICKER=BTC
export CRYPTO_NAME=bitcoin

python main.py
```

To see a list of cryptos avalible (we are using the coingecko API):

```
curl -X GET "https://api.coingecko.com/api/v3/coins/list" -H  "accept: application/json" | jq '.[].id'
```

## Support

If you have a request for a new ticker or issues with a current one, please open a github issue or find me on discord at `jonesbooned#1111` or [join the support server](https://discord.gg/CQqnCYEtG7).

Love these bots? Maybe [buy me a coffee](https://ko-fi.com/rileysnyder)!