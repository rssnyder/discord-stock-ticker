package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/rssnyder/discord-stock-ticker/utils"
)

type Stock struct {
	Ticker      string        `json:"ticker"`   // stock symbol
	Name        string        `json:"name"`     // override for symbol as shown on the bot
	Nickname    bool          `json:"nickname"` // flag for changing nickname
	Color       bool          `json:"color"`
	FlashChange bool          `json:"flash_change"`
	Frequency   time.Duration `json:"frequency"` // how often to update in seconds
	Price       int           `json:"price"`
	token       string        `json:"-"` // discord token
	close       chan int      `json:"-"`
}

// NewStock saves information about the stock and starts up a watcher on it
func NewStock(ticker string, token string, name string, nickname bool, color bool, flashChange bool, frequency int) *Stock {
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
func NewCrypto(ticker string, token string, name string, nickname bool, color bool, flashChange bool, frequency int) *Stock {
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

	// create a new discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + s.token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// show as online
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening discord connection,", err)
		return
	}

	// get bot id
	botUser, err := dg.User("@me")
	if err != nil {
		fmt.Println("Error getting bot id: ", err)
		return
	}

	// Get guides for bot
	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		fmt.Println("Error getting guilds: ", err)
		s.Nickname = false
	}

	logger.Infof("Watching stock price for %s", s.Name)
	ticker := time.NewTicker(s.Frequency)

	// continuously watch
	for {
		select {
		case <-s.close:
			logger.Infof("Shutting down price watching for %s", s.Name)
			return
		case <-ticker.C:
			logger.Infof("Fetching stock price for %s", s.Name)

			var priceData utils.PriceResults
			var fmtPrice string
			var fmtDiff string

			// save the price struct & do something with it
			priceData, err = utils.GetStockPrice(s.Ticker)
			if err != nil {
				logger.Errorf("Unable to fetch stock price for %s", s.Name)
			}
			fmtPrice = priceData.QuoteSummary.Results[0].Price.RegularMarketPrice.Fmt

			// check for day or after hours change
			var emptyChange utils.Change

			if priceData.QuoteSummary.Results[0].Price.PostMarketChange != emptyChange {
				fmtDiff = priceData.QuoteSummary.Results[0].Price.PostMarketChange.Fmt
			} else {
				fmtDiff = priceData.QuoteSummary.Results[0].Price.RegularMarketChange.Fmt
			}

			// calculate if price has moved up or down
			var increase bool
			if string(fmtDiff[0]) == "-" {
				increase = false
			} else {
				increase = true
			}

			if s.Nickname {
				// update nickname instead of activity
				var nickname string
				var activity string

				// format nickname
				nickname = fmt.Sprintf("%s - $%s", s.Name, fmtPrice)

				// format activity based on trading time
				if priceData.QuoteSummary.Results[0].Price.PostMarketChange != emptyChange {
					activity = fmt.Sprintf("Change: %s", fmtDiff)
				} else {
					activity = fmt.Sprintf("AHT: %s", fmtDiff)
				}

				// Update nickname in guilds
				for _, g := range guilds {
					err = dg.GuildMemberNickname(g.ID, "@me", nickname)
					if err != nil {
						fmt.Println("Error updating nickname: ", err)
						continue
					}
					logger.Infof("Set nickname in %s: %s", g.Name, nickname)

					if s.Color {
						// get roles for colors
						var redRole string
						var greeenRole string

						roles, err := dg.GuildRoles(g.ID)
						if err != nil {
							fmt.Println("Error getting guilds: ", err)
							continue
						}

						// find role ids
						for _, r := range roles {
							if r.Name == "tickers-red" {
								redRole = r.ID
							} else if r.Name == "tickers-green" {
								greeenRole = r.ID
							}
						}

						if len(redRole) == 0 || len(greeenRole) == 0 {
							logger.Error("Unable to find roles for color changes")
							continue
						}

						// assign role based on change
						if increase {
							dg.GuildMemberRoleRemove(g.ID, botUser.ID, redRole)
							err = dg.GuildMemberRoleAdd(g.ID, botUser.ID, greeenRole)
							if err != nil {
								logger.Errorf("Unable to set role: ", err)
							}
						} else {
							dg.GuildMemberRoleRemove(g.ID, botUser.ID, greeenRole)
							err = dg.GuildMemberRoleAdd(g.ID, botUser.ID, redRole)
							if err != nil {
								logger.Errorf("Unable to set role: ", err)
							}
						}
					}
				}

				err = dg.UpdateListeningStatus(activity)
				if err != nil {
					logger.Errorf("Unable to set activity: ", err)
				} else {
					logger.Infof("Set activity: %s", activity)
				}

			} else {
				var activity string

				// format activity based on trading time
				if priceData.QuoteSummary.Results[0].Price.PostMarketChange != emptyChange {
					activity = fmt.Sprintf("%s AHT %s", fmtPrice, fmtDiff)
				} else {
					activity = fmt.Sprintf("%s - $%s", fmtPrice, fmtDiff)
				}

				err = dg.UpdateListeningStatus(activity)
				if err != nil {
					logger.Errorf("Unable to set activity: ", err)
				} else {
					logger.Infof("Set activity: %s", activity)
				}

			}

		}

	}

}

