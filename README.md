# discord-stock-ticker

Live stock and crypto tickers for your discord server.

With these bots you can track prices of...

- Coins and Tokens on CoinGecko
- Marketcaps of Coins and Tokens on CoinGecko
- Stocks on Yahoo Finance
- Tokens on Pancakeswap
- Tokens on Dexlab
- Tokens on 1Inch
- NFT Collections on OpenSea, Solanart, Solart, and Magiceden
- Gas on Ethereum, Binance, and Polygon Chains
- Number of holders of a token on Ethereum and Binance Chains

🍾 100 public tickers with over 15k installs across 3k discord servers!

🛠️ Use theis project to host your own tickers, or .[pay for custom tickers to be made](https://github.com/rssnyder/discord-stock-ticker/blob/master/README.md#premium).

[![Publish](https://github.com/rssnyder/discord-stock-ticker/actions/workflows/deploy.yml/badge.svg)](https://github.com/rssnyder/discord-stock-ticker/actions/workflows/deploy.yml)
[![MIT License](https://img.shields.io/apm/l/atomic-design-ui.svg?)](https://github.com/tterb/atomic-design-ui/blob/master/LICENSEs)

[![GitHub last commit](https://img.shields.io/github/last-commit/rssnyder/discord-stock-ticker.svg?style=flat)](https://github.com/rssnyder/discord-stock-ticker/pulse)
[![GitHub stars](https://img.shields.io/github/stars/rssnyder/discord-stock-ticker.svg?style=social&label=Star)](https://github.com/rssnyder/discord-stock-ticker/pulse)
[![GitHub watchers](https://img.shields.io/github/watchers/rssnyder/discord-stock-ticker.svg?style=social&label=Watch)](https://github.com/rssnyder/discord-stock-ticker/pulse)

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
  - [Premium](#premium)
  - [Self-Hosting](#self-hosting)
      - [Setting options](#setting-options)
      - [Systemd (linux)](#systemd-linux)
  - [Managing bots](#managing-bots)
  - [Stock and Crypto Price Tickers](#stock-and-crypto-price-tickers)
    - [Bot Configuration (stock)](#bot-configuration-stock)
    - [Bot Configuration (crypto)](#bot-configuration-crypto)
  - [Stock and Crypto Price Tickerboards](#stock-and-crypto-price-tickerboards)
    - [Bot Configuration (stock)](#bot-configuration-stock-1)
    - [Bot Configuration (crypto)](#bot-configuration-crypto-1)
  - [Crypto Market Cap](#crypto-market-cap)
    - [Bot Configuration](#bot-configuration)
  - [Crypto Circulating Supply](#crypto-circulating-supply)
    - [Bot Configuration](#bot-configuration-1)
  - [Crypto Total Value Locked](#crypto-total-value-locked)
    - [Bot Configuration](#bot-configuration-2)
  - [Ethereum, BSC, and Polygon Gas Prices](#ethereum-bsc-and-polygon-gas-prices)
    - [Bot Configuration](#bot-configuration-3)
  - [Ethereum, BSC, or Polygon Token Holders](#ethereum-bsc-or-polygon-token-holders)
    - [Bot Configuration](#bot-configuration-4)
  - [ETH/BSC/MATIC Token Price](#ethbscmatic-token-price)
    - [Bot Configuration](#bot-configuration-5)
  - [OpenSea/Solanart NFT Collection Floor Price](#openseasolanart-nft-collection-floor-price)
    - [Bot Configuration](#bot-configuration-6)
  - [Roles for colors](#roles-for-colors)
  - [Kubernetes](#kubernetes)
  - [Louie](#louie)

## Preview

![image](https://user-images.githubusercontent.com/7338312/127577682-70b67f31-59c9-427b-b9dc-2736a2b4e378.png)![TICKERS](https://user-images.githubusercontent.com/7338312/126001327-2d7167d2-e998-4e13-9272-61feb4e9bf7a.png)![BOARDS](https://user-images.githubusercontent.com/7338312/126001753-4f0ec66e-5737-495a-a85b-cafeef6f5cea.gif)![image](https://user-images.githubusercontent.com/7338312/127577601-43500287-1cf4-47ee-9f21-67c22f606850.png)![HOLDERS](https://user-images.githubusercontent.com/7338312/126001392-dfb72cc1-d526-40e8-9982-077bb22fc44c.png)![FLOOR](https://user-images.githubusercontent.com/7338312/148694075-7ca93668-2ce6-4e26-af9c-1dc032bf6980.png)

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

## Premium

If you are interested in a ticker that isnt on this list, you can host your own using the code here or pay to have them made for you.

Price per bot (paid monthly): $1

Price per bot (paid yearly):  $10

If you are interested please see the [contact info on my github page](https://github.com/rssnyder) and send me a messgae via your platform of choice (discord perferred). For a live demo, join the support discord linked at the top or bottom of this page.

## Self-Hosting

[Click here](https://youtu.be/LhgCdtE8kmc) to watch a quick video tutorial on how to self-host these bots on linux. There is also an in depth written format [gist here](https://gist.github.com/rssnyder/55eb4e0b18cca399592a557e95b5547b). If you are familar with ansible, I have a [playbook here](https://github.com/rssnyder/isengard/blob/master/playbooks/discord-stock-ticker.yml).

Pull down the latest release for your OS [here](https://github.com/rssnyder/discord-stock-ticker/releases). Extract. Run.

```shell
wget https://github.com/rssnyder/discord-stock-ticker/releases/download/v2.0.0/discord-stock-ticker-v3.3.0-linux-amd64.tar.gz

tar zxf discord-stock-ticker-v3.3.0-linux-amd64.tar.gz

./discord-stock-ticker
```

### Setting options

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

### Systemd (linux)

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

If you need to make modifications to the setting of the service, just edit the `/etc/systemd/system/discord-stock-ticker.service` file on the line with `ExecStart=`. An example walkthrough can be found in [this issue](https://github.com/rssnyder/discord-stock-ticker/issues/137). Be sure to run `systemctl daemon-reload` to pick up and changes.

Now that you have the service running, you can add bots using the API exposed on the addres and port that the service runs on (this address is shown when you start the service).

## Managing bots

All bots are controlled via an API interface and follow the same api template for management:

Available methods:
  
```text
GET     # show all currently running bots and their configuration
POST    # create a new bot
PATCH   # restart a running bot
DELETE  # delete a running bot
```

If you are new to using an API to manage things, there are several ways to make API calls:

1) Curl. This is a command available on virtually all Linux distros. Replace anything between < and > with the appropriate information.

The generic format for a curl API call:

```shell
curl -X <method> -H "Content-type: application/json" -d <inline json or from file> <hostname>:<port>/<bot type>
```

GET is the default method for curl, so you may omit the method. Also since you're just retrieving your bots, you can omit the -d flag as well.

Get a listing of all your bots:

```shell
curl localhost:8080/<bot type>
```

Create a new bot:
(In this example, the bot configuration is located in a file 'btc.json', in the folder bots/crypo)

```shell
curl -X POST -H "Content-type: application/json" -d @bots/crypto/btc.json localhost:8080/ticker
```

Instructions for restarting running bots and deleting bots are forthcoming.

2) Powershell:
  
```shell
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

3) [postman](https://www.postman.com/)

## Stock and Crypto Price Tickers

bot type: `ticker`

Tracks stock or crypto prices. Uses Yahoo for stock or CoinGecko for crypto.

### Bot Configuration (stock)

```json
{
  "ticker": "pfg",                                  # string: symbol for the stock from yahoo finance
  "name": "2) PFG",                                 # string/OPTIONAL: overwrites display name of bot
  "color": true,                                    # bool/OPTIONAL: requires nickname
  "decorator": "@",                                 # string/OPTIONAL: what to show instead of arrows
  "currency": "aud",                                # string/OPTIONAL: alternative curreny
  "activity": "Hello;Its;Me",                       # string/OPTIONAL: list of strings to show in activity section
  "nickname": true,                                 # bool/OPTIONAL: display information in nickname vs activity
  "frequency": 10,                                  # int/OPTIONAL: seconds between refresh
  "twelve_data_key": "xxx",                         # string/OPTIONAL: use twelve data as source, pass in api key
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxx"   # string: dicord bot token
}
```

### Bot Configuration (crypto)

```json
{
  "name": "bitcoin",                                # string: name of the crypto from coingecko
  "crypto": true,                                   # bool: always true for crypto
  "ticker": "1) BTC",                               # string/OPTIONAL: overwrites display name of bot
  "color": true,                                    # bool/OPTIONAL: requires nickname
  "decorator": "@",                                 # string/OPTIONAL: what to show instead of arrows
  "currency": "aud",                                # string/OPTIONAL: alternative curreny
  "currency_symbol": "AUD",                         # string/OPTIONAL: alternative curreny symbol
  "pair": "binancecoin",                            # string/OPTIONAL: pair the coin with another coin, replaces activity section
  "pair_flip": true,                                # bool/OPTIONAL: show <pair>/<coin> rather than <coin>/<pair>
  "activity": "Hello;Its;Me",                       # string/OPTIONAL: list of strings to show in activity section
  "decimals": 3,                                    # int/OPTIONAL: set number of decimal places
  "nickname": true,                                 # bool/OPTIONAL: display information in nickname vs activity
  "frequency": 10,                                  # int/OPTIONAL: seconds between refresh
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxx"   # string: dicord bot token
}
```

## Stock and Crypto Price Tickerboards

bot type: `tickerboard`

Tracks multiple stock or crypto prices. Uses Yahoo for stock or CoinGecko for crypto.

### Bot Configuration (stock)

```json
{
  "name": "Stocks",                                 # string: name of your board
  "items": ["PFG", "GME", "AMC"],                   # list of strings: symbols from yahoo finance to rotate through
  "header": "1. ",                                  # string/OPTIONAL: adds a header to the nickname to help sort bots
  "color": true,                                    # bool/OPTIONAL: requires nickname
  "arrows": true,                                   # bool/OPTIONAL: show arrows in ticker names
  "nickname": true,                                 # bool/OPTIONAL: display information in nickname vs activity
  "frequency": 10,                                  # int/OPTIONAL: seconds between refresh
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxx"   # string: dicord bot token
}
```

### Bot Configuration (crypto)

```json
{
  "name": "Cryptos",                                # string: name of your board
  "crypto": true,                                   # bool: always true for crypto
  "items": ["bitcoin", "ethereum", "dogecoin"],     # list of strings: names from coingecko to rotate through
  "header": "2. ",                                  # string/OPTIONAL: adds a header to the nickname to help sort bots
  "color": true,                                    # bool/OPTIONAL: requires nickname
  "arrows": true,                                   # bool/OPTIONAL: show arrows in ticker names
  "nickname": true,                                 # bool/OPTIONAL: display information in nickname vs activity
  "frequency": 10,                                  # int/OPTIONAL: seconds between refresh
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxx"   # string: dicord bot token
}
```

## Crypto Market Cap

bot type: `marketcap`

Tracks the marketcap of a coin. Uses CoinGecko for data.

### Bot Configuration

```json
{
  "name": "bitcoin",                                # string: name of the crypto from coingecko
  "ticker": "1) BTC",                               # string/OPTIONAL: overwrites display name of bot
  "color": true,                                    # bool/OPTIONAL: requires nickname
  "decorator": "@",                                 # string/OPTIONAL: what to show instead of arrows
  "currency": "aud",                                # string/OPTIONAL: alternative curreny
  "currency_symbol": "AUD",                         # string/OPTIONAL: alternative curreny symbol
  "activity": "Hello;Its;Me",                       # string/OPTIONAL: list of strings to show in activity section
  "decimals": 3,                                    # int/OPTIONAL: set number of decimal places
  "nickname": true,                                 # bool/OPTIONAL: display information in nickname vs activity
  "frequency": 10,                                  # int/OPTIONAL: seconds between refresh
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxx"   # string: dicord bot token
}
```

## Crypto Circulating Supply

bot type: `circulating`

Tracks the circulating supply of a coin. Uses CoinGecko for data.

### Bot Configuration

```json
{
  "name": "bitcoin",                                # string: name of the crypto from coingecko
  "ticker": "1) BTC",                               # string/OPTIONAL: overwrites display name of bot
  "currency_symbol": "BITCOIN",                     # string/OPTIONAL: alternative curreny symbol
  "activity": "Hello;Its;Me",                       # string/OPTIONAL: list of strings to show in activity section
  "decimals": 3,                                    # int/OPTIONAL: set number of decimal places
  "nickname": true,                                 # bool/OPTIONAL: display information in nickname vs activity
  "frequency": 10,                                  # int/OPTIONAL: seconds between refresh
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxx"   # string: dicord bot token
}
```

## Crypto Total Value Locked

bot type: `valuelocked`

Tracks the total value locked of a coin. Uses CoinGecko for data.

### Bot Configuration

```json
{
  "name": "bitcoin",                                # string: name of the crypto from coingecko
  "ticker": "1) BTC",                               # string/OPTIONAL: overwrites display name of bot
  "currency": "aud",                                # string/OPTIONAL: alternative curreny
  "currency_symbol": "AUD",                         # string/OPTIONAL: alternative curreny symbol
  "activity": "Hello;Its;Me",                       # string/OPTIONAL: list of strings to show in activity section
  "decimals": 3,                                    # int/OPTIONAL: set number of decimal places
  "nickname": true,                                 # bool/OPTIONAL: display information in nickname vs activity
  "frequency": 10,                                  # int/OPTIONAL: seconds between refresh
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxx"   # string: dicord bot token
}
```

## Gas Prices

bot type: `gas`

![image](https://user-images.githubusercontent.com/7338312/127577601-43500287-1cf4-47ee-9f21-67c22f606850.png)

Track the gas price on:
- Ethereum
- Binance
- Polygon
- ..and many more

Uses [Zapper](https://api.zapper.fi/api/static/index.html#/Miscellaneous%20Data%20Endpoints/GasPriceController_getGasPrice) for data. For now always uses the eip1559 chains.

### Bot Configuration

```json
{
  "network": "ethereum",                            # string: one of: ethereum, binance-smart-chain, or polygon
  "nickname": true,                                 # bool/OPTIONAL: display information in nickname vs activity
  "frequency": 10,                                  # int/OPTIONAL: seconds between refresh
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxx"   # string: dicord bot token
}
```

## Ethereum, BSC, or Polygon Token Holders

bot type: `holders`

![HOLDERS](https://user-images.githubusercontent.com/7338312/126001392-dfb72cc1-d526-40e8-9982-077bb22fc44c.png)

Track the number of token holders on Ethereum or Binance chains. Uses etherscan or bscscan for data.

### Bot Configuration

```json
{
  "network": "ethereum",                            # string: one of: ethereum, binance-smart-chain, or polygon
  "address": "0x00000000000000000000000000",        # string: address of contract for token
  "activity": "ethereum",                           # string: text to show in activity section of the bot
  "nickname": true,                                 # bool/OPTIONAL: display information in nickname vs activity
  "frequency": 10,                                  # int/OPTIONAL: seconds between refresh
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxx"   # string: dicord bot token
}
```

## ETH/BSC/MATIC Token Price

bot type: `token`

Track the price of a token on Ethereum, Binance, or Polygon chains. Uses 1inch by default, or pancakeswap/dexlab if specified.

### Bot Configuration

```json
{
  "network": "ethereum",                            # string: network of token, options are ethereum, binance-smart-chain, or polygon
  "name": "my token",                               # string: display name of token
  "contract": "0x00000",                            # string: contract address of token
  "currency": "0x00000",                            # string/OPTIONAL: contract address of token to price against, default is USDC
  "nickname": true,                                 # bool/OPTIONAL: display information in nickname vs activity
  "color": true,                                    # bool/OPTIONAL: requires nickname
  "decorator": "@",                                 # string/OPTIONAL: what to show instead of arrows
  "activity": "Hello;Its;Me",                       # string/OPTIONAL: list of strings to show in activity section
  "source": "pancakeswap",                          # string/OPTIONAL: if the token is a BSC token, you can set pancakeswap here to use it vs 1inch; you can also set dexlab for solana tokens
  "frequency": 10,                                  # int/OPTIONAL: seconds between refresh
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxx"   # string: dicord bot token
}
```

## OpenSea/Solanart NFT Collection Floor Price

bot type: `floor`

![image](https://user-images.githubusercontent.com/7338312/148694081-9e90f75d-dcef-4a68-b98a-11c37b2b616a.png)

Track the floor price of an NFT collection on OpenSea or Solanart.

### Bot Configuration

```json
{
  "marketplace": "opensea",                         # string: one of: opensea, solsea or solanart
  "name": "ethereum",                               # string: collection name/id from source
  "nickname": true,                                 # bool/OPTIONAL: display information in nickname vs activity
  "frequency": 10,                                  # int/OPTIONAL: seconds between refresh
  "discord_bot_token": "xxxxxxxxxxxxxxxxxxxxxxxx"   # string: dicord bot token
}
```

## Roles for colors

To enable color changing you will need to create three roles.

The first role is the role the tickers will appear under. It can be named _anything you want_. You need to check the **Display role members seperatly from other online members** option for this role, but _do not_ assign a custom color for this role, leave it default.

Then you need to make two other roles. These roles need to be named _exactly_ **tickers-red** & **tickers-green**. **Do not** check the Display role members seperatly from other online members option for these roles, but do assign colors to these roles, red and green (or whatever color you want to represent gain/loss) respectively.

The last two roles tickers-green and tickers-red need to be below the first role in the role list in your server settings. You should then add all your ticker bots to the first role.

![roles example](https://user-images.githubusercontent.com/7338312/131678207-b1510955-f762-46e3-ae5c-1b5eddb68844.jpg)

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
