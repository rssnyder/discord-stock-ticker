# discord-stock-ticker

Live stock and crypto tickers for your discord server.

Now with five different types of tickers!

🍾 400+ public tickers with over 15k installs across 3k discord servers!

*Are you just looking to add free tickers to your discord server? Click the discord icon below to join the support server and get the list of avalible bots!*

[![Releases](https://github.com/rssnyder/discord-stock-ticker/workflows/Build%20and%20Publish%20Container%20Image/badge.svg)](https://github.com/rssnyder/discord-stock-ticker/releases)
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
    - [Other (not my) crypto discord bots](#other-not-my-crypto-discord-bots)
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

[![aa](https://img.shields.io/badge/stock-aa-ff69b4)](https://discord.com/oauth2/authorize?client_id=819316554486251530&permissions=0&scope=bot)
[![arch](https://img.shields.io/badge/stock-arch-ff69b4)](https://discord.com/oauth2/authorize?client_id=819319796189233154&permissions=0&scope=bot)
[![lac.to](https://img.shields.io/badge/stock-lac.to-ff69b4)](https://discord.com/oauth2/authorize?client_id=819320002578481203&permissions=0&scope=bot)
[![x](https://img.shields.io/badge/stock-x-ff69b4)](https://discord.com/oauth2/authorize?client_id=819320113945509899&permissions=0&scope=bot)
[![scr.to](https://img.shields.io/badge/stock-scr.to-ff69b4)](https://discord.com/oauth2/authorize?client_id=819320277368307795&permissions=0&scope=bot)
[![bb](https://img.shields.io/badge/stock-bb-ff69b4)](https://discord.com/oauth2/authorize?client_id=805289769272999986&permissions=0&scope=bot)
[![amc](https://img.shields.io/badge/stock-amc-ff69b4)](https://discord.com/oauth2/authorize?client_id=805294017441038357&permissions=0&scope=bot)
[![nok](https://img.shields.io/badge/stock-nok-ff69b4)](https://discord.com/oauth2/authorize?client_id=805294107962245120&permissions=0&scope=bot)
[![pfg](https://img.shields.io/badge/stock-pfg-ff69b4)](https://discord.com/oauth2/authorize?client_id=805466470930055189&permissions=0&scope=bot)
[![aapl](https://img.shields.io/badge/stock-aapl-ff69b4)](https://discord.com/oauth2/authorize?client_id=806569145184550922&permissions=0&scope=bot)
[![amzn](https://img.shields.io/badge/stock-amzn-ff69b4)](https://discord.com/oauth2/authorize?client_id=806570287042002945&permissions=0&scope=bot)
[![goog](https://img.shields.io/badge/stock-goog-ff69b4)](https://discord.com/oauth2/authorize?client_id=806570628156882945&permissions=0&scope=bot)
[![gspc](https://img.shields.io/badge/stock-gspc-ff69b4)](https://discord.com/oauth2/authorize?client_id=808431853363134484&permissions=0&scope=bot)
[![dji](https://img.shields.io/badge/stock-dji-ff69b4)](https://discord.com/oauth2/authorize?client_id=808432655746072596&permissions=0&scope=bot)
[![ixic](https://img.shields.io/badge/stock-ixic-ff69b4)](https://discord.com/oauth2/authorize?client_id=808432811644026921&permissions=0&scope=bot)
[![tsla](https://img.shields.io/badge/stock-tsla-ff69b4)](https://discord.com/oauth2/authorize?client_id=808723743069306882&permissions=0&scope=bot)
[![dkng](https://img.shields.io/badge/stock-dkng-ff69b4)](https://discord.com/oauth2/authorize?client_id=808724381608968202&permissions=0&scope=bot)
[![spy](https://img.shields.io/badge/stock-spy-ff69b4)](https://discord.com/oauth2/authorize?client_id=811418568846737500&permissions=0&scope=bot)
[![sprq](https://img.shields.io/badge/stock-sprq-ff69b4)](https://discord.com/oauth2/authorize?client_id=812041982980521995&permissions=0&scope=bot)
[![acic](https://img.shields.io/badge/stock-acic-ff69b4)](https://discord.com/oauth2/authorize?client_id=812041922499706890&permissions=0&scope=bot)
[![bmbl](https://img.shields.io/badge/stock-bmbl-ff69b4)](https://discord.com/oauth2/authorize?client_id=812041842903482378&permissions=0&scope=bot)
[![plug](https://img.shields.io/badge/stock-plug-ff69b4)](https://discord.com/oauth2/authorize?client_id=812041750528000031&permissions=0&scope=bot)
[![fcel](https://img.shields.io/badge/stock-fcel-ff69b4)](https://discord.com/oauth2/authorize?client_id=812041645008093186&permissions=0&scope=bot)
[![ipod](https://img.shields.io/badge/stock-ipod-ff69b4)](https://discord.com/oauth2/authorize?client_id=814168658082791454&permissions=0&scope=bot)
[![ipof](https://img.shields.io/badge/stock-ipof-ff69b4)](https://discord.com/oauth2/authorize?client_id=814169320233369640&permissions=0&scope=bot)
[![amd](https://img.shields.io/badge/stock-amd-ff69b4)](https://discord.com/oauth2/authorize?client_id=816049122850897923&permissions=0&scope=bot)
[![nio](https://img.shields.io/badge/stock-nio-ff69b4)](https://discord.com/oauth2/authorize?client_id=816049780546994196&permissions=0&scope=bot)
[![esgc](https://img.shields.io/badge/stock-esgc-ff69b4)](https://discord.com/oauth2/authorize?client_id=816372064494551141&permissions=0&scope=bot)
[![gc=f](https://img.shields.io/badge/stock-gc=f-ff69b4)](https://discord.com/oauth2/authorize?client_id=816375134122147850&permissions=0&scope=bot)
[![si=f](https://img.shields.io/badge/stock-si=f-ff69b4)](https://discord.com/oauth2/authorize?client_id=816375586066661456&permissions=0&scope=bot)
[![cl=f](https://img.shields.io/badge/stock-cl=f-ff69b4)](https://discord.com/oauth2/authorize?client_id=816375780442636288&permissions=0&scope=bot)
[![pltr](https://img.shields.io/badge/stock-pltr-ff69b4)](https://discord.com/oauth2/authorize?client_id=818471352415551518&permissions=0&scope=bot)
[![qqqj](https://img.shields.io/badge/stock-qqqj-ff69b4)](https://discord.com/oauth2/authorize?client_id=819320358045351937&permissions=0&scope=bot)
[![pypl](https://img.shields.io/badge/stock-pypl-ff69b4)](https://discord.com/oauth2/authorize?client_id=819395115465572352&permissions=0&scope=bot)
[![apha](https://img.shields.io/badge/stock-apha-ff69b4)](https://discord.com/oauth2/authorize?client_id=819395294949015552&permissions=0&scope=bot)
[![sndl](https://img.shields.io/badge/stock-sndl-ff69b4)](https://discord.com/oauth2/authorize?client_id=819638060894388305&permissions=0&scope=bot)
[![hut.to](https://img.shields.io/badge/stock-hut.to-ff69b4)](https://discord.com/oauth2/authorize?client_id=819639798482075678&permissions=0&scope=bot)
[![nhic](https://img.shields.io/badge/stock-nhic-ff69b4)](https://discord.com/oauth2/authorize?client_id=819639927146676304&permissions=0&scope=bot)
[![rty=f](https://img.shields.io/badge/stock-rty=f-ff69b4)](https://discord.com/oauth2/authorize?client_id=819578053881102336&permissions=0&scope=bot)
[![vix](https://img.shields.io/badge/stock-vix-ff69b4)](https://discord.com/oauth2/authorize?client_id=819576744176386048&permissions=0&scope=bot)
[![tnx](https://img.shields.io/badge/stock-tnx-ff69b4)](https://discord.com/oauth2/authorize?client_id=819576354361704459&permissions=0&scope=bot)
[![cciv](https://img.shields.io/badge/stock-cciv-ff69b4)](https://discord.com/oauth2/authorize?client_id=819575244691865623&permissions=0&scope=bot)
[![hcmc](https://img.shields.io/badge/stock-hcmc-ff69b4)](https://discord.com/oauth2/authorize?client_id=819575463957102593&permissions=0&scope=bot)
[![ctrm](https://img.shields.io/badge/stock-ctrm-ff69b4)](https://discord.com/oauth2/authorize?client_id=819939407754690572&permissions=0&scope=bot)
[![etfm](https://img.shields.io/badge/stock-etfm-ff69b4)](https://discord.com/oauth2/authorize?client_id=819939500687360071&permissions=0&scope=bot)
[![sens](https://img.shields.io/badge/stock-sens-ff69b4)](https://discord.com/oauth2/authorize?client_id=819939795580354620&permissions=0&scope=bot)
[![ftcv](https://img.shields.io/badge/stock-ftcv-ff69b4)](https://discord.com/oauth2/authorize?client_id=820048409016533015&permissions=0&scope=bot)
[![mvis](https://img.shields.io/badge/stock-mvis-ff69b4)](https://discord.com/oauth2/authorize?client_id=820048636649144330&permissions=0&scope=bot)
[![cake](https://img.shields.io/badge/stock-cake-ff69b4)](https://discord.com/oauth2/authorize?client_id=820048736201605141&permissions=0&scope=bot)
[![chwy](https://img.shields.io/badge/stock-chwy-ff69b4)](https://discord.com/oauth2/authorize?client_id=826512272506880071&permissions=0&scope=bot)
[![cmcsa](https://img.shields.io/badge/stock-cmcsa-ff69b4)](https://discord.com/oauth2/authorize?client_id=826512493953417226&permissions=0&scope=bot)
[![mstr](https://img.shields.io/badge/stock-mstr-ff69b4)](https://discord.com/oauth2/authorize?client_id=826512588078710794&permissions=0&scope=bot)
[![ethe](https://img.shields.io/badge/stock-ethe-ff69b4)](https://discord.com/oauth2/authorize?client_id=826512680005140550&permissions=0&scope=bot)
[![arkk](https://img.shields.io/badge/stock-arkk-ff69b4)](https://discord.com/oauth2/authorize?client_id=826512765580869673&permissions=0&scope=bot)
[![discb](https://img.shields.io/badge/stock-discb-ff69b4)](https://discord.com/oauth2/authorize?client_id=827236022466117642&permissions=0&scope=bot)
[![nvta](https://img.shields.io/badge/stock-nvta-ff69b4)](https://discord.com/oauth2/authorize?client_id=827236111440674876&permissions=0&scope=bot)
[![msft](https://img.shields.io/badge/stock-msft-ff69b4)](https://discord.com/oauth2/authorize?client_id=827236206575222794&permissions=0&scope=bot)
[![nflx](https://img.shields.io/badge/stock-nflx-ff69b4)](https://discord.com/oauth2/authorize?client_id=827236295323025408&permissions=0&scope=bot)
[![fcx](https://img.shields.io/badge/stock-fcx-ff69b4)](https://discord.com/oauth2/authorize?client_id=827986004618117170&permissions=0&scope=bot)
[![bidu](https://img.shields.io/badge/stock-bidu-ff69b4)](https://discord.com/oauth2/authorize?client_id=827986153147858946&permissions=0&scope=bot)
[![ttd](https://img.shields.io/badge/stock-ttd-ff69b4)](https://discord.com/oauth2/authorize?client_id=828326607411871754&permissions=0&scope=bot)
[![gme](https://img.shields.io/badge/stock-gme-ff69b4)](https://discord.com/oauth2/authorize?client_id=828417475624435712&permissions=0&scope=bot)
[![scr](https://img.shields.io/badge/stock-scr-ff69b4)](https://discord.com/oauth2/authorize?client_id=828613724608135259&permissions=0&scope=bot)
[![fubo](https://img.shields.io/badge/stock-fubo-ff69b4)](https://discord.com/oauth2/authorize?client_id=828613871823618128&permissions=0&scope=bot)
[![dis](https://img.shields.io/badge/stock-dis-ff69b4)](https://discord.com/oauth2/authorize?client_id=828613902740357121&permissions=0&scope=bot)
[![buzz](https://img.shields.io/badge/stock-buzz-ff69b4)](https://discord.com/oauth2/authorize?client_id=829119893714108477&permissions=0&scope=bot)
[![bngo](https://img.shields.io/badge/stock-bngo-ff69b4)](https://discord.com/oauth2/authorize?client_id=831958131214385152&permissions=0&scope=bot)
[![uavs](https://img.shields.io/badge/stock-uavs-ff69b4)](https://discord.com/oauth2/authorize?client_id=831958814419845130&permissions=0&scope=bot)
[![arkg](https://img.shields.io/badge/stock-arkg-ff69b4)](https://discord.com/oauth2/authorize?client_id=831958893801897994&permissions=0&scope=bot)
[![es=f](https://img.shields.io/badge/stock-es=f-ff69b4)](https://discord.com/oauth2/authorize?client_id=831958978455535667&permissions=0&scope=bot)
[![nq=f](https://img.shields.io/badge/stock-nq=f-ff69b4)](https://discord.com/oauth2/authorize?client_id=833759608861622353&permissions=0&scope=bot)
[![ym=f](https://img.shields.io/badge/stock-ym=f-ff69b4)](https://discord.com/oauth2/authorize?client_id=833759696783409195&permissions=0&scope=bot)
[![arkx](https://img.shields.io/badge/stock-arkx-ff69b4)](https://discord.com/oauth2/authorize?client_id=833796890092109834&permissions=0&scope=bot)
[![arkw](https://img.shields.io/badge/stock-arkw-ff69b4)](https://discord.com/oauth2/authorize?client_id=834494803244220456&permissions=0&scope=bot)
[![arkf](https://img.shields.io/badge/stock-arkf-ff69b4)](https://discord.com/oauth2/authorize?client_id=834494886693568572&permissions=0&scope=bot)
[![v](https://img.shields.io/badge/stock-v-ff69b4)](https://discord.com/oauth2/authorize?client_id=834494985092202537&permissions=0&scope=bot)
[![sq](https://img.shields.io/badge/stock-sq-ff69b4)](https://discord.com/oauth2/authorize?client_id=834543616314966036&permissions=0&scope=bot)
[![jpm](https://img.shields.io/badge/stock-jpm-ff69b4)](https://discord.com/oauth2/authorize?client_id=834543707314192444&permissions=0&scope=bot)
[![ma](https://img.shields.io/badge/stock-ma-ff69b4)](https://discord.com/oauth2/authorize?client_id=834543958276964362&permissions=0&scope=bot)
[![bac](https://img.shields.io/badge/stock-bac-ff69b4)](https://discord.com/oauth2/authorize?client_id=834544024055840768&permissions=0&scope=bot)
[![wfc](https://img.shields.io/badge/stock-wfc-ff69b4)](https://discord.com/oauth2/authorize?client_id=834544089621200967&permissions=0&scope=bot)
[![zb=f](https://img.shields.io/badge/stock-zb=f-ff69b4)](https://discord.com/oauth2/authorize?client_id=834544178456952912&permissions=0&scope=bot)
[![crl](https://img.shields.io/badge/stock-crl-ff69b4)](https://discord.com/oauth2/authorize?client_id=836339309602144257&permissions=0&scope=bot)
[![mnmd](https://img.shields.io/badge/stock-mnmd-ff69b4)](https://discord.com/oauth2/authorize?client_id=836339349133066241&permissions=0&scope=bot)
[![c](https://img.shields.io/badge/stock-c-ff69b4)](https://discord.com/oauth2/authorize?client_id=836385880175673415&permissions=0&scope=bot)
[![nvda](https://img.shields.io/badge/stock-nvda-ff69b4)](https://discord.com/oauth2/authorize?client_id=836385897372581910&permissions=0&scope=bot)
[![comp](https://img.shields.io/badge/stock-comp-ff69b4)](https://discord.com/oauth2/authorize?client_id=836385913604669440&permissions=0&scope=bot)
[![nsei](https://img.shields.io/badge/stock-nsei-ff69b4)](https://discord.com/oauth2/authorize?client_id=836385929560195083&permissions=0&scope=bot)
[![brka](https://img.shields.io/badge/stock-brka-ff69b4)](https://discord.com/oauth2/authorize?client_id=836385946970750988&permissions=0&scope=bot)
[![nsebank](https://img.shields.io/badge/stock-nsebank-ff69b4)](https://discord.com/oauth2/authorize?client_id=836385970785615923&permissions=0&scope=bot)
[![fb](https://img.shields.io/badge/stock-fb-ff69b4)](https://discord.com/oauth2/authorize?client_id=836385990113361951&permissions=0&scope=bot)
[![sol](https://img.shields.io/badge/stock-sol-ff69b4)](https://discord.com/oauth2/authorize?client_id=836386005493612544&permissions=0&scope=bot)
[![sos](https://img.shields.io/badge/stock-sos-ff69b4)](https://discord.com/oauth2/authorize?client_id=837015771053686831&permissions=0&scope=bot)
[![mara](https://img.shields.io/badge/stock-mara-ff69b4)](https://discord.com/oauth2/authorize?client_id=837015831794679848&permissions=0&scope=bot)
[![ebon](https://img.shields.io/badge/stock-ebon-ff69b4)](https://discord.com/oauth2/authorize?client_id=838824558647705681&permissions=0&scope=bot)
[![riot](https://img.shields.io/badge/stock-riot-ff69b4)](https://discord.com/oauth2/authorize?client_id=838824632537710643&permissions=0&scope=bot)
[![ocgn](https://img.shields.io/badge/stock-ocgn-ff69b4)](https://discord.com/oauth2/authorize?client_id=839201266520555562&permissions=0&scope=bot)
[![bfarf](https://img.shields.io/badge/stock-bfarf-ff69b4)](https://discord.com/oauth2/authorize?client_id=839288940665372694&permissions=0&scope=bot)
[![hutmf](https://img.shields.io/badge/stock-hutmf-ff69b4)](https://discord.com/oauth2/authorize?client_id=839289012883947550&permissions=0&scope=bot)
[![rkt](https://img.shields.io/badge/stock-rkt-ff69b4)](https://discord.com/oauth2/authorize?client_id=839289069045547018&permissions=0&scope=bot)
[![btc](https://img.shields.io/badge/stock-btc-ff69b4)](https://discord.com/oauth2/authorize?client_id=839289136418783322&permissions=0&scope=bot)
[![tlry](https://img.shields.io/badge/stock-tlry-ff69b4)](https://discord.com/oauth2/authorize?client_id=840022914690842624&permissions=0&scope=bot)
[![ogi](https://img.shields.io/badge/stock-ogi-ff69b4)](https://discord.com/oauth2/authorize?client_id=840023145758851103&permissions=0&scope=bot)
[![cidm](https://img.shields.io/badge/stock-cidm-ff69b4)](https://discord.com/oauth2/authorize?client_id=840023236640505867&permissions=0&scope=bot)
[![roku](https://img.shields.io/badge/stock-roku-ff69b4)](https://discord.com/oauth2/authorize?client_id=840023321034883072&permissions=0&scope=bot)
[![btbt](https://img.shields.io/badge/stock-btbt-ff69b4)](https://discord.com/oauth2/authorize?client_id=841850350226898975&permissions=0&scope=bot)
[![omx](https://img.shields.io/badge/stock-omx-ff69b4)](https://discord.com/oauth2/authorize?client_id=841850514429444158&permissions=0&scope=bot)
[![tcnnf](https://img.shields.io/badge/stock-tcnnf-ff69b4)](https://discord.com/oauth2/authorize?client_id=841412636792979526&permissions=0&scope=bot)
[![trul](https://img.shields.io/badge/stock-trul-ff69b4)](https://discord.com/oauth2/authorize?client_id=841850589402628107&permissions=0&scope=bot)
[![adbe](https://img.shields.io/badge/stock-adbe-ff69b4)](https://discord.com/oauth2/authorize?client_id=844012640878723103&permissions=0&scope=bot)
[![docu](https://img.shields.io/badge/stock-docu-ff69b4)](https://discord.com/oauth2/authorize?client_id=844934496955400223&permissions=0&scope=bot)
[![lmt](https://img.shields.io/badge/stock-lmt-ff69b4)](https://discord.com/oauth2/authorize?client_id=844934534089015337&permissions=0&scope=bot)
[![wmt](https://img.shields.io/badge/stock-wmt-ff69b4)](https://discord.com/oauth2/authorize?client_id=844934568796618772&permissions=0&scope=bot)
[![cost](https://img.shields.io/badge/stock-cost-ff69b4)](https://discord.com/oauth2/authorize?client_id=844934604674039818&permissions=0&scope=bot)
[![nya](https://img.shields.io/badge/stock-nya-ff69b4)](https://discord.com/oauth2/authorize?client_id=844934639574974475&permissions=0&scope=bot)
[![gold](https://img.shields.io/badge/stock-gold-ff69b4)](https://discord.com/oauth2/authorize?client_id=844934740171948032&permissions=0&scope=bot)
[![ttt](https://img.shields.io/badge/stock-ttt-ff69b4)](https://discord.com/oauth2/authorize?client_id=846441735923957861&permissions=0&scope=bot)
[![spce](https://img.shields.io/badge/stock-spce-ff69b4)](https://discord.com/oauth2/authorize?client_id=846441819453784114&permissions=0&scope=bot)
[![hd](https://img.shields.io/badge/stock-hd-ff69b4)](https://discord.com/oauth2/authorize?client_id=846442063427272774&permissions=0&scope=bot)
[![iwm](https://img.shields.io/badge/stock-iwm-ff69b4)](https://discord.com/oauth2/authorize?client_id=848232028511272972&permissions=0&scope=bot)
[![nkla](https://img.shields.io/badge/stock-nkla-ff69b4)](https://discord.com/oauth2/authorize?client_id=848232142545616927&permissions=0&scope=bot)
[![sklz](https://img.shields.io/badge/stock-sklz-ff69b4)](https://discord.com/oauth2/authorize?client_id=848232593401053204&permissions=0&scope=bot)
[![coin](https://img.shields.io/badge/stock-coin-ff69b4)](https://discord.com/oauth2/authorize?client_id=848232686320484382&permissions=0&scope=bot)
[![snow](https://img.shields.io/badge/stock-snow-ff69b4)](https://discord.com/oauth2/authorize?client_id=848232784959111259&permissions=0&scope=bot)
[![li](https://img.shields.io/badge/stock-li-ff69b4)](https://discord.com/oauth2/authorize?client_id=850012125546938379&permissions=0&scope=bot)
[![ldos](https://img.shields.io/badge/stock-ldos-ff69b4)](https://discord.com/oauth2/authorize?client_id=850012308456341504&permissions=0&scope=bot)
[![goev](https://img.shields.io/badge/stock-goev-ff69b4)](https://discord.com/oauth2/authorize?client_id=850012384391069726&permissions=0&scope=bot)
[![ctxr](https://img.shields.io/badge/stock-ctxr-ff69b4)](https://discord.com/oauth2/authorize?client_id=850012459674894376&permissions=0&scope=bot)
[![ndaq](https://img.shields.io/badge/stock-ndaq-ff69b4)](https://discord.com/oauth2/authorize?client_id=851203267840966696&permissions=0&scope=bot)
[![wkhs](https://img.shields.io/badge/stock-wkhs-ff69b4)](https://discord.com/oauth2/authorize?client_id=851203379286769664&permissions=0&scope=bot)
[![clov](https://img.shields.io/badge/stock-clov-ff69b4)](https://discord.com/oauth2/authorize?client_id=851811767960600577&permissions=0&scope=bot)
[![vti](https://img.shields.io/badge/stock-vti-ff69b4)](https://discord.com/oauth2/authorize?client_id=851811845621678100&permissions=0&scope=bot)
[![fsr](https://img.shields.io/badge/stock-fsr-ff69b4)](https://discord.com/oauth2/authorize?client_id=854499151407742997&permissions=0&scope=bot)
[![qqq](https://img.shields.io/badge/stock-qqq-ff69b4)](https://discord.com/oauth2/authorize?client_id=854499452840312842&permissions=0&scope=bot)
[![es](https://img.shields.io/badge/stock-es-ff69b4)](https://discord.com/oauth2/authorize?client_id=854499554774876190&permissions=0&scope=bot)
[![si](https://img.shields.io/badge/stock-si-ff69b4)](https://discord.com/oauth2/authorize?client_id=856537356970885191&permissions=0&scope=bot)
[![recaf](https://img.shields.io/badge/stock-recaf-ff69b4)](https://discord.com/oauth2/authorize?client_id=856621703556104213&permissions=0&scope=bot)
[![reco.v](https://img.shields.io/badge/stock-reco.v-ff69b4)](https://discord.com/oauth2/authorize?client_id=856621799096844298&permissions=0&scope=bot)
[![xle](https://img.shields.io/badge/stock-xle-ff69b4)](https://discord.com/oauth2/authorize?client_id=856621874896699412&permissions=0&scope=bot)
[![baba](https://img.shields.io/badge/stock-baba-ff69b4)](https://discord.com/oauth2/authorize?client_id=856892723725729812&permissions=0&scope=bot)
[![2222.sr](https://img.shields.io/badge/stock-2222.sr-ff69b4)](https://discord.com/oauth2/authorize?client_id=856892815367471115&permissions=0&scope=bot)
[![tcehy](https://img.shields.io/badge/stock-tcehy-ff69b4)](https://discord.com/oauth2/authorize?client_id=856892893137731655&permissions=0&scope=bot)
[![ko](https://img.shields.io/badge/stock-ko-ff69b4)](https://discord.com/oauth2/authorize?client_id=856893055365939260&permissions=0&scope=bot)
[![wish](https://img.shields.io/badge/stock-wish-ff69b4)](https://discord.com/oauth2/authorize?client_id=857265820534308865&permissions=0&scope=bot)
[![reliance.ns](https://img.shields.io/badge/stock-reliance.ns-ff69b4)](https://discord.com/oauth2/authorize?client_id=857265911488708619&permissions=0&scope=bot)
[![vbiv](https://img.shields.io/badge/stock-vbiv-ff69b4)](https://discord.com/oauth2/authorize?client_id=857266064056647700&permissions=0&scope=bot)
[![xbi](https://img.shields.io/badge/stock-xbi-ff69b4)](https://discord.com/oauth2/authorize?client_id=857266136142708736&permissions=0&scope=bot)
[![cstm](https://img.shields.io/badge/stock-cstm-ff69b4)](https://discord.com/oauth2/authorize?client_id=857358104105320468&permissions=0&scope=bot)
[![upst](https://img.shields.io/badge/stock-upst-ff69b4)](https://discord.com/oauth2/authorize?client_id=858008502957375499&permissions=0&scope=bot)
[![rblx](https://img.shields.io/badge/stock-rblx-ff69b4)](https://discord.com/oauth2/authorize?client_id=858008578055995443&permissions=0&scope=bot)
[![abnb](https://img.shields.io/badge/stock-abnb-ff69b4)](https://discord.com/oauth2/authorize?client_id=858008652794167306&permissions=0&scope=bot)
[![stem](https://img.shields.io/badge/stock-stem-ff69b4)](https://discord.com/oauth2/authorize?client_id=858009557754183751&permissions=0&scope=bot)
[![atos](https://img.shields.io/badge/stock-atos-ff69b4)](https://discord.com/oauth2/authorize?client_id=859094401325989898&permissions=0&scope=bot)
[![mmat](https://img.shields.io/badge/stock-mmat-ff69b4)](https://discord.com/oauth2/authorize?client_id=859148180594622484&permissions=0&scope=bot)
[![dnut](https://img.shields.io/badge/stock-dnut-ff69b4)](https://discord.com/oauth2/authorize?client_id=860213362741477467&permissions=0&scope=bot)
[![upro](https://img.shields.io/badge/stock-upro-ff69b4)](https://discord.com/oauth2/authorize?client_id=860516151238066207&permissions=0&scope=bot)
[![tmf](https://img.shields.io/badge/stock-tmf-ff69b4)](https://discord.com/oauth2/authorize?client_id=860516222633508914&permissions=0&scope=bot)
[![tqqq](https://img.shields.io/badge/stock-tqqq-ff69b4)](https://discord.com/oauth2/authorize?client_id=860516468722630677&permissions=0&scope=bot)
[![shop](https://img.shields.io/badge/stock-shop-ff69b4)](https://discord.com/oauth2/authorize?client_id=862083201610678283&permissions=0&scope=bot)
[![indiavix](https://img.shields.io/badge/stock-indiavix-ff69b4)](https://discord.com/oauth2/authorize?client_id=862732016298622976&permissions=0&scope=bot)
[![aty](https://img.shields.io/badge/stock-aty-ff69b4)](https://discord.com/oauth2/authorize?client_id=862732160724369408&permissions=0&scope=bot)
[![at.to](https://img.shields.io/badge/stock-at.to-ff69b4)](https://discord.com/oauth2/authorize?client_id=862732228797399070&permissions=0&scope=bot)
[![hgen](https://img.shields.io/badge/stock-hgen-ff69b4)](https://discord.com/oauth2/authorize?client_id=864864511864995870&permissions=0&scope=bot)
[![inn21.si](https://img.shields.io/badge/stock-inn21.si-ff69b4)](https://discord.com/oauth2/authorize?client_id=864864678991495168&permissions=0&scope=bot)
[![nvax](https://img.shields.io/badge/stock-nvax-ff69b4)](https://discord.com/oauth2/authorize?client_id=867440167258161162&permissions=0&scope=bot)
[![eurusd=x](https://img.shields.io/badge/stock-eurusd=x-ff69b4)](https://discord.com/oauth2/authorize?client_id=869243628164354109&permissions=0&scope=bot)
[![gbpusd=x](https://img.shields.io/badge/stock-gbpusd=x-ff69b4)](https://discord.com/oauth2/authorize?client_id=869243765523644476&permissions=0&scope=bot)
[![audusd=x](https://img.shields.io/badge/stock-audusd=x-ff69b4)](https://discord.com/oauth2/authorize?client_id=869247986058817577&permissions=0&scope=bot)
[![cad=x](https://img.shields.io/badge/stock-cad=x-ff69b4)](https://discord.com/oauth2/authorize?client_id=869571837636534283&permissions=0&scope=bot)
[![nzdusd=x](https://img.shields.io/badge/stock-nzdusd=x-ff69b4)](https://discord.com/oauth2/authorize?client_id=869571906901278740&permissions=0&scope=bot)
[![dxy.nyb](https://img.shields.io/badge/stock-dxy.nyb-ff69b4)](https://discord.com/oauth2/authorize?client_id=869572043945963550&permissions=0&scope=bot)
[![chf=x](https://img.shields.io/badge/stock-chf=x-ff69b4)](https://discord.com/oauth2/authorize?client_id=867840575474892851&permissions=0&scope=bot)
[![jpy=x](https://img.shields.io/badge/stock-jpy=x-ff69b4)](https://discord.com/oauth2/authorize?client_id=869950362939969536&permissions=0&scope=bot)
[![ftse](https://img.shields.io/badge/stock-ftse-ff69b4)](https://discord.com/oauth2/authorize?client_id=869998805809000468&permissions=0&scope=bot)
[![hood](https://img.shields.io/badge/stock-hood-ff69b4)](https://discord.com/oauth2/authorize?client_id=869998928664338432&permissions=0&scope=bot)
[![ndx](https://img.shields.io/badge/stock-ndx-ff69b4)](https://discord.com/oauth2/authorize?client_id=874361894335569951&permissions=0&scope=bot)
[![slp](https://img.shields.io/badge/stock-slp-ff69b4)](https://discord.com/oauth2/authorize?client_id=877903126521339945&permissions=0&scope=bot)
[![psfe](https://img.shields.io/badge/stock-psfe-ff69b4)](https://discord.com/oauth2/authorize?client_id=877903459796529223&permissions=0&scope=bot)
[![tlt](https://img.shields.io/badge/stock-tlt-ff69b4)](https://discord.com/oauth2/authorize?client_id=880628730442174494&permissions=0&scope=bot)
[![polis](https://img.shields.io/badge/stock-polis-ff69b4)](https://discord.com/oauth2/authorize?client_id=&permissions=0&scope=bot)
[![alix](https://img.shields.io/badge/stock-alix-ff69b4)](https://discord.com/oauth2/authorize?client_id=&permissions=0&scope=bot)
[![poco](https://img.shields.io/badge/stock-poco-ff69b4)](https://discord.com/oauth2/authorize?client_id=&permissions=0&scope=bot)
[![ceek](https://img.shields.io/badge/stock-ceek-ff69b4)](https://discord.com/oauth2/authorize?client_id=&permissions=0&scope=bot)
[![smon](https://img.shields.io/badge/stock-smon-ff69b4)](https://discord.com/oauth2/authorize?client_id=&permissions=0&scope=bot)
[![imx](https://img.shields.io/badge/stock-imx-ff69b4)](https://discord.com/oauth2/authorize?client_id=&permissions=0&scope=bot)
[![xil](https://img.shields.io/badge/stock-xil-ff69b4)](https://discord.com/oauth2/authorize?client_id=&permissions=0&scope=bot)
[![moni](https://img.shields.io/badge/stock-moni-ff69b4)](https://discord.com/oauth2/authorize?client_id=&permissions=0&scope=bot)

### Crypto

 [![bitcoincash](https://img.shields.io/badge/crypto-bitcoincash-blueviolet)](https://discord.com/oauth2/authorize?client_id=805604560013230170&permissions=0&scope=bot)
[![ethereum](https://img.shields.io/badge/crypto-ethereum-blueviolet)](https://discord.com/oauth2/authorize?client_id=805605209522962452&permissions=0&scope=bot)
[![dogecoin](https://img.shields.io/badge/crypto-dogecoin-blueviolet)](https://discord.com/oauth2/authorize?client_id=805605888387186699&permissions=0&scope=bot)
[![monero](https://img.shields.io/badge/crypto-monero-blueviolet)](https://discord.com/oauth2/authorize?client_id=806282848045629451&permissions=0&scope=bot)
[![litecoin](https://img.shields.io/badge/crypto-litecoin-blueviolet)](https://discord.com/oauth2/authorize?client_id=806635240482668574&permissions=0&scope=bot)
[![ripple](https://img.shields.io/badge/crypto-ripple-blueviolet)](https://discord.com/oauth2/authorize?client_id=806634757168693258&permissions=0&scope=bot)
[![polkadot](https://img.shields.io/badge/crypto-polkadot-blueviolet)](https://discord.com/oauth2/authorize?client_id=806633568787890217&permissions=0&scope=bot)
[![cardano](https://img.shields.io/badge/crypto-cardano-blueviolet)](https://discord.com/oauth2/authorize?client_id=807311315055542272&permissions=0&scope=bot)
[![chainlink](https://img.shields.io/badge/crypto-chainlink-blueviolet)](https://discord.com/oauth2/authorize?client_id=808407486860230747&permissions=0&scope=bot)
[![stellar](https://img.shields.io/badge/crypto-stellar-blueviolet)](https://discord.com/oauth2/authorize?client_id=808409647731179534&permissions=0&scope=bot)
[![0x](https://img.shields.io/badge/crypto-0x-blueviolet)](https://discord.com/oauth2/authorize?client_id=810892119362895872&permissions=0&scope=bot)
[![balancer](https://img.shields.io/badge/crypto-balancer-blueviolet)](https://discord.com/oauth2/authorize?client_id=810894385360535572&permissions=0&scope=bot)
[![iota](https://img.shields.io/badge/crypto-iota-blueviolet)](https://discord.com/oauth2/authorize?client_id=814170376254652486&permissions=0&scope=bot)
[![reeffinance](https://img.shields.io/badge/crypto-reeffinance-blueviolet)](https://discord.com/oauth2/authorize?client_id=814288538107379742&permissions=0&scope=bot)
[![algorand](https://img.shields.io/badge/crypto-algorand-blueviolet)](https://discord.com/oauth2/authorize?client_id=819274628778164265&permissions=0&scope=bot)
[![tezos](https://img.shields.io/badge/crypto-tezos-blueviolet)](https://discord.com/oauth2/authorize?client_id=811609668991975484&permissions=0&scope=bot)
[![ethereumclassic](https://img.shields.io/badge/crypto-ethereumclassic-blueviolet)](https://discord.com/oauth2/authorize?client_id=819395405980762182&permissions=0&scope=bot)
[![ravencoin](https://img.shields.io/badge/crypto-ravencoin-blueviolet)](https://discord.com/oauth2/authorize?client_id=819395519708921866&permissions=0&scope=bot)
[![binancecoin](https://img.shields.io/badge/crypto-binancecoin-blueviolet)](https://discord.com/oauth2/authorize?client_id=819395643193688124&permissions=0&scope=bot)
[![ethernitychain](https://img.shields.io/badge/crypto-ethernitychain-blueviolet)](https://discord.com/oauth2/authorize?client_id=819939616349749249&permissions=0&scope=bot)
[![ecomi](https://img.shields.io/badge/crypto-ecomi-blueviolet)](https://discord.com/oauth2/authorize?client_id=819939716228579360&permissions=0&scope=bot)
[![reserverightstoken](https://img.shields.io/badge/crypto-reserverightstoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=820048829000581142&permissions=0&scope=bot)
[![aave](https://img.shields.io/badge/crypto-aave-blueviolet)](https://discord.com/oauth2/authorize?client_id=826512401502961756&permissions=0&scope=bot)
[![rulerprotocol](https://img.shields.io/badge/crypto-rulerprotocol-blueviolet)](https://discord.com/oauth2/authorize?client_id=827236401329340476&permissions=0&scope=bot)
[![polkamon](https://img.shields.io/badge/crypto-polkamon-blueviolet)](https://discord.com/oauth2/authorize?client_id=827984460859310080&permissions=0&scope=bot)
[![uniswap](https://img.shields.io/badge/crypto-uniswap-blueviolet)](https://discord.com/oauth2/authorize?client_id=827985872389275658&permissions=0&scope=bot)
[![bittorrent2](https://img.shields.io/badge/crypto-bittorrent2-blueviolet)](https://discord.com/oauth2/authorize?client_id=827986251264819201&permissions=0&scope=bot)
[![tron](https://img.shields.io/badge/crypto-tron-blueviolet)](https://discord.com/oauth2/authorize?client_id=828326036785463307&permissions=0&scope=bot)
[![vechain](https://img.shields.io/badge/crypto-vechain-blueviolet)](https://discord.com/oauth2/authorize?client_id=828326223306424350&permissions=0&scope=bot)
[![vethortoken](https://img.shields.io/badge/crypto-vethortoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=828326375911194635&permissions=0&scope=bot)
[![siacoin](https://img.shields.io/badge/crypto-siacoin-blueviolet)](https://discord.com/oauth2/authorize?client_id=828326519625613393&permissions=0&scope=bot)
[![bitcoin](https://img.shields.io/badge/crypto-bitcoin-blueviolet)](https://discord.com/oauth2/authorize?client_id=828417381779898368&permissions=0&scope=bot)
[![illuvium](https://img.shields.io/badge/crypto-illuvium-blueviolet)](https://discord.com/oauth2/authorize?client_id=828417571354968104&permissions=0&scope=bot)
[![cosmos](https://img.shields.io/badge/crypto-cosmos-blueviolet)](https://discord.com/oauth2/authorize?client_id=828417242570948638&permissions=0&scope=bot)
[![zilliqa](https://img.shields.io/badge/crypto-zilliqa-blueviolet)](https://discord.com/oauth2/authorize?client_id=828417678976745483&permissions=0&scope=bot)
[![pangolin](https://img.shields.io/badge/crypto-pangolin-blueviolet)](https://discord.com/oauth2/authorize?client_id=828613694845747260&permissions=0&scope=bot)
[![orionprotocol](https://img.shields.io/badge/crypto-orionprotocol-blueviolet)](https://discord.com/oauth2/authorize?client_id=828613759781961790&permissions=0&scope=bot)
[![maticnetwork](https://img.shields.io/badge/crypto-maticnetwork-blueviolet)](https://discord.com/oauth2/authorize?client_id=828613785345458206&permissions=0&scope=bot)
[![basicattentiontoken](https://img.shields.io/badge/crypto-basicattentiontoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=828613810355961898&permissions=0&scope=bot)
[![wink](https://img.shields.io/badge/crypto-wink-blueviolet)](https://discord.com/oauth2/authorize?client_id=828613846889136148&permissions=0&scope=bot)
[![shibainu](https://img.shields.io/badge/crypto-shibainu-blueviolet)](https://discord.com/oauth2/authorize?client_id=829119870556831816&permissions=0&scope=bot)
[![pancakeswaptoken](https://img.shields.io/badge/crypto-pancakeswaptoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=831957913819021322&permissions=0&scope=bot)
[![graphlinqprotocol](https://img.shields.io/badge/crypto-graphlinqprotocol-blueviolet)](https://discord.com/oauth2/authorize?client_id=831958048523419696&permissions=0&scope=bot)
[![solana](https://img.shields.io/badge/crypto-solana-blueviolet)](https://discord.com/oauth2/authorize?client_id=836339329084948490&permissions=0&scope=bot)
[![banano](https://img.shields.io/badge/crypto-banano-blueviolet)](https://discord.com/oauth2/authorize?client_id=836339375398453268&permissions=0&scope=bot)
[![raydium](https://img.shields.io/badge/crypto-raydium-blueviolet)](https://discord.com/oauth2/authorize?client_id=836385816409407488&permissions=0&scope=bot)
[![cope](https://img.shields.io/badge/crypto-cope-blueviolet)](https://discord.com/oauth2/authorize?client_id=836385861641830420&permissions=0&scope=bot)
[![safemoon](https://img.shields.io/badge/crypto-safemoon-blueviolet)](https://discord.com/oauth2/authorize?client_id=837015732919074878&permissions=0&scope=bot)
[![nervefinance](https://img.shields.io/badge/crypto-nervefinance-blueviolet)](https://discord.com/oauth2/authorize?client_id=837015750845530122&permissions=0&scope=bot)
[![lightningprotocol](https://img.shields.io/badge/crypto-lightningprotocol-blueviolet)](https://discord.com/oauth2/authorize?client_id=837015792464953347&permissions=0&scope=bot)
[![ftxtoken](https://img.shields.io/badge/crypto-ftxtoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=837015812677304380&permissions=0&scope=bot)
[![enjincoin](https://img.shields.io/badge/crypto-enjincoin-blueviolet)](https://discord.com/oauth2/authorize?client_id=838824692478771232&permissions=0&scope=bot)
[![quick](https://img.shields.io/badge/crypto-quick-blueviolet)](https://discord.com/oauth2/authorize?client_id=838824749052330044&permissions=0&scope=bot)
[![decentraland](https://img.shields.io/badge/crypto-decentraland-blueviolet)](https://discord.com/oauth2/authorize?client_id=840023058734776351&permissions=0&scope=bot)
[![fantom](https://img.shields.io/badge/crypto-fantom-blueviolet)](https://discord.com/oauth2/authorize?client_id=841412194578071612&permissions=0&scope=bot)
[![spookyswap](https://img.shields.io/badge/crypto-spookyswap-blueviolet)](https://discord.com/oauth2/authorize?client_id=841412279987470346&permissions=0&scope=bot)
[![apeswapfinance](https://img.shields.io/badge/crypto-apeswapfinance-blueviolet)](https://discord.com/oauth2/authorize?client_id=841412365585219625&permissions=0&scope=bot)
[![locgame](https://img.shields.io/badge/crypto-locgame-blueviolet)](https://discord.com/oauth2/authorize?client_id=841412468404518974&permissions=0&scope=bot)
[![coti](https://img.shields.io/badge/crypto-coti-blueviolet)](https://discord.com/oauth2/authorize?client_id=841412539854094416&permissions=0&scope=bot)
[![caspernetwork](https://img.shields.io/badge/crypto-caspernetwork-blueviolet)](https://discord.com/oauth2/authorize?client_id=841849975134355506&permissions=0&scope=bot)
[![hederahashgraph](https://img.shields.io/badge/crypto-hederahashgraph-blueviolet)](https://discord.com/oauth2/authorize?client_id=841850253418430495&permissions=0&scope=bot)
[![waultswap](https://img.shields.io/badge/crypto-waultswap-blueviolet)](https://discord.com/oauth2/authorize?client_id=841850425312935938&permissions=0&scope=bot)
[![rope](https://img.shields.io/badge/crypto-rope-blueviolet)](https://discord.com/oauth2/authorize?client_id=844012291304980490&permissions=0&scope=bot)
[![wootradenetwork](https://img.shields.io/badge/crypto-wootradenetwork-blueviolet)](https://discord.com/oauth2/authorize?client_id=844012368626188298&permissions=0&scope=bot)
[![sushi](https://img.shields.io/badge/crypto-sushi-blueviolet)](https://discord.com/oauth2/authorize?client_id=844012506670432277&permissions=0&scope=bot)
[![luksotoken](https://img.shields.io/badge/crypto-luksotoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=844012575615615007&permissions=0&scope=bot)
[![elevenfinance](https://img.shields.io/badge/crypto-elevenfinance-blueviolet)](https://discord.com/oauth2/authorize?client_id=844934678083403826&permissions=0&scope=bot)
[![kusama](https://img.shields.io/badge/crypto-kusama-blueviolet)](https://discord.com/oauth2/authorize?client_id=844934710597779456&permissions=0&scope=bot)
[![eos](https://img.shields.io/badge/crypto-eos-blueviolet)](https://discord.com/oauth2/authorize?client_id=844934770314969109&permissions=0&scope=bot)
[![moonstar](https://img.shields.io/badge/crypto-moonstar-blueviolet)](https://discord.com/oauth2/authorize?client_id=846441900910051378&permissions=0&scope=bot)
[![peacockcoin](https://img.shields.io/badge/crypto-peacockcoin-blueviolet)](https://discord.com/oauth2/authorize?client_id=846441985429733416&permissions=0&scope=bot)
[![esterfinance](https://img.shields.io/badge/crypto-esterfinance-blueviolet)](https://discord.com/oauth2/authorize?client_id=849375576941658143&permissions=0&scope=bot)
[![terraluna](https://img.shields.io/badge/crypto-terraluna-blueviolet)](https://discord.com/oauth2/authorize?client_id=849375907856384020&permissions=0&scope=bot)
[![piratechain](https://img.shields.io/badge/crypto-piratechain-blueviolet)](https://discord.com/oauth2/authorize?client_id=850012220473344020&permissions=0&scope=bot)
[![thegraph](https://img.shields.io/badge/crypto-thegraph-blueviolet)](https://discord.com/oauth2/authorize?client_id=850384459488297000&permissions=0&scope=bot)
[![rune](https://img.shields.io/badge/crypto-rune-blueviolet)](https://discord.com/oauth2/authorize?client_id=850384555210047509&permissions=0&scope=bot)
[![dfynnetwork](https://img.shields.io/badge/crypto-dfynnetwork-blueviolet)](https://discord.com/oauth2/authorize?client_id=850384639368495114&permissions=0&scope=bot)
[![celo](https://img.shields.io/badge/crypto-celo-blueviolet)](https://discord.com/oauth2/authorize?client_id=850384720084205637&permissions=0&scope=bot)
[![pussyfinancial](https://img.shields.io/badge/crypto-pussyfinancial-blueviolet)](https://discord.com/oauth2/authorize?client_id=851202946821652520&permissions=0&scope=bot)
[![irontitaniumtoken](https://img.shields.io/badge/crypto-irontitaniumtoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=851203079487619073&permissions=0&scope=bot)
[![chia](https://img.shields.io/badge/crypto-chia-blueviolet)](https://discord.com/oauth2/authorize?client_id=851203175200456704&permissions=0&scope=bot)
[![lifetoken](https://img.shields.io/badge/crypto-lifetoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=851811517875355680&permissions=0&scope=bot)
[![clucoin](https://img.shields.io/badge/crypto-clucoin-blueviolet)](https://discord.com/oauth2/authorize?client_id=851811687522631710&permissions=0&scope=bot)
[![steel](https://img.shields.io/badge/crypto-steel-blueviolet)](https://discord.com/oauth2/authorize?client_id=851811917333135360&permissions=0&scope=bot)
[![thetatoken](https://img.shields.io/badge/crypto-thetatoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=852238599163871286&permissions=0&scope=bot)
[![forcedao](https://img.shields.io/badge/crypto-forcedao-blueviolet)](https://discord.com/oauth2/authorize?client_id=852239848496562247&permissions=0&scope=bot)
[![icetoken](https://img.shields.io/badge/crypto-icetoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=852239941840011275&permissions=0&scope=bot)
[![tomb](https://img.shields.io/badge/crypto-tomb-blueviolet)](https://discord.com/oauth2/authorize?client_id=852240031090343976&permissions=0&scope=bot)
[![aquariusfi](https://img.shields.io/badge/crypto-aquariusfi-blueviolet)](https://discord.com/oauth2/authorize?client_id=852240124104015932&permissions=0&scope=bot)
[![amptoken](https://img.shields.io/badge/crypto-amptoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=852323149134954506&permissions=0&scope=bot)
[![unicrypt2](https://img.shields.io/badge/crypto-unicrypt2-blueviolet)](https://discord.com/oauth2/authorize?client_id=852323334229327912&permissions=0&scope=bot)
[![cumrocket](https://img.shields.io/badge/crypto-cumrocket-blueviolet)](https://discord.com/oauth2/authorize?client_id=852323588542562334&permissions=0&scope=bot)
[![komodo](https://img.shields.io/badge/crypto-komodo-blueviolet)](https://discord.com/oauth2/authorize?client_id=852323686425952298&permissions=0&scope=bot)
[![waultswappolygon](https://img.shields.io/badge/crypto-waultswappolygon-blueviolet)](https://discord.com/oauth2/authorize?client_id=854392209062887464&permissions=0&scope=bot)
[![ironstablecoin](https://img.shields.io/badge/crypto-ironstablecoin-blueviolet)](https://discord.com/oauth2/authorize?client_id=854392293590433872&permissions=0&scope=bot)
[![xdollar](https://img.shields.io/badge/crypto-xdollar-blueviolet)](https://discord.com/oauth2/authorize?client_id=854392375400464444&permissions=0&scope=bot)
[![xdollarstablecoin](https://img.shields.io/badge/crypto-xdollarstablecoin-blueviolet)](https://discord.com/oauth2/authorize?client_id=854392452034986024&permissions=0&scope=bot)
[![evai](https://img.shields.io/badge/crypto-evai-blueviolet)](https://discord.com/oauth2/authorize?client_id=854392532544520292&permissions=0&scope=bot)
[![polycatfinance](https://img.shields.io/badge/crypto-polycatfinance-blueviolet)](https://discord.com/oauth2/authorize?client_id=854499334455689267&permissions=0&scope=bot)
[![spiritswap](https://img.shields.io/badge/crypto-spiritswap-blueviolet)](https://discord.com/oauth2/authorize?client_id=854499653792956426&permissions=0&scope=bot)
[![comfytoken](https://img.shields.io/badge/crypto-comfytoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=854764554985668629&permissions=0&scope=bot)
[![hodltoken](https://img.shields.io/badge/crypto-hodltoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=854764648938078240&permissions=0&scope=bot)
[![cheecoin](https://img.shields.io/badge/crypto-cheecoin-blueviolet)](https://discord.com/oauth2/authorize?client_id=854764729771098163&permissions=0&scope=bot)
[![particle2](https://img.shields.io/badge/crypto-particle2-blueviolet)](https://discord.com/oauth2/authorize?client_id=856537250833367081&permissions=0&scope=bot)
[![1inch](https://img.shields.io/badge/crypto-1inch-blueviolet)](https://discord.com/oauth2/authorize?client_id=856537449722019880&permissions=0&scope=bot)
[![dero](https://img.shields.io/badge/crypto-dero-blueviolet)](https://discord.com/oauth2/authorize?client_id=856537535597379605&permissions=0&scope=bot)
[![zeppelindao](https://img.shields.io/badge/crypto-zeppelindao-blueviolet)](https://discord.com/oauth2/authorize?client_id=856537610054139929&permissions=0&scope=bot)
[![eject](https://img.shields.io/badge/crypto-eject-blueviolet)](https://discord.com/oauth2/authorize?client_id=856621620154990652&permissions=0&scope=bot)
[![compoundcoin](https://img.shields.io/badge/crypto-compoundcoin-blueviolet)](https://discord.com/oauth2/authorize?client_id=856622019696787516&permissions=0&scope=bot)
[![compoundgovernancetoken](https://img.shields.io/badge/crypto-compoundgovernancetoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=856892639759826984&permissions=0&scope=bot)
[![tombshares](https://img.shields.io/badge/crypto-tombshares-blueviolet)](https://discord.com/oauth2/authorize?client_id=857265987488186389&permissions=0&scope=bot)
[![everrise](https://img.shields.io/badge/crypto-everrise-blueviolet)](https://discord.com/oauth2/authorize?client_id=857357797167333438&permissions=0&scope=bot)
[![dash](https://img.shields.io/badge/crypto-dash-blueviolet)](https://discord.com/oauth2/authorize?client_id=857357876271775774&permissions=0&scope=bot)
[![ergo](https://img.shields.io/badge/crypto-ergo-blueviolet)](https://discord.com/oauth2/authorize?client_id=857357953131348019&permissions=0&scope=bot)
[![tether](https://img.shields.io/badge/crypto-tether-blueviolet)](https://discord.com/oauth2/authorize?client_id=857358033709301792&permissions=0&scope=bot)
[![babydogecoin](https://img.shields.io/badge/crypto-babydogecoin-blueviolet)](https://discord.com/oauth2/authorize?client_id=858008376712888320&permissions=0&scope=bot)
[![swissborg](https://img.shields.io/badge/crypto-swissborg-blueviolet)](https://discord.com/oauth2/authorize?client_id=858367813232361484&permissions=0&scope=bot)
[![utrust](https://img.shields.io/badge/crypto-utrust-blueviolet)](https://discord.com/oauth2/authorize?client_id=859094549354512385&permissions=0&scope=bot)
[![bitbooktoken](https://img.shields.io/badge/crypto-bitbooktoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=859094646492102666&permissions=0&scope=bot)
[![mbitbooks](https://img.shields.io/badge/crypto-mbitbooks-blueviolet)](https://discord.com/oauth2/authorize?client_id=859094736597418024&permissions=0&scope=bot)
[![quantnetwork](https://img.shields.io/badge/crypto-quantnetwork-blueviolet)](https://discord.com/oauth2/authorize?client_id=859094810678001745&permissions=0&scope=bot)
[![ankr](https://img.shields.io/badge/crypto-ankr-blueviolet)](https://discord.com/oauth2/authorize?client_id=859148266121461790&permissions=0&scope=bot)
[![vectorspace](https://img.shields.io/badge/crypto-vectorspace-blueviolet)](https://discord.com/oauth2/authorize?client_id=859148344621924384&permissions=0&scope=bot)
[![just](https://img.shields.io/badge/crypto-just-blueviolet)](https://discord.com/oauth2/authorize?client_id=859148421901713438&permissions=0&scope=bot)
[![smoothlovepotion](https://img.shields.io/badge/crypto-smoothlovepotion-blueviolet)](https://discord.com/oauth2/authorize?client_id=859148505469943860&permissions=0&scope=bot)
[![osmosis](https://img.shields.io/badge/crypto-osmosis-blueviolet)](https://discord.com/oauth2/authorize?client_id=859932007910866944&permissions=0&scope=bot)
[![liqprotocol](https://img.shields.io/badge/crypto-liqprotocol-blueviolet)](https://discord.com/oauth2/authorize?client_id=859932144209756180&permissions=0&scope=bot)
[![axieinfinity](https://img.shields.io/badge/crypto-axieinfinity-blueviolet)](https://discord.com/oauth2/authorize?client_id=859932709262589992&permissions=0&scope=bot)
[![boxaxis](https://img.shields.io/badge/crypto-boxaxis-blueviolet)](https://discord.com/oauth2/authorize?client_id=859932784801742858&permissions=0&scope=bot)
[![dopex](https://img.shields.io/badge/crypto-dopex-blueviolet)](https://discord.com/oauth2/authorize?client_id=860157918186569760&permissions=0&scope=bot)
[![defipulseindex](https://img.shields.io/badge/crypto-defipulseindex-blueviolet)](https://discord.com/oauth2/authorize?client_id=860213184728006686&permissions=0&scope=bot)
[![indexcooperative](https://img.shields.io/badge/crypto-indexcooperative-blueviolet)](https://discord.com/oauth2/authorize?client_id=860213290344775690&permissions=0&scope=bot)
[![huobitoken](https://img.shields.io/badge/crypto-huobitoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=860515915769577492&permissions=0&scope=bot)
[![maker](https://img.shields.io/badge/crypto-maker-blueviolet)](https://discord.com/oauth2/authorize?client_id=860516003964649482&permissions=0&scope=bot)
[![neo](https://img.shields.io/badge/crypto-neo-blueviolet)](https://discord.com/oauth2/authorize?client_id=860516079797534720&permissions=0&scope=bot)
[![scorpiontoken](https://img.shields.io/badge/crypto-scorpiontoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=860516532604633138&permissions=0&scope=bot)
[![cryptocomchain](https://img.shields.io/badge/crypto-cryptocomchain-blueviolet)](https://discord.com/oauth2/authorize?client_id=860516690647580712&permissions=0&scope=bot)
[![cuminu](https://img.shields.io/badge/crypto-cuminu-blueviolet)](https://discord.com/oauth2/authorize?client_id=860516776038236170&permissions=0&scope=bot)
[![akashnetwork](https://img.shields.io/badge/crypto-akashnetwork-blueviolet)](https://discord.com/oauth2/authorize?client_id=862082930062655499&permissions=0&scope=bot)
[![sentinel](https://img.shields.io/badge/crypto-sentinel-blueviolet)](https://discord.com/oauth2/authorize?client_id=862083026275663913&permissions=0&scope=bot)
[![dogelonmars](https://img.shields.io/badge/crypto-dogelonmars-blueviolet)](https://discord.com/oauth2/authorize?client_id=862083128751685662&permissions=0&scope=bot)
[![secret](https://img.shields.io/badge/crypto-secret-blueviolet)](https://discord.com/oauth2/authorize?client_id=862083275720097813&permissions=0&scope=bot)
[![harmony](https://img.shields.io/badge/crypto-harmony-blueviolet)](https://discord.com/oauth2/authorize?client_id=862731942193791015&permissions=0&scope=bot)
[![sifchain](https://img.shields.io/badge/crypto-sifchain-blueviolet)](https://discord.com/oauth2/authorize?client_id=862732082476220426&permissions=0&scope=bot)
[![sefi](https://img.shields.io/badge/crypto-sefi-blueviolet)](https://discord.com/oauth2/authorize?client_id=862737844124778506&permissions=0&scope=bot)
[![boneshibaswap](https://img.shields.io/badge/crypto-boneshibaswap-blueviolet)](https://discord.com/oauth2/authorize?client_id=862737917741891636&permissions=0&scope=bot)
[![leash](https://img.shields.io/badge/crypto-leash-blueviolet)](https://discord.com/oauth2/authorize?client_id=862738128335142922&permissions=0&scope=bot)
[![crowns](https://img.shields.io/badge/crypto-crowns-blueviolet)](https://discord.com/oauth2/authorize?client_id=862738209975697479&permissions=0&scope=bot)
[![okb](https://img.shields.io/badge/crypto-okb-blueviolet)](https://discord.com/oauth2/authorize?client_id=862738288527147038&permissions=0&scope=bot)
[![hex](https://img.shields.io/badge/crypto-hex-blueviolet)](https://discord.com/oauth2/authorize?client_id=864864993404387380&permissions=0&scope=bot)
[![securedmoonrattoken](https://img.shields.io/badge/crypto-securedmoonrattoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=864865710277656586&permissions=0&scope=bot)
[![pancakehunny](https://img.shields.io/badge/crypto-pancakehunny-blueviolet)](https://discord.com/oauth2/authorize?client_id=867439869516840981&permissions=0&scope=bot)
[![fegtoken](https://img.shields.io/badge/crypto-fegtoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=867439942598787072&permissions=0&scope=bot)
[![pornrocket](https://img.shields.io/badge/crypto-pornrocket-blueviolet)](https://discord.com/oauth2/authorize?client_id=867440022755213342&permissions=0&scope=bot)
[![cryptoblades](https://img.shields.io/badge/crypto-cryptoblades-blueviolet)](https://discord.com/oauth2/authorize?client_id=867745267327041556&permissions=0&scope=bot)
[![seedifyfund](https://img.shields.io/badge/crypto-seedifyfund-blueviolet)](https://discord.com/oauth2/authorize?client_id=867745367730159616&permissions=0&scope=bot)
[![ion](https://img.shields.io/badge/crypto-ion-blueviolet)](https://discord.com/oauth2/authorize?client_id=867745444570464266&permissions=0&scope=bot)
[![yieldly](https://img.shields.io/badge/crypto-yieldly-blueviolet)](https://discord.com/oauth2/authorize?client_id=867745508973608971&permissions=0&scope=bot)
[![dungeonswap](https://img.shields.io/badge/crypto-dungeonswap-blueviolet)](https://discord.com/oauth2/authorize?client_id=867745588233895976&permissions=0&scope=bot)
[![frogefinance](https://img.shields.io/badge/crypto-frogefinance-blueviolet)](https://discord.com/oauth2/authorize?client_id=869243381224734760&permissions=0&scope=bot)
[![frogdaodime](https://img.shields.io/badge/crypto-frogdaodime-blueviolet)](https://discord.com/oauth2/authorize?client_id=869243493808214046&permissions=0&scope=bot)
[![nano](https://img.shields.io/badge/crypto-nano-blueviolet)](https://discord.com/oauth2/authorize?client_id=869243560258592829&permissions=0&scope=bot)
[![mydefipet](https://img.shields.io/badge/crypto-mydefipet-blueviolet)](https://discord.com/oauth2/authorize?client_id=869247729132515418&permissions=0&scope=bot)
[![coin98](https://img.shields.io/badge/crypto-coin98-blueviolet)](https://discord.com/oauth2/authorize?client_id=869247821554012192&permissions=0&scope=bot)
[![internetcomputer](https://img.shields.io/badge/crypto-internetcomputer-blueviolet)](https://discord.com/oauth2/authorize?client_id=869247919340023848&permissions=0&scope=bot)
[![thorchain](https://img.shields.io/badge/crypto-thorchain-blueviolet)](https://discord.com/oauth2/authorize?client_id=869369232301916160&permissions=0&scope=bot)
[![metahero](https://img.shields.io/badge/crypto-metahero-blueviolet)](https://discord.com/oauth2/authorize?client_id=869369336899457034&permissions=0&scope=bot)
[![hanuyokia](https://img.shields.io/badge/crypto-hanuyokia-blueviolet)](https://discord.com/oauth2/authorize?client_id=869369442126168096&permissions=0&scope=bot)
[![starlink](https://img.shields.io/badge/crypto-starlink-blueviolet)](https://discord.com/oauth2/authorize?client_id=869369510606544897&permissions=0&scope=bot)
[![fibotoken](https://img.shields.io/badge/crypto-fibotoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=869369587882418217&permissions=0&scope=bot)
[![zcash](https://img.shields.io/badge/crypto-zcash-blueviolet)](https://discord.com/oauth2/authorize?client_id=869571753607852032&permissions=0&scope=bot)
[![tokocrypto](https://img.shields.io/badge/crypto-tokocrypto-blueviolet)](https://discord.com/oauth2/authorize?client_id=869571979911499786&permissions=0&scope=bot)
[![nafter](https://img.shields.io/badge/crypto-nafter-blueviolet)](https://discord.com/oauth2/authorize?client_id=856621765919768606&permissions=0&scope=bot)
[![moneyplanttoken](https://img.shields.io/badge/crypto-moneyplanttoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=869950420133507144&permissions=0&scope=bot)
[![myneighboralice](https://img.shields.io/badge/crypto-myneighboralice-blueviolet)](https://discord.com/oauth2/authorize?client_id=869950487678554162&permissions=0&scope=bot)
[![mist](https://img.shields.io/badge/crypto-mist-blueviolet)](https://discord.com/oauth2/authorize?client_id=869998726628925480&permissions=0&scope=bot)
[![digibyte](https://img.shields.io/badge/crypto-digibyte-blueviolet)](https://discord.com/oauth2/authorize?client_id=869998871676354620&permissions=0&scope=bot)
[![elongate](https://img.shields.io/badge/crypto-elongate-blueviolet)](https://discord.com/oauth2/authorize?client_id=869998983253213226&permissions=0&scope=bot)
[![babyaxie](https://img.shields.io/badge/crypto-babyaxie-blueviolet)](https://discord.com/oauth2/authorize?client_id=869999047161815090&permissions=0&scope=bot)
[![serum](https://img.shields.io/badge/crypto-serum-blueviolet)](https://discord.com/oauth2/authorize?client_id=869999361331953695&permissions=0&scope=bot)
[![darkenergycrystals](https://img.shields.io/badge/crypto-darkenergycrystals-blueviolet)](https://discord.com/oauth2/authorize?client_id=869999418248675379&permissions=0&scope=bot)
[![drakoin](https://img.shields.io/badge/crypto-drakoin-blueviolet)](https://discord.com/oauth2/authorize?client_id=869999472686551070&permissions=0&scope=bot)
[![satoshivisioncoin](https://img.shields.io/badge/crypto-satoshivisioncoin-blueviolet)](https://discord.com/oauth2/authorize?client_id=869999526646280273&permissions=0&scope=bot)
[![weth](https://img.shields.io/badge/crypto-weth-blueviolet)](https://discord.com/oauth2/authorize?client_id=871068140623257610&permissions=0&scope=bot)
[![signum](https://img.shields.io/badge/crypto-signum-blueviolet)](https://discord.com/oauth2/authorize?client_id=871068306591871077&permissions=0&scope=bot)
[![zoocoin](https://img.shields.io/badge/crypto-zoocoin-blueviolet)](https://discord.com/oauth2/authorize?client_id=871068364410355743&permissions=0&scope=bot)
[![casperdefi](https://img.shields.io/badge/crypto-casperdefi-blueviolet)](https://discord.com/oauth2/authorize?client_id=871068423302570066&permissions=0&scope=bot)
[![unidex](https://img.shields.io/badge/crypto-unidex-blueviolet)](https://discord.com/oauth2/authorize?client_id=871068490235273247&permissions=0&scope=bot)
[![plantvsundeadtoken](https://img.shields.io/badge/crypto-plantvsundeadtoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=871812819266461807&permissions=0&scope=bot)
[![yelfinance](https://img.shields.io/badge/crypto-yelfinance-blueviolet)](https://discord.com/oauth2/authorize?client_id=871812952225878086&permissions=0&scope=bot)
[![curvedaotoken](https://img.shields.io/badge/crypto-curvedaotoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=873190952553046057&permissions=0&scope=bot)
[![qidao](https://img.shields.io/badge/crypto-qidao-blueviolet)](https://discord.com/oauth2/authorize?client_id=873191030655160371&permissions=0&scope=bot)
[![cryptozoon](https://img.shields.io/badge/crypto-cryptozoon-blueviolet)](https://discord.com/oauth2/authorize?client_id=873191092659572797&permissions=0&scope=bot)
[![binanceusd](https://img.shields.io/badge/crypto-binanceusd-blueviolet)](https://discord.com/oauth2/authorize?client_id=873191153825095720&permissions=0&scope=bot)
[![diamondplatformtoken](https://img.shields.io/badge/crypto-diamondplatformtoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=873191221298872380&permissions=0&scope=bot)
[![draken](https://img.shields.io/badge/crypto-draken-blueviolet)](https://discord.com/oauth2/authorize?client_id=873589811976495115&permissions=0&scope=bot)
[![nafty](https://img.shields.io/badge/crypto-nafty-blueviolet)](https://discord.com/oauth2/authorize?client_id=874299070548684870&permissions=0&scope=bot)
[![catbread](https://img.shields.io/badge/crypto-catbread-blueviolet)](https://discord.com/oauth2/authorize?client_id=874299138265727018&permissions=0&scope=bot)
[![splinterlands](https://img.shields.io/badge/crypto-splinterlands-blueviolet)](https://discord.com/oauth2/authorize?client_id=874299210076393503&permissions=0&scope=bot)
[![itamgames](https://img.shields.io/badge/crypto-itamgames-blueviolet)](https://discord.com/oauth2/authorize?client_id=874299280913997915&permissions=0&scope=bot)
[![starshiptoken](https://img.shields.io/badge/crypto-starshiptoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=874299343937630208&permissions=0&scope=bot)
[![bitcoinhd](https://img.shields.io/badge/crypto-bitcoinhd-blueviolet)](https://discord.com/oauth2/authorize?client_id=874361647140057089&permissions=0&scope=bot)
[![paxgold](https://img.shields.io/badge/crypto-paxgold-blueviolet)](https://discord.com/oauth2/authorize?client_id=874361710549532742&permissions=0&scope=bot)
[![starship](https://img.shields.io/badge/crypto-starship-blueviolet)](https://discord.com/oauth2/authorize?client_id=874361797392601108&permissions=0&scope=bot)
[![dinox](https://img.shields.io/badge/crypto-dinox-blueviolet)](https://discord.com/oauth2/authorize?client_id=874361972186034188&permissions=0&scope=bot)
[![skillchain](https://img.shields.io/badge/crypto-skillchain-blueviolet)](https://discord.com/oauth2/authorize?client_id=874679320864501820&permissions=0&scope=bot)
[![farmdefi](https://img.shields.io/badge/crypto-farmdefi-blueviolet)](https://discord.com/oauth2/authorize?client_id=874679412510060645&permissions=0&scope=bot)
[![vitae](https://img.shields.io/badge/crypto-vitae-blueviolet)](https://discord.com/oauth2/authorize?client_id=875019754761490522&permissions=0&scope=bot)
[![blockcreatures](https://img.shields.io/badge/crypto-blockcreatures-blueviolet)](https://discord.com/oauth2/authorize?client_id=875019836609167411&permissions=0&scope=bot)
[![drakeballtoken](https://img.shields.io/badge/crypto-drakeballtoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=875019902052872212&permissions=0&scope=bot)
[![zootoken](https://img.shields.io/badge/crypto-zootoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=875019962832539658&permissions=0&scope=bot)
[![polkamonster](https://img.shields.io/badge/crypto-polkamonster-blueviolet)](https://discord.com/oauth2/authorize?client_id=875431986351702106&permissions=0&scope=bot)
[![revomon](https://img.shields.io/badge/crypto-revomon-blueviolet)](https://discord.com/oauth2/authorize?client_id=875432060209213450&permissions=0&scope=bot)
[![gala](https://img.shields.io/badge/crypto-gala-blueviolet)](https://discord.com/oauth2/authorize?client_id=875432125246079006&permissions=0&scope=bot)
[![flow](https://img.shields.io/badge/crypto-flow-blueviolet)](https://discord.com/oauth2/authorize?client_id=875432186088652801&permissions=0&scope=bot)
[![wax](https://img.shields.io/badge/crypto-wax-blueviolet)](https://discord.com/oauth2/authorize?client_id=875432263184187393&permissions=0&scope=bot)
[![arweave](https://img.shields.io/badge/crypto-arweave-blueviolet)](https://discord.com/oauth2/authorize?client_id=875757316413227008&permissions=0&scope=bot)
[![singularitynet](https://img.shields.io/badge/crypto-singularitynet-blueviolet)](https://discord.com/oauth2/authorize?client_id=875757390644002826&permissions=0&scope=bot)
[![singularitydao](https://img.shields.io/badge/crypto-singularitydao-blueviolet)](https://discord.com/oauth2/authorize?client_id=875757452765831188&permissions=0&scope=bot)
[![bandprotocol](https://img.shields.io/badge/crypto-bandprotocol-blueviolet)](https://discord.com/oauth2/authorize?client_id=875757573213655070&permissions=0&scope=bot)
[![charli3](https://img.shields.io/badge/crypto-charli3-blueviolet)](https://discord.com/oauth2/authorize?client_id=875757649449320488&permissions=0&scope=bot)
[![alienworlds](https://img.shields.io/badge/crypto-alienworlds-blueviolet)](https://discord.com/oauth2/authorize?client_id=876825386078568458&permissions=0&scope=bot)
[![cardstarter](https://img.shields.io/badge/crypto-cardstarter-blueviolet)](https://discord.com/oauth2/authorize?client_id=876825456995889172&permissions=0&scope=bot)
[![mooncake](https://img.shields.io/badge/crypto-mooncake-blueviolet)](https://discord.com/oauth2/authorize?client_id=876885138594758667&permissions=0&scope=bot)
[![saitamainu](https://img.shields.io/badge/crypto-saitamainu-blueviolet)](https://discord.com/oauth2/authorize?client_id=876885219888734219&permissions=0&scope=bot)
[![gulagtoken](https://img.shields.io/badge/crypto-gulagtoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=876885288536911923&permissions=0&scope=bot)
[![coinarytoken](https://img.shields.io/badge/crypto-coinarytoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=877903037564346398&permissions=0&scope=bot)
[![avalanche2](https://img.shields.io/badge/crypto-avalanche2-blueviolet)](https://discord.com/oauth2/authorize?client_id=877903202211754094&permissions=0&scope=bot)
[![mobox](https://img.shields.io/badge/crypto-mobox-blueviolet)](https://discord.com/oauth2/authorize?client_id=877903375956594719&permissions=0&scope=bot)
[![blockapescissors](https://img.shields.io/badge/crypto-blockapescissors-blueviolet)](https://discord.com/oauth2/authorize?client_id=878464235691204638&permissions=0&scope=bot)
[![polkaswap](https://img.shields.io/badge/crypto-polkaswap-blueviolet)](https://discord.com/oauth2/authorize?client_id=878464354973016074&permissions=0&scope=bot)
[![chiliz](https://img.shields.io/badge/crypto-chiliz-blueviolet)](https://discord.com/oauth2/authorize?client_id=878464463911677973&permissions=0&scope=bot)
[![thecryptoprophecies](https://img.shields.io/badge/crypto-thecryptoprophecies-blueviolet)](https://discord.com/oauth2/authorize?client_id=879447028017561670&permissions=0&scope=bot)
[![catgirl](https://img.shields.io/badge/crypto-catgirl-blueviolet)](https://discord.com/oauth2/authorize?client_id=879447101732454430&permissions=0&scope=bot)
[![cryptocars](https://img.shields.io/badge/crypto-cryptocars-blueviolet)](https://discord.com/oauth2/authorize?client_id=879447174210011146&permissions=0&scope=bot)
[![mxc](https://img.shields.io/badge/crypto-mxc-blueviolet)](https://discord.com/oauth2/authorize?client_id=879447248986046464&permissions=0&scope=bot)
[![binamon](https://img.shields.io/badge/crypto-binamon-blueviolet)](https://discord.com/oauth2/authorize?client_id=880628836679680000&permissions=0&scope=bot)
[![minaprotocol](https://img.shields.io/badge/crypto-minaprotocol-blueviolet)](https://discord.com/oauth2/authorize?client_id=880629111503073301&permissions=0&scope=bot)
[![binemon](https://img.shields.io/badge/crypto-binemon-blueviolet)](https://discord.com/oauth2/authorize?client_id=880629231116238888&permissions=0&scope=bot)
[![dehive](https://img.shields.io/badge/crypto-dehive-blueviolet)](https://discord.com/oauth2/authorize?client_id=880629354583949362&permissions=0&scope=bot)
[![joe](https://img.shields.io/badge/crypto-joe-blueviolet)](https://discord.com/oauth2/authorize?client_id=880806610400845864&permissions=0&scope=bot)
[![tripscommunity](https://img.shields.io/badge/crypto-tripscommunity-blueviolet)](https://discord.com/oauth2/authorize?client_id=880806704642654239&permissions=0&scope=bot)
[![wrappedbitcoin](https://img.shields.io/badge/crypto-wrappedbitcoin-blueviolet)](https://discord.com/oauth2/authorize?client_id=880806808111964251&permissions=0&scope=bot)
[![mimatic](https://img.shields.io/badge/crypto-mimatic-blueviolet)](https://discord.com/oauth2/authorize?client_id=880806882124632114&permissions=0&scope=bot)
[![dogecola](https://img.shields.io/badge/crypto-dogecola-blueviolet)](https://discord.com/oauth2/authorize?client_id=880806963674488955&permissions=0&scope=bot)
[![constellationlabs](https://img.shields.io/badge/crypto-constellationlabs-blueviolet)](https://discord.com/oauth2/authorize?client_id=881324625705992262&permissions=0&scope=bot)
[![dragonslayer](https://img.shields.io/badge/crypto-dragonslayer-blueviolet)](https://discord.com/oauth2/authorize?client_id=881324711181697045&permissions=0&scope=bot)
[![bunnypark](https://img.shields.io/badge/crypto-bunnypark-blueviolet)](https://discord.com/oauth2/authorize?client_id=881324784129024041&permissions=0&scope=bot)
[![hibikifinance](https://img.shields.io/badge/crypto-hibikifinance-blueviolet)](https://discord.com/oauth2/authorize?client_id=882321215421816924&permissions=0&scope=bot)
[![revolvertoken](https://img.shields.io/badge/crypto-revolvertoken-blueviolet)](https://discord.com/oauth2/authorize?client_id=882321305565810729&permissions=0&scope=bot)
[![kucoinshares](https://img.shields.io/badge/crypto-kucoinshares-blueviolet)](https://discord.com/oauth2/authorize?client_id=882321377321955348&permissions=0&scope=bot)
[![atari](https://img.shields.io/badge/crypto-atari-blueviolet)](https://discord.com/oauth2/authorize?client_id=882420406894616606&permissions=0&scope=bot)
[![polkastarter](https://img.shields.io/badge/crypto-polkastarter-blueviolet)](https://discord.com/oauth2/authorize?client_id=882420501329375252&permissions=0&scope=bot)
[![polylauncher](https://img.shields.io/badge/crypto-polylauncher-blueviolet)](https://discord.com/oauth2/authorize?client_id=882420591322333324&permissions=0&scope=bot)
[![wanakafarm](https://img.shields.io/badge/crypto-wanakafarm-blueviolet)](https://discord.com/oauth2/authorize?client_id=882420667725795438&permissions=0&scope=bot)
[![cryptohounds](https://img.shields.io/badge/crypto-cryptohounds-blueviolet)](https://discord.com/oauth2/authorize?client_id=882712094812799066&permissions=0&scope=bot)
[![verasity](https://img.shields.io/badge/crypto-verasity-blueviolet)](https://discord.com/oauth2/authorize?client_id=882712166992576574&permissions=0&scope=bot)
[![alphafinance](https://img.shields.io/badge/crypto-alphafinance-blueviolet)](https://discord.com/oauth2/authorize?client_id=882712243303755807&permissions=0&scope=bot)
[![alchemypay](https://img.shields.io/badge/crypto-alchemypay-blueviolet)](https://discord.com/oauth2/authorize?client_id=882712403631046687&permissions=0&scope=bot)
[![strong](https://img.shields.io/badge/crypto-strong-blueviolet)](https://discord.com/oauth2/authorize?client_id=882712509335891998&permissions=0&scope=bot)
[![derace](https://img.shields.io/badge/crypto-derace-blueviolet)](https://discord.com/oauth2/authorize?client_id=885172849445318677&permissions=0&scope=bot)
[![staratlas](https://img.shields.io/badge/crypto-staratlas-blueviolet)](https://discord.com/oauth2/authorize?client_id=885172879556235295&permissions=0&scope=bot)
[![deathroad](https://img.shields.io/badge/crypto-deathroad-blueviolet)](https://discord.com/oauth2/authorize?client_id=885172906815029290&permissions=0&scope=bot)
[![polis](https://img.shields.io/badge/crypto-polis-blueviolet)](https://discord.com/oauth2/authorize?client_id=885172933994110996&permissions=0&scope=bot)

### Gas Prices

[![Ethereum Invite Link](https://user-images.githubusercontent.com/7338312/127579033-8785ed17-2bcc-474c-80d4-8ea356da70e6.png)](https://discord.com/api/oauth2/authorize?client_id=833797002684661821&permissions=0&scope=bot)[![Binance Smart Chain Invite Link](https://user-images.githubusercontent.com/7338312/127578976-d47069cb-c162-4ab5-ad73-be17b2c1796d.png)](https://discord.com/api/oauth2/authorize?client_id=856947934452645898&permissions=0&scope=bot)[![Polygon Invite Link](https://user-images.githubusercontent.com/7338312/127578967-a7097067-9b0a-44d2-baf6-e3541a511c70.png)](https://discord.com/api/oauth2/authorize?client_id=857023179210096674&permissions=0&scope=bot)

### Other (not my) crypto discord bots

[![image](https://user-images.githubusercontent.com/7338312/135726609-f3504a1e-7c2a-457e-9476-b50e0974e764.png)](https://discord.com/oauth2/authorize?client_id=893362064842706994&permissions=0&scope=bot)

## Premium

![Discord Sidebar w/ Premium Bots](https://s3.cloud.rileysnyder.org/public/assets/sidebar-premium.png)

For advanced features like faster update times and color changing names on price changes you can subscribe to my premuim offering.

Price per bot (paid monthly): $1
Price per bot (paid yearly):  $10

If you are interested please see the [contact info on my github page](https://github.com/rssnyder) and send me a messgae via your platform of choice (discord perferred). For a live demo, join the support discord linked at the top or bottom of this page.

## Self-Hosting - Docker

⚠️ As of version **3.5.0** we are using `mattn/go-sqlite3` to store state. Since this is a CGO package cross-compilation is more difficult. Because of this running on non linux-x86 machines may require you to build from source. I am currently working on publishing offical builds again for other OS/ARCH and will remove this warning when the work has been completed.

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

⚠️ As of version **3.5.0** we are using `mattn/go-sqlite3` to store state. Since this is a CGO package cross-compilation is more difficult. Because of this running on non linux-x86 machines may require you to build from source. I am currently working on publishing offical builds again for other OS/ARCH and will remove this warning when the work has been completed.

This bot is distributed as a docker image and a binary.

The program acts as a manager of one to many bots. You can have one running instance of the program and have any number of bots running within it.

[Click here](https://youtu.be/LhgCdtE8kmc) to watch a quick video tutorial on how to self-host these bots on linux.

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
  "currency_sumbol": "AUD",                         # string/OPTIONAL: alternative curreny symbol
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
  "currency_sumbol": "AUD",                         # string/OPTIONAL: alternative curreny symbol
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