func (s *Stock) watchCryptoPrice() {

	// create a new discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + s.token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// show as online
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening discord connection,", err)
		return
	}

	// get bot id
	botUser, err := dg.User("@me")
	if err != nil {
		fmt.Println("Error getting bot id: ", err)
		return
	}

	// Get guides for bot
	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		fmt.Println("Error getting guilds: ", err)
		s.Nickname = false
	}

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

			var priceData utils.GeckoPriceResults
			var fmtPrice string
			var fmtDiff string

			// save the price struct & do something with it
			priceData, err = utils.GetCryptoPrice(s.Name)
			if err != nil {
				logger.Errorf("Unable to fetch stock price for %s: %s", s.Name, err)
			}
			fmtPrice = strconv.Itoa(priceData.MarketData.CurrentPrice.USD)
			fmtDiff = strconv.Itoa(priceData.MarketData.PriceChange)

			// calculate if price has moved up or down
			var increase bool
			if string(fmtDiff[0]) == "-" {
				increase = false
			} else {
				increase = true
			}

			if s.Nickname {
				// update nickname instead of activity
				var nickname string
				var activity string

				// format nickname
				nickname = fmt.Sprintf("%s - $%s", s.Ticker, fmtPrice)

				// format activity
				activity = fmt.Sprintf("24hr - $%s", fmtDiff)

				// Update nickname in guilds
				for _, g := range guilds {
					err = dg.GuildMemberNickname(g.ID, "@me", nickname)
					if err != nil {
						fmt.Println("Error updating nickname: ", err)
						continue
					}
					logger.Infof("Set nickname in %s: %s", g.Name, nickname)

					if s.Color {
						// get roles for colors
						var redRole string
						var greeenRole string

						roles, err := dg.GuildRoles(g.ID)
						if err != nil {
							fmt.Println("Error getting guilds: ", err)
							continue
						}

						// find role ids
						for _, r := range roles {
							if r.Name == "tickers-red" {
								redRole = r.ID
							} else if r.Name == "tickers-green" {
								greeenRole = r.ID
							}
						}

						if len(redRole) == 0 || len(greeenRole) == 0 {
							logger.Error("Unable to find roles for color changes")
							continue
						}

						// assign role based on change
						if increase {
							err = dg.GuildMemberRoleRemove(g.ID, botUser.ID, redRole)
							if err != nil {
								logger.Errorf("Unable to remove role: ", err)
							}
							err = dg.GuildMemberRoleAdd(g.ID, botUser.ID, greeenRole)
							if err != nil {
								logger.Errorf("Unable to set role: ", err)
							}
						} else {
							err = dg.GuildMemberRoleRemove(g.ID, botUser.ID, greeenRole)
							if err != nil {
								logger.Errorf("Unable to remove role: ", err)
							}
							err = dg.GuildMemberRoleAdd(g.ID, botUser.ID, redRole)
							if err != nil {
								logger.Errorf("Unable to set role: ", err)
							}
						}
					}
				}

				err = dg.UpdateListeningStatus(activity)
				if err != nil {
					logger.Errorf("Unable to set activity: ", err)
				} else {
					logger.Infof("Set activity: %s", activity)
				}

			} else {
				var activity string

				// format activity
				activity = fmt.Sprintf("%s - $%s", s.Ticker, fmtPrice)
				err = dg.UpdateListeningStatus(activity)
				if err != nil {
					logger.Errorf("Unable to set activity: ", err)
				} else {
					logger.Infof("Set activity: %s", activity)
				}

			}

		}
	}
}
