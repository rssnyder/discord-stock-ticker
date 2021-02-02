# discord-stock-ticker

Live stock tickers for your discord server.

Due to discord limitations the prices in the usernames will be updated every hour while the prices in the activity section update every 60 seconds.

![Discord Sidebar w/ Bots](/assets/sidebar.png)

## Add tickers to your servers

Stock | Crypto
------------ | -------------
[![GameStop](https://logo.clearbit.com/gamestop.com)](https://discord.com/api/oauth2/authorize?client_id=805268557994262529&permissions=0&scope=bot) | [![Bitcoin](https://logo.clearbit.com/bitcoin.org)](https://discord.com/api/oauth2/authorize?client_id=805599050871210014&permissions=0&scope=bot)
[![Blackberry](https://logo.clearbit.com/blackberry.com)](https://discord.com/api/oauth2/authorize?client_id=805289769272999986&permissions=0&scope=bot) | [![Bitcoin Cash](https://logo.clearbit.com/bitcoin.com)](https://discord.com/api/oauth2/authorize?client_id=805604560013230170&permissions=0&scope=bot)
[![AMC Theatres](https://logo.clearbit.com/amctheatres.com)](https://discord.com/api/oauth2/authorize?client_id=805294017441038357&permissions=0&scope=bot) | [![Ethereum](https://logo.clearbit.com/ethereum.org)](https://discord.com/api/oauth2/authorize?client_id=805605209522962452&permissions=0&scope=bot)
[![Nokia](https://logo.clearbit.com/nokia.com)](https://discord.com/api/oauth2/authorize?client_id=805294107962245120&permissions=0&scope=bot) | [![Dogecoin](https://logo.clearbit.com/dogecoin.com)](https://discord.com/api/oauth2/authorize?client_id=805605888387186699&permissions=0&scope=bot)
[![Principal Financial Group](https://logo.clearbit.com/principal.com)](https://discord.com/api/oauth2/authorize?client_id=805466470930055189&permissions=0&scope=bot) | [![Monero](https://logo.clearbit.com/getmonero.org)](https://discord.com/api/oauth2/authorize?client_id=806282848045629451&permissions=0&scope=bot)

## Hosting

To run for youself, simply set DISCORD_BOT_TOKEN and TICKER in your environment, and run `main.py`.

If you want to watch a crypto, you must also set CRYPTO_NAME, where CRYPTO_NAME is the full name (eg. Bitcoin) and TICKER is how you want the coin to appear (eg. BTC).

The bots above are hosted using [piku](https://github.com/piku/piku) on a local server. The logging stack includes loki & promtail with grafana for visualization.

![Really cool grafana dashboard](/assets/grafana.png)

## Support

If you have a request for a new ticker or issues with a current one, please open a github issue or find me on discord at `jonesbooned#1111`.