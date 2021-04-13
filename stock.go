package main

import "time"

type Stock struct {
	Ticker      string        `json:"ticker"`   // stock symbol
	Name        string        `json:"name"`     // override for symbol as shown on the bot
	Nickname    string        `json:"nickname"` // flag for changing nickname
	Color       string        `json:"color"`
	FlashChange string        `json:"flash_change"`
	Frequency   time.Duration `json:"frequency"` // how often to update in seconds
	Price       int           `json:"price"`
	token       string        `json:"-"` // discord token
	close       chan int      `json:"-"`
}

// enumerate the types
const (
	CryptoType = iota
	StockType
)

// NewStock saves information about the stock and starts up a watcher on it
func NewStock(ticker string, token string, name string, nickname string, color string, flashChange string, frequency int) *Stock {
	s := &Stock{
		Ticker:      ticker,
		Name:        name,
		Nickname:    nickname,
		Color:       color,
		FlashChange: flashChange,
		Frequency:   time.Duration(frequency) * time.Second,
		token:       token,
		close:       make(chan int, 1),
	}

	// spin off go routine to watch the price
	go s.watchStockPrice()
	return s
}

// NewCrypto saves information about the crypto and starts up a watcher on it
func NewCrypto(ticker string, token string, name string, nickname string, color string, flashChange string, frequency int) *Stock {
	s := &Stock{
		Ticker:      ticker,
		Name:        name,
		Nickname:    nickname,
		Color:       color,
		FlashChange: flashChange,
		Frequency:   time.Duration(frequency) * time.Second,
		token:       token,
		close:       make(chan int, 1),
	}

	// spin off go routine to watch the price
	go s.watchCryptoPrice()
	return s
}

// Shutdown sends a signal to shut off the goroutine
func (s *Stock) Shutdown() {
	s.close <- 1
}

func (s *Stock) watchStockPrice() {
	logger.Debugf("Watching stock price for %s", s.Name)
	ticker := time.NewTicker(s.Frequency)
	// continuously watch
	for {
		select {
		case <-s.close:
			logger.Infof("Shutting down price watching for %s", s.Name)
			return
		case <-ticker.C:
			logger.Debugf("Fetching stock price for %s", s.Name)
			/*
					data = get_stock_price(ticker)
				   price_data = data.get('quoteSummary', {}).get('result', []).pop().get('price', {})
				   price = price_data.get('regularMarketPrice', {}).get('raw', 0.00)

				   # If after hours, get change
				   if price_data.get('postMarketChange'):

				       # Get difference or new price
				       if getenv('POST_MARKET_PRICE'):
				           post_market_target = 'postMarketPrice'
				       else:
				           post_market_target = 'postMarketChange'

				       raw_diff = price_data.get(post_market_target, {}).get('raw', 0.00)
				       diff = round(raw_diff, 2)

				       if not getenv('POST_MARKET_PRICE'):
				           if diff >= 0.0:
				               change_up = True
				               diff = '+' + str(diff)
				           else:
				               change_up = False

				       activity_content = f'${price} AHT {diff}'
				       logging.info(f'stock after hours price retrived: {activity_content}')
				   else:
				       raw_diff = price_data.get('regularMarketChange', {}).get('raw', 0.00)
				       diff = round(raw_diff, 2)
				       if diff >= 0.0:
				           diff = '+' + str(diff)
				       else:
				           change_up = False


				       activity_content = f'${price} / {diff}'
				       logging.info(f'stock price retrived: {activity_content}')

				   # Change name via nickname if set
				   if change_nick:

				       for server in self.guilds:

				           green = discord.utils.get(server.roles, name="tickers-green")
				           red = discord.utils.get(server.roles, name="tickers-red")

				           try:
				               await server.me.edit(
				                   nick=f'{name} - ${price}'
				               )

				               if change_color:

				                   if flash_change:
				                       # Flash price change
				                       if price >= old_price:
				                           await server.me.add_roles(green)
				                           await server.me.remove_roles(red)
				                       else:
				                           await server.me.add_roles(red)
				                           await server.me.remove_roles(green)

				                   # Stay on day change
				                   if change_up:
				                       await server.me.add_roles(green)
				                       await server.me.remove_roles(red)
				                   else:
				                       await server.me.add_roles(red)
				                       await server.me.remove_roles(green)

				           except discord.HTTPException as e:
				               logging.error(f'updating nick failed: {e.status}: {e.text}')
				           except discord.Forbidden as f:
				               logging.error(f'lacking perms for chaning nick: {f.status}: {f.text}')

				           logging.info(f'stock updated nick in {server.name}')

				       # Check what price we are displaying
				       if price_data.get('postMarketChange'):
				           activity_content_header = 'After Hours'
				       else:
				           activity_content_header = 'Day Change'

				       activity_content = f'{activity_content_header}: {diff}'

				   # Change activity
				   try:
				       await self.change_presence(
				           activity=discord.Activity(
				               type=discord.ActivityType.watching,
				               name=activity_content
				           )
				       )

				       logging.info(f'stock activity updated: {activity_content}')

				   except discord.InvalidArgument as e:
				       logging.error(f'updating activity failed: {e.status}: {e.text}')

				   old_price = price

				   logging.info(f'stock sleeping for {frequency}s')
				   await asyncio.sleep(int(frequency))
				   logging.info('stock sleep ended')
			*/
		}

	}

}

