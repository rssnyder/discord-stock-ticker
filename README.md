# discord-stock-ticker

Live stock and crypto tickers for your discord server.

Now with five different types of tickers!

🍾 400+ public tickers with over 15k installs across 3k discord servers!

*Are you just looking to add free tickers to your discord server? Click the discord icon below to join the support server and get the list of avalible bots!*

[![Publish](https://github.com/rssnyder/discord-stock-ticker/actions/workflows/deploy.yml/badge.svg)](https://github.com/rssnyder/discord-stock-ticker/actions/workflows/deploy.yml)
[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/tterb/atomic-design-ui/blob/master/LICENSEs)

[![GitHub last commit](https://img.shields.io/github/last-commit/rssnyder/discord-stock-ticker.svg?style=flat)]()
[![GitHub stars](https://img.shields.io/github/stars/rssnyder/discord-stock-ticker.svg?style=social&label=Star)]()
[![GitHub watchers](https://img.shields.io/github/watchers/rssnyder/discord-stock-ticker.svg?style=social&label=Watch)]()

## Contents

- [discord-stock-ticker](#discord-stock-ticker)
  - [Contents](#contents)
  - [Preview](#preview)
  - [Join the discord server](#join-the-discord-server)
  - [Support this project](#support-this-project)
  - [Add free tickers to your servers](#add-free-tickers-to-your-servers)
    - [Stocks](#stocks)
    - [Crypto](#crypto)
    - [Gas Prices](#gas-prices)
    - [Other crypto bots I make (click for details)](#other-crypto-bots-i-make-click-for-details)
  - [Premium](#premium)
  - [Self-Hosting - Docker](#self-hosting---docker)
  - [Self-Hosting - binary](#self-hosting---binary)
    - [Roles for colors](#roles-for-colors)
    - [Using the binary](#using-the-binary)
      - [Setting options](#setting-options)
      - [Systemd service](#systemd-service)
  - [Stock and Crypto Price Tickers](#stock-and-crypto-price-tickers)
    - [List current running bots](#list-current-running-bots)
    - [Add a new bot](#add-a-new-bot)
    - [Restart a bot](#restart-a-bot)
    - [Remove a bot](#remove-a-bot)
  - [Crypto Market Cap](#crypto-market-cap)
    - [List current running bots](#list-current-running-bots-1)
    - [Add a new bot](#add-a-new-bot-1)
    - [Restart a bot](#restart-a-bot-1)
    - [Remove a bot](#remove-a-bot-1)
  - [Stock and Crypto Price Tickerboards](#stock-and-crypto-price-tickerboards)
    - [List current running Boards](#list-current-running-boards)
    - [Add a new Board](#add-a-new-board)
    - [Restart a Board](#restart-a-board)
    - [Remove a Board](#remove-a-board)
  - [Ethereum, BSC, and Polygon Gas Prices](#ethereum-bsc-and-polygon-gas-prices)
    - [List current running Gas](#list-current-running-gas)
    - [Add a new Gas](#add-a-new-gas)
    - [Restart a Gas](#restart-a-gas)
    - [Remove a Gas](#remove-a-gas)
  - [Ethereum, BSC, or Polygon Token Holders](#ethereum-bsc-or-polygon-token-holders)
    - [List current running Holders](#list-current-running-holders)
    - [Add a new Holder](#add-a-new-holder)
    - [Restart a Holder](#restart-a-holder)
    - [Remove a Holder](#remove-a-holder)
  - [ETH/BSC/MATIC Token Price](#ethbscmatic-token-price)
    - [List current running Tokens](#list-current-running-tokens)
    - [Add a new Token](#add-a-new-token)
    - [Restart a Token](#restart-a-token)
    - [Remove a Token](#remove-a-token)
  - [OpenSea/Solanart NFT Collection Floor Price](#openseasolanart-nft-collection-floor-price)
    - [List current running Floors](#list-current-running-floors)
    - [Add a new Floor](#add-a-new-floor)
    - [Restart a Floor](#restart-a-floor)
    - [Remove a Floor](#remove-a-floor)
  - [Kubernetes](#kubernetes)
  - [Louie](#louie)

## Preview

![image](https://user-images.githubusercontent.com/7338312/127577682-70b67f31-59c9-427b-b9dc-2736a2b4e378.png)![TICKERS](https://user-images.githubusercontent.com/7338312/126001327-2d7167d2-e998-4e13-9272-61feb4e9bf7a.png)![BOARDS](https://user-images.githubusercontent.com/7338312/126001753-4f0ec66e-5737-495a-a85b-cafeef6f5cea.gif)![image](https://user-images.githubusercontent.com/7338312/127577601-43500287-1cf4-47ee-9f21-67c22f606850.png)![HOLDERS](https://user-images.githubusercontent.com/7338312/126001392-dfb72cc1-d526-40e8-9982-077bb22fc44c.png)

## Join the discord server

[![Discord Chat](https://logo.clearbit.com/discord.com)](https://discord.gg/CQqnCYEtG7)

## Support this project

<a href='https://ko-fi.com/rileysnyder' target='_blank'><img height='35' style='border:0px;height:46px;' src='https://az743702.vo.msecnd.net/cdn/kofi3.png?v=0' border='0' alt='Buy Me a Coffee' /></a>

<a href="https://www.digitalocean.com/?refcode=1acd6d377e8b&utm_campaign=Referral_Invite&utm_medium=Referral_Program&utm_source=badge"><img src="https://web-platforms.sfo2.digitaloceanspaces.com/WWW/Badge%203.svg" alt="DigitalOcean Referral Badge" /></a>

Love these bots? You can support this project by subscribing to the [premium version](https://github.com/rssnyder/discord-stock-ticker/blob/master/README.md#premium), [buying me a coffee](https://ko-fi.com/rileysnyder), [using my digital ocean referral link](https://m.do.co/c/1acd6d377e8b), or [hiring me](https://github.com/rssnyder) to write or host **your** discord bot!

## Add free tickers to your servers

### Stocks

[bb](https://discord.com/oauth2/authorize?client_id=805289769272999986&permissions=0&scope=bot) |
[amc](https://discord.com/oauth2/authorize?client_id=805294017441038357&permissions=0&scope=bot) |
[nok](https://discord.com/oauth2/authorize?client_id=805294107962245120&permissions=0&scope=bot) |
[aapl](https://discord.com/oauth2/authorize?client_id=806569145184550922&permissions=0&scope=bot) |
[amzn](https://discord.com/oauth2/authorize?client_id=806570287042002945&permissions=0&scope=bot) |
[goog](https://discord.com/oauth2/authorize?client_id=806570628156882945&permissions=0&scope=bot) |
[^gspc](https://discord.com/oauth2/authorize?client_id=808431853363134484&permissions=0&scope=bot) |
[^dji](https://discord.com/oauth2/authorize?client_id=808432655746072596&permissions=0&scope=bot) |
[^ixic](https://discord.com/oauth2/authorize?client_id=808432811644026921&permissions=0&scope=bot) |
[tsla](https://discord.com/oauth2/authorize?client_id=808723743069306882&permissions=0&scope=bot) |
[dkng](https://discord.com/oauth2/authorize?client_id=808724381608968202&permissions=0&scope=bot) |
[spy](https://discord.com/oauth2/authorize?client_id=811418568846737500&permissions=0&scope=bot) |
[amd](https://discord.com/oauth2/authorize?client_id=816049122850897923&permissions=0&scope=bot) |
[nio](https://discord.com/oauth2/authorize?client_id=816049780546994196&permissions=0&scope=bot) |
[gc=f](https://discord.com/oauth2/authorize?client_id=816375134122147850&permissions=0&scope=bot) |
[si=f](https://discord.com/oauth2/authorize?client_id=816375586066661456&permissions=0&scope=bot) |
[cl=f](https://discord.com/oauth2/authorize?client_id=816375780442636288&permissions=0&scope=bot) |
[pltr](https://discord.com/oauth2/authorize?client_id=818471352415551518&permissions=0&scope=bot) |
[pypl](https://discord.com/oauth2/authorize?client_id=819395115465572352&permissions=0&scope=bot) |
[sndl](https://discord.com/oauth2/authorize?client_id=819638060894388305&permissions=0&scope=bot) |
[rty=f](https://discord.com/oauth2/authorize?client_id=819578053881102336&permissions=0&scope=bot) |
[^vix](https://discord.com/oauth2/authorize?client_id=819576744176386048&permissions=0&scope=bot) |
[arkk](https://discord.com/oauth2/authorize?client_id=826512765580869673&permissions=0&scope=bot) |
[msft](https://discord.com/oauth2/authorize?client_id=827236206575222794&permissions=0&scope=bot) |
[nflx](https://discord.com/oauth2/authorize?client_id=827236295323025408&permissions=0&scope=bot) |
[gme](https://discord.com/oauth2/authorize?client_id=828417475624435712&permissions=0&scope=bot) |
[dis](https://discord.com/oauth2/authorize?client_id=828613902740357121&permissions=0&scope=bot) |
[es=f](https://discord.com/oauth2/authorize?client_id=831958978455535667&permissions=0&scope=bot) |
[nq=f](https://discord.com/oauth2/authorize?client_id=833759608861622353&permissions=0&scope=bot) |
[ym=f](https://discord.com/oauth2/authorize?client_id=833759696783409195&permissions=0&scope=bot) |
[nvda](https://discord.com/oauth2/authorize?client_id=836385897372581910&permissions=0&scope=bot) |
[fb](https://discord.com/oauth2/authorize?client_id=836385990113361951&permissions=0&scope=bot) |
[btc](https://discord.com/oauth2/authorize?client_id=839289136418783322&permissions=0&scope=bot) |
[coin](https://discord.com/oauth2/authorize?client_id=848232686320484382&permissions=0&scope=bot) |
[ndaq](https://discord.com/oauth2/authorize?client_id=851203267840966696&permissions=0&scope=bot) |
[qqq](https://discord.com/oauth2/authorize?client_id=854499452840312842&permissions=0&scope=bot) |

### Crypto

[bitcoin](https://discord.com/oauth2/authorize?client_id=828417381779898368&permissions=0&scope=bot) |
[ethereum](https://discord.com/oauth2/authorize?client_id=805605209522962452&permissions=0&scope=bot) |
[bitcoin-cash](https://discord.com/oauth2/authorize?client_id=805604560013230170&permissions=0&scope=bot) |
[dogecoin](https://discord.com/oauth2/authorize?client_id=805605888387186699&permissions=0&scope=bot) |
[monero](https://discord.com/oauth2/authorize?client_id=806282848045629451&permissions=0&scope=bot) |
[litecoin](https://discord.com/oauth2/authorize?client_id=806635240482668574&permissions=0&scope=bot) |
[ripple](https://discord.com/oauth2/authorize?client_id=806634757168693258&permissions=0&scope=bot) |
[polkadot](https://discord.com/oauth2/authorize?client_id=806633568787890217&permissions=0&scope=bot) |
[cardano](https://discord.com/oauth2/authorize?client_id=807311315055542272&permissions=0&scope=bot) |
[chainlink](https://discord.com/oauth2/authorize?client_id=808407486860230747&permissions=0&scope=bot) |
[stellar](https://discord.com/oauth2/authorize?client_id=808409647731179534&permissions=0&scope=bot) |
[iota](https://discord.com/oauth2/authorize?client_id=814170376254652486&permissions=0&scope=bot) |
[reef-finance](https://discord.com/oauth2/authorize?client_id=814288538107379742&permissions=0&scope=bot) |
[algorand](https://discord.com/oauth2/authorize?client_id=819274628778164265&permissions=0&scope=bot) |
[tezos](https://discord.com/oauth2/authorize?client_id=811609668991975484&permissions=0&scope=bot) |
[ethereum-classic](https://discord.com/oauth2/authorize?client_id=819395405980762182&permissions=0&scope=bot) |
[ravencoin](https://discord.com/oauth2/authorize?client_id=819395519708921866&permissions=0&scope=bot) |
[binancecoin](https://discord.com/oauth2/authorize?client_id=819395643193688124&permissions=0&scope=bot) |
[ecomi](https://discord.com/oauth2/authorize?client_id=819939716228579360&permissions=0&scope=bot) |
[aave](https://discord.com/oauth2/authorize?client_id=826512401502961756&permissions=0&scope=bot) |
[uniswap](https://discord.com/oauth2/authorize?client_id=827985872389275658&permissions=0&scope=bot) |
[bittorrent-2](https://discord.com/oauth2/authorize?client_id=827986251264819201&permissions=0&scope=bot) |
[tron](https://discord.com/oauth2/authorize?client_id=828326036785463307&permissions=0&scope=bot) |
[vechain](https://discord.com/oauth2/authorize?client_id=828326223306424350&permissions=0&scope=bot) |
[illuvium](https://discord.com/oauth2/authorize?client_id=828417571354968104&permissions=0&scope=bot) |
[cosmos](https://discord.com/oauth2/authorize?client_id=828417242570948638&permissions=0&scope=bot) |
[zilliqa](https://discord.com/oauth2/authorize?client_id=828417678976745483&permissions=0&scope=bot) |
[matic-network](https://discord.com/oauth2/authorize?client_id=828613785345458206&permissions=0&scope=bot) |
[basic-attention-token](https://discord.com/oauth2/authorize?client_id=828613810355961898&permissions=0&scope=bot) |
[shiba-inu](https://discord.com/oauth2/authorize?client_id=829119870556831816&permissions=0&scope=bot) |
[pancakeswap-token](https://discord.com/oauth2/authorize?client_id=831957913819021322&permissions=0&scope=bot) |
[solana](https://discord.com/oauth2/authorize?client_id=836339329084948490&permissions=0&scope=bot) |
[raydium](https://discord.com/oauth2/authorize?client_id=836385816409407488&permissions=0&scope=bot) |
[safemoon](https://discord.com/oauth2/authorize?client_id=837015732919074878&permissions=0&scope=bot) |
[ftx-token](https://discord.com/oauth2/authorize?client_id=837015812677304380&permissions=0&scope=bot) |
[enjincoin](https://discord.com/oauth2/authorize?client_id=838824692478771232&permissions=0&scope=bot) |
[decentraland](https://discord.com/oauth2/authorize?client_id=840023058734776351&permissions=0&scope=bot) |
[fantom](https://discord.com/oauth2/authorize?client_id=841412194578071612&permissions=0&scope=bot) |
[coti](https://discord.com/oauth2/authorize?client_id=841412539854094416&permissions=0&scope=bot) |
[hedera-hashgraph](https://discord.com/oauth2/authorize?client_id=841850253418430495&permissions=0&scope=bot) |
[sushi](https://discord.com/oauth2/authorize?client_id=844012506670432277&permissions=0&scope=bot) |
[kusama](https://discord.com/oauth2/authorize?client_id=844934710597779456&permissions=0&scope=bot) |
[eos](https://discord.com/oauth2/authorize?client_id=844934770314969109&permissions=0&scope=bot) |
[terra-luna](https://discord.com/oauth2/authorize?client_id=849375907856384020&permissions=0&scope=bot) |
[chia](https://discord.com/oauth2/authorize?client_id=851203175200456704&permissions=0&scope=bot) |
[theta-token](https://discord.com/oauth2/authorize?client_id=852238599163871286&permissions=0&scope=bot) |
[tether](https://discord.com/oauth2/authorize?client_id=857358033709301792&permissions=0&scope=bot) |
[smooth-love-potion](https://discord.com/oauth2/authorize?client_id=859148505469943860&permissions=0&scope=bot) |
[axie-infinity](https://discord.com/oauth2/authorize?client_id=859932709262589992&permissions=0&scope=bot) |
[harmony](https://discord.com/oauth2/authorize?client_id=862731942193791015&permissions=0&scope=bot) |
[cryptoblades](https://discord.com/oauth2/authorize?client_id=867745267327041556&permissions=0&scope=bot) |
[my-defi-pet](https://discord.com/oauth2/authorize?client_id=869247729132515418&permissions=0&scope=bot) |
[mist](https://discord.com/oauth2/authorize?client_id=869998726628925480&permissions=0&scope=bot) |
[weth](https://discord.com/oauth2/authorize?client_id=871068140623257610&permissions=0&scope=bot) |
[plant-vs-undead-token](https://discord.com/oauth2/authorize?client_id=871812819266461807&permissions=0&scope=bot) |
[cryptozoon](https://discord.com/oauth2/authorize?client_id=873191092659572797&permissions=0&scope=bot) |
[binance-usd](https://discord.com/oauth2/authorize?client_id=873191153825095720&permissions=0&scope=bot) |
[splinterlands](https://discord.com/oauth2/authorize?client_id=874299210076393503&permissions=0&scope=bot) |
[wax](https://discord.com/oauth2/authorize?client_id=875432263184187393&permissions=0&scope=bot) |
[coinary-token](https://discord.com/oauth2/authorize?client_id=877903037564346398&permissions=0&scope=bot) |
[avalanche-2](https://discord.com/oauth2/authorize?client_id=877903202211754094&permissions=0&scope=bot) |
[cryptocars](https://discord.com/oauth2/authorize?client_id=879447174210011146&permissions=0&scope=bot) |
[binamon](https://discord.com/oauth2/authorize?client_id=880628836679680000&permissions=0&scope=bot) |
[wanaka-farm](https://discord.com/oauth2/authorize?client_id=882420667725795438&permissions=0&scope=bot) |

### Gas Prices

[![Ethereum Invite Link](https://user-images.githubusercontent.com/7338312/127579033-8785ed17-2bcc-474c-80d4-8ea356da70e6.png)](https://discord.com/api/oauth2/authorize?client_id=833797002684661821&permissions=0&scope=bot)[![Binance Smart Chain Invite Link](https://user-images.githubusercontent.com/7338312/127578976-d47069cb-c162-4ab5-ad73-be17b2c1796d.png)](https://discord.com/api/oauth2/authorize?client_id=856947934452645898&permissions=0&scope=bot)[![Polygon Invite Link](https://user-images.githubusercontent.com/7338312/127578967-a7097067-9b0a-44d2-baf6-e3541a511c70.png)](https://discord.com/api/oauth2/authorize?client_id=857023179210096674&permissions=0&scope=bot)

### Other crypto bots I make (click for details)

[![image](https://user-images.githubusercontent.com/7338312/147755805-b2562443-205c-44d3-8742-64d90ff81963.png)](https://github.com/rssnyder/discord-nft-floor-price)

## Premium

For advanced features like faster update times and color changing names on price changes you can subscribe to my premuim offering.

Price per bot (paid monthly): $1

Price per bot (paid yearly):  $10

If you are interested please see the [contact info on my github page](https://github.com/rssnyder) and send me a messgae via your platform of choice (discord perferred). For a live demo, join the support discord linked at the top or bottom of this page.

## Self-Hosting - Docker

Grab the current release number from the [release page](https://github.com/rssnyder/discord-stock-ticker/releases) and expose your designated API port:

```shell
docker run -p "8080:8080" ghcr.io/rssnyder/discord-stock-ticker:3.4.1
```

You can set the config via ENV vars, since we use [namsral/flag](https://github.com/namsral/flag) the variables are the same as the flag inputs, but all uppercase:

When using the binary...

```shell
  -address="localhost:8080": address:port to bind http server to.
  -cache=false: enable cache for coingecko
  -db="": file to store tickers in
  -frequency=0: set frequency for all tickers
  -logLevel=0: defines the log level. 0=production builds. 1=dev builds.
  -redisAddress="localhost:6379": address:port for redis server.
  -redisDB=0: redis db to use
  -redisPassword="": redis password
```

When using env (docker)...

```shell
export ADDRESS="localhost:8080" # address:port to bind http server to.
export CACHE=false # enable cache for coingecko
export DB="" # file to store tickers in
export FREQUENCY=60 # set frequency for all tickers
export LOGLEVEL=0 # defines the log level. 0=production builds. 1=dev builds.
export REDISADDRESS="localhost:6379" # address:port for redis server.
export REDISDB=0 # redis db to use
export REDISPASSWORD="" # redis password
```

```shell
docker run -p "8080:8080" --env CACHE=true ghcr.io/rssnyder/discord-stock-ticker:3.4.1
```

Then you can pass a volume to store the state (and at the same time, upgrade to using [docker-compose](https://docs.docker.com/compose/)):

```shell
---
version: "3"
services:

  discordstockticker:
    image: ghcr.io/rssnyder/discord-stock-ticker:3.4.1
    environment:
      - DB=/dst.db
      - CACHE=true
    volumes:
      - /home/infra/dst.db:/dst.db
    ports:
      - "8112:8080"
```

## Self-Hosting - binary

This bot is distributed as a docker image and a binary.

The program acts as a manager of one to many bots. You can have one running instance of the program and have any number of bots running within it.

[Click here](https://youtu.be/LhgCdtE8kmc) to watch a quick video tutorial on how to self-host these bots on linux.

If you are using windows and do not have a unix shell to use, you should use powershell. Here is an example of an API call using powershell:

```powershell
$Body = @{
  name = "bitcoin"
  crypto = $true
  discord_bot_token = "xxxxxxxxxxxxxxxxxxxxxxxxx"
}
 
$Parameters = @{
    Method = "POST"
    Uri =  "127.0.0.1:8080/ticker"
    Body = ($Body | ConvertTo-Json) 
    ContentType = "application/json"
}

Invoke-RestMethod @Parameters
```

### Roles for colors

To enable color changing you will need to create three roles.

The first role is the role the tickers will appear under. It can be named _anything you want_. You need to check the **Display role members seperatly from other online members** option for this role, but _do not_ assign a custom color for this role, leave it default.

Then you need to make two other roles. These roles need to be named _exactly_ **tickers-red** & **tickers-green**. **Do not** check the Display role members seperatly from other online members option for these roles, but do assign colors to these roles, red and green (or whatever color you want to represent gain/loss) respectively.

The last two roles tickers-green and tickers-red need to be below the first role in the role list in your server settings. You should then add all your ticker bots to the first role.

![roles example](https://user-images.githubusercontent.com/7338312/131678207-b1510955-f762-46e3-ae5c-1b5eddb68844.jpg)

### Using the binary

Pull down the latest release for your OS [here](https://github.com/rssnyder/discord-stock-ticker/releases).

```shell
wget https://github.com/rssnyder/discord-stock-ticker/releases/download/v2.0.0/discord-stock-ticker-v3.3.0-linux-amd64.tar.gz

tar zxf discord-stock-ticker-v3.3.0-linux-amd64.tar.gz

./discord-stock-ticker
```

#### Setting options

There are options you can set for the service using flags:

```shell
  -address="localhost:8080": address:port to bind http server to.
  -cache=false: enable cache for coingecko
  -db="": file to store tickers in
  -frequency=0: set frequency for all tickers
  -logLevel=0: defines the log level. 0=production builds. 1=dev builds.
  -redisAddress="localhost:6379": address:port for redis server.
  -redisDB=0: redis db to use
  -redisPassword="": redis password
```

#### Systemd service

The script here (ran as root) will download and install a `discord-stock-ticker` service on your linux machine with an API avalible on port `8080` to manage bots.

```shell
wget https://github.com/rssnyder/discord-stock-ticker/releases/download/v3.3.0/discord-stock-ticker-v3.3.0-linux-amd64.tar.gz

tar zxf discord-stock-ticker-v3.3.0-linux-amd64.tar.gz

mkdir -p /etc/discord-stock-ticker

mv discord-stock-ticker /etc/discord-stock-ticker/

wget https://raw.githubusercontent.com/rssnyder/discord-stock-ticker/master/discord-stock-ticker.service

mv discord-stock-ticker.service /etc/systemd/system/

systemctl daemon-reload

systemctl start discord-stock-ticker.service
```

If you need to make modifications to the setting of the service, just edit the `/etc/systemd/system/discord-stock-ticker.service` file on the line with `ExecStart=`.

Now that you have the service running, you can add bots using the API exposed on the addres and port that the service runs on (this address is shown when you start the service).

## Stock and Crypto Price Tickers

### List current running bots

```shell
curl localhost:8080/ticker
```

### Add a new bot

Stock Payload:

```json
{
  "ticker": "pfg",                                  # string: symbol for the stock from yahoo finance
  "name": "2) PFG",                                 # string/OPTIONAL: overwrites display name of bot
  "set_color": true,                                # bool/OPTIONAL: requires set_nickname
  "decorator": "@",                                 # string/OPTIONAL: what to show instead of arrows
  "currency": "aud",                                # string/OPTIONAL: alternative curreny
  "activity": "Hello;Its;Me",                       # string/OPTIONAL: list of strings to show in activity section
  "set_nickname": true,                             # bool/OPTIONAL: display information in nickname vs activity
  "frequency": 10,                                  # int/OPTIONAL: seconds between refresh
  "twelve_data_key": "xxx",                         # string/OPTIONAL: use twelve data as source, pass in api key
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxx"   # string: dicord bot token
}
```

Crypto Payload:

```json
{
  "name": "bitcoin",                                # string: name of the crypto from coingecko
  "crypto": true,                                   # bool: always true for crypto
  "ticker": "1) BTC",                               # string/OPTIONAL: overwrites display name of bot
  "set_color": true,                                # bool/OPTIONAL: requires set_nickname
  "decorator": "@",                                 # string/OPTIONAL: what to show instead of arrows
  "currency": "aud",                                # string/OPTIONAL: alternative curreny
  "currency_symbol": "AUD",                         # string/OPTIONAL: alternative curreny symbol
  "pair": "binancecoin",                            # string/OPTIONAL: pair the coin with another coin, replaces activity section
  "pair_flip": true,                                # bool/OPTIONAL: show <pair>/<coin> rather than <coin>/<pair>
  "activity": "Hello;Its;Me",                       # string/OPTIONAL: list of strings to show in activity section
  "decimals": 3,                                    # int/OPTIONAL: set number of decimal places
  "set_nickname": true,                             # bool/OPTIONAL: display information in nickname vs activity
  "frequency": 10,                                  # int/OPTIONAL: seconds between refresh
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxx"   # string: dicord bot token
}
```

Example:

```shell
curl -X POST -H "Content-Type: application/json" --data '{
  "ticker": "pfg",
  "name": "PFG",
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxxx"
}' localhost:8080/ticker
```

### Restart a bot

```shell
curl -X PATCH localhost:8080/ticker/pfg
```

```shell
curl -X PATCH localhost:8080/ticker/bitcoin
```

### Remove a bot

```shell
curl -X DELETE localhost:8080/ticker/pfg
```

```shell
curl -X DELETE localhost:8080/ticker/bitcoin
```

## Crypto Market Cap

### List current running bots

```shell
curl localhost:8080/marketcap
```

### Add a new bot

```json
{
  "name": "bitcoin",                                # string: name of the crypto from coingecko
  "ticker": "1) BTC",                               # string/OPTIONAL: overwrites display name of bot
  "set_color": true,                                # bool/OPTIONAL: requires set_nickname
  "decorator": "@",                                 # string/OPTIONAL: what to show instead of arrows
  "currency": "aud",                                # string/OPTIONAL: alternative curreny
  "currency_symbol": "AUD",                         # string/OPTIONAL: alternative curreny symbol
  "activity": "Hello;Its;Me",                       # string/OPTIONAL: list of strings to show in activity section
  "decimals": 3,                                    # int/OPTIONAL: set number of decimal places
  "set_nickname": true,                             # bool/OPTIONAL: display information in nickname vs activity
  "frequency": 10,                                  # int/OPTIONAL: seconds between refresh
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxx"   # string: dicord bot token
}
```

Example:

```shell
curl -X POST -H "Content-Type: application/json" --data '{
  "name": "bitcoin",
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxxx"
}' localhost:8080/marketcap
```

### Restart a bot

```shell
curl -X PATCH localhost:8080/marketcap/bitcoin
```

### Remove a bot

```shell
curl -X DELETE localhost:8080/marketcap/bitcoin
```

## Stock and Crypto Price Tickerboards

Tickerboards are tickers that rotate though several stocks or cryptos. This bot is a newer release, and is not as stable as the rest of the bots.

![BOARDS](https://user-images.githubusercontent.com/7338312/126001753-4f0ec66e-5737-495a-a85b-cafeef6f5cea.gif)

### List current running Boards

```shell
curl localhost:8080/tickerboard
```

### Add a new Board

Stock Payload:

```json
{
  "name": "Stocks",                                 # string: name of your board
  "items": ["PFG", "GME", "AMC"],                   # list of strings: symbols from yahoo finance to rotate through
  "header": "1. ",                                  # string/OPTIONAL: adds a header to the nickname to help sort bots
  "set_color": true,                                # bool/OPTIONAL: requires set_nickname
  "arrows": true,                                   # bool/OPTIONAL: show arrows in ticker names
  "set_nickname": true,                             # bool/OPTIONAL: display information in nickname vs activity
  "frequency": 10,                                  # int/OPTIONAL: seconds between refresh
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxx"   # string: dicord bot token
}
```

Crypto Payload:

```json
{
  "name": "Cryptos",                                # string: name of your board
  "crypto": true,                                   # bool: always true for crypto
  "items": ["bitcoin", "ethereum", "dogecoin"],     # list of strings: names from coingecko to rotate through
  "header": "2. ",                                  # string/OPTIONAL: adds a header to the nickname to help sort bots
  "set_color": true,                                # bool/OPTIONAL: requires set_nickname
  "arrows": true,                                   # bool/OPTIONAL: show arrows in ticker names
  "set_nickname": true,                             # bool/OPTIONAL: display information in nickname vs activity
  "frequency": 10,                                  # int/OPTIONAL: seconds between refresh
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxx"   # string: dicord bot token
}
```

Example:

```shell
curl -X POST -H "Content-Type: application/json" --data '{
  "name": "Stocks",
  "frequency": 3,
  "set_nickname": true,
  "set_color": true,
  "percentage": true,
  "arrows": true,
  "discord_bot_token": "xxxxxxx",
  "items": ["PFG", "GME", "AMC"]
}' localhost:8080/tickerboard
```

### Restart a Board

```shell
curl -X PATCH localhost:8080/tickerboard/stocks
```

### Remove a Board

```shell
curl -X DELETE localhost:8080/tickerboard/stocks
```

## Ethereum, BSC, and Polygon Gas Prices

These bots shows the current recommended gas prices for three types of transactions. You can choose either the ethereum, binance smart chain, or polygon blockchain.

![image](https://user-images.githubusercontent.com/7338312/127577601-43500287-1cf4-47ee-9f21-67c22f606850.png)

### List current running Gas

```shell
curl localhost:8080/gas
```

### Add a new Gas

Payload:

```json
{
  "network": "ethereum",                            # string: one of: ethereum, binance-smart-chain, or polygon
  "set_nickname": true,                             # bool/OPTIONAL: display information in nickname vs activity
  "frequency": 10,                                  # int/OPTIONAL: seconds between refresh
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxx"   # string: dicord bot token
}
```

Example:

```shell
curl -X POST -H "Content-Type: application/json" --data '{
  "network": "polygon",
  "frequency": 3,
  "set_nickname": true,
  "discord_bot_token": "xxxxxxx"
}' localhost:8080/gas
```

### Restart a Gas

```shell
curl -X PATCH localhost:8080/gas/polygon
```

### Remove a Gas

```shell
curl -X DELETE localhost:8080/gas/polygon
```

## Ethereum, BSC, or Polygon Token Holders

This bot lists the number of addresses that hold a particular token. You can choose from the ethereum or binance smart chain blockchains.

![HOLDERS](https://user-images.githubusercontent.com/7338312/126001392-dfb72cc1-d526-40e8-9982-077bb22fc44c.png)

### List current running Holders

```shell
curl localhost:8080/holders
```

### Add a new Holder

Payload:

```json
{
  "network": "ethereum",                            # string: one of: ethereum, binance-smart-chain, or polygon
  "address": "0x00000000000000000000000000",        # string: address of contract for token
  "activity": "ethereum",                           # string: text to show in activity section of the bot
  "set_nickname": true,                             # bool/OPTIONAL: display information in nickname vs activity
  "frequency": 10,                                  # int/OPTIONAL: seconds between refresh
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxx"   # string: dicord bot token
}
```

Example:

```shell
curl -X POST -H "Content-Type: application/json" --data '{
  "network": "ethereum",
  "address": "0x00000000000000",
  "activity": "Holders of MyToken",
  "set_nickname": true,
  "frequency": 120,
  "discord_bot_token": "xxxxxxx"
}' localhost:8080/holders
```

### Restart a Holder

```shell
curl -X PATCH localhost:8080/holders/ethereum-0x00000000000000
```

### Remove a Holder

```shell
curl -X DELETE localhost:8080/holders/ethereum-0x00000000000000
```

## ETH/BSC/MATIC Token Price

This bot gets the current rate for a given token. You can choose another token to pair with on price, or by default USDC is used. You can choose either the ethereum, binance smart chain, or polygon blockchain.

### List current running Tokens

```shell
curl localhost:8080/token
```

### Add a new Token

Payload:

```json
{
  "network": "ethereum",                            # string: network of token, options are ethereum, binance-smart-chain, or polygon
  "name": "my token",                               # string: display name of token
  "contract": "0x00000",                            # string: contract address of token
  "currency": "0x00000",                            # string/OPTIONAL: contract address of token to price against, default is USDC
  "set_nickname": true,                             # bool/OPTIONAL: display information in nickname vs activity
  "set_color": true,                                # bool/OPTIONAL: requires set_nickname
  "decorator": "@",                                 # string/OPTIONAL: what to show instead of arrows
  "activity": "Hello;Its;Me",                       # string/OPTIONAL: list of strings to show in activity section
  "source": "pancakeswap",                          # string/OPTIONAL: if the token is a BSC token, you can set pancakeswap here to use it vs 1inch; you can also set dexlab for solana tokens
  "frequency": 10,                                  # int/OPTIONAL: seconds between refresh
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxx"   # string: dicord bot token
}
```

Example:

```shell
curl -X POST -H "Content-Type: application/json" --data '{
  "network": "polygon",
  "contract": "0x0000000",
  "frequency": 3,
  "set_nickname": true,
  "discord_bot_token": "xxxxxxx"
}' localhost:8080/token
```

### Restart a Token

```shell
curl -X PATCH localhost:8080/token/polygon-0x0000000
```

### Remove a Token

```shell
curl -X DELETE localhost:8080/token/polygon-0x0000000
```

## OpenSea/Solanart NFT Collection Floor Price

This bot gets the current floor price for a NFT collection.

### List current running Floors

```shell
curl localhost:8080/floor
```

### Add a new Floor

Payload:

```json
{
  "marketplace": "opensea",                         # string: one of: opensea or solanart
  "name": "ethereum",                               # string: one of: ethereum, binance-smart-chain, or polygon
  "set_nickname": true,                             # bool/OPTIONAL: display information in nickname vs activity
  "frequency": 10,                                  # int/OPTIONAL: seconds between refresh
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxx"   # string: dicord bot token
}
```

Example:

```shell
curl -X POST -H "Content-Type: application/json" --data '{
  "marketplace": "solanart",
  "name": "solpunks",
  "frequency": 3,
  "set_nickname": true,
  "discord_bot_token": "xxxxxxx"
}' localhost:8080/floor
```


### Restart a Floor

```shell
curl -X PATCH localhost:8080/floor/solanart-solpunks
```

### Remove a Floor

```shell
curl -X DELETE localhost:8080/floor/solanart-solpunks
```

## Kubernetes

Thanks to @jr0dd there is a helm chart for deploying to k8s clusters. His chart can be found [here](https://github.com/jr0dd/charts/tree/master/discord-stock-ticker)

You can also use a simple deployment file:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    environment: public
  name: ticker-cardano
spec:
  replicas: 1
  selector:
    matchLabels:
      environment: public
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        environment: public
    spec:
      containers:
        - env:
            - name: CRYPTO_NAME
              value: cardano
            - name: DISCORD_BOT_TOKEN
              value: xxxxxxxxxxxxxxxxxxxxxx
            - name: FREQUENCY
              value: "1"
            - name: SET_COLOR
              value: "1"
            - name: SET_NICKNAME
              value: "1"
            - name: TICKER
              value: ADA
            - name: TZ
              value: America/Chicago
          image: ghcr.io/rssnyder/discord-stock-ticker:1.8.1
          name: ticker-cardano
          resources: {}
      restartPolicy: Always
status: {}
```

## Louie

Since you have read this far, here is a picture of Louie at his favorite park:

![PXL_20210424_185951005 PORTRAIT](https://user-images.githubusercontent.com/7338312/129428365-38d1c7c5-547e-48d4-8702-44f35541eac5.jpg)