func (s *Stock) watchCryptoPrice() {
	ticker := time.NewTicker(s.Frequency)
	logger.Debugf("Watching crypto price for %s", s.Name)
	// continuously watch
	for {
		select {
		case <-s.close:
			logger.Infof("Shutting down price watching for %s", s.Name)
			return
		case <-ticker.C:
			logger.Debugf("Fetching crypto price for %s", s.Name)
			/*

			   # Grab the current price data
			   data = get_crypto_price(crypto_name)
			   price = data.get('market_data', {}).get('current_price', {}).get(CURRENCY, 0.0)
			   change = data.get('market_data', {}).get('price_change_24h', 0)
			   change_header = ''
			   if change >= 0.0:
			      change_header = '+'
			   else:
			      change_up = False

			   logging.info(f'crypto price retrived {price}')

			   activity_content = f'${price} / {change_header}{change}'

			   # Change name via nickname if set
			   if change_nick:

			      for server in self.guilds:

			   	   green = discord.utils.get(server.roles, name="tickers-green")
			   	   red = discord.utils.get(server.roles, name="tickers-red")

			   	   try:
			   		   await server.me.edit(
			   			   nick=f'{ticker} - ${price}'
			   		   )

			   		   if change_color:

			   			   if flash_change:
			   				   # Flash price change
			   				   if price >= old_price:
			   					   await server.me.add_roles(green)
			   					   await server.me.remove_roles(red)
			   				   else:
			   					   await server.me.add_roles(red)
			   					   await server.me.remove_roles(green)

			   			   # Stay on day change
			   			   if change_up:
			   				   await server.me.add_roles(green)
			   				   await server.me.remove_roles(red)
			   			   else:
			   				   await server.me.add_roles(red)
			   				   await server.me.remove_roles(green)

			   	   except discord.HTTPException as e:
			   		   logging.error(f'updating nick failed: {e.status}: {e.text}')
			   	   except discord.Forbidden as f:
			   		   logging.error(f'lacking perms for chaning nick: {f.status}: {f.text}')

			   	   logging.info(f'{crypto_name} updated nick in {server.name}')

			      # Use activity for other fun stuff
			      activity_content = f'24hr: {change_header}{change}'

			   # Change activity
			   try:
			      await self.change_presence(
			   	   activity=discord.Activity(
			   		   type=discord.ActivityType.watching,
			   		   name=activity_content
			   	   )
			      )

			      old_price = price
			      logging.info(f'crypto activity updated {activity_content}')
			   except discord.InvalidArgument as e:
			      logging.error(f'updating activity failed: {e.status}: {e.text}')

			   # Only update every min
			   logging.info(f'crypto sleeping for {frequency}s')
			   await asyncio.sleep(int(frequency))
			   logging.info('crypto sleep ended')
			*/
		}
	}
}
