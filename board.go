package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/rssnyder/discord-stock-ticker/utils"
)

type Board struct {
	Items      []string `json:"items"`
	Name       string   `json:"name"`
	Crypto     bool     `json:"crypto"`
	Header     string   `json:"header"`
	Nickname   bool     `json:"nickname"`
	Color      bool     `json:"color"`
	Percentage bool     `json:"percentage"`
	Arrows     bool     `json:"arrows"`
	Frequency  int      `json:"frequency"`
	ClientID   string   `json:"client_id"`
	Token      string   `json:"discord_bot_token"`
	Close      chan int `json:"-"`
}

// label returns a human readble id for this bot
func (b *Board) label() string {
	label := strings.ToLower(b.Name)
	if len(label) > 32 {
		label = label[:32]
	}
	return label
}

func (b *Board) watchStockPrice() {

	// create a new discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + b.Token)
	if err != nil {
		logger.Errorf("Error creating Discord session: %s\n", err)
		lastUpdate.With(prometheus.Labels{"type": "board", "ticker": b.Name, "guild": "None"}).Set(0)
		return
	}

	// show as online
	err = dg.Open()
	if err != nil {
		logger.Errorf("error opening discord connection: %s\n", err)
		lastUpdate.With(prometheus.Labels{"type": "board", "ticker": b.Name, "guild": "None"}).Set(0)
		return
	}

	// Get guides for bot
	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		logger.Errorf("Error getting guilds: %s\n", err)
		b.Nickname = false
	}
	if len(guilds) == 0 {
		b.Nickname = false
	}

	// perform management operations
	if *managed {
		setName(dg, b.label())
	}

	logger.Infof("Watching board for %s", b.Name)
	ticker := time.NewTicker(time.Duration(b.Frequency) * time.Second)

	// continuously watch
	for {
		for _, symbol := range b.Items {
			select {
			case <-b.Close:
				logger.Infof("Shutting down price watching for %s", b.Name)
				return
			case <-ticker.C:

				logger.Infof("Fetching stock price for %s", symbol)

				var priceData utils.PriceResults
				var fmtPrice string
				var fmtDiff string

				// save the price struct & do something with it
				priceData, err = utils.GetStockPrice(symbol)
				if err != nil {
					logger.Errorf("Unable to fetch stock price for %s", symbol)
				}

				if len(priceData.QuoteSummary.Results) == 0 {
					logger.Errorf("Yahoo returned bad data for %s", symbol)
					continue
				}
				fmtPrice = priceData.QuoteSummary.Results[0].Price.RegularMarketPrice.Fmt

				var activityHeader string

				if b.Percentage {
					activityHeader = ""
				} else {
					activityHeader = "$"
				}

				// check for day or after hours change
				var emptyChange utils.Change

				if priceData.QuoteSummary.Results[0].Price.PostMarketChange != emptyChange {
					if b.Percentage {
						fmtDiff = priceData.QuoteSummary.Results[0].Price.PostMarketChangePercent.Fmt
					} else {
						fmtDiff = priceData.QuoteSummary.Results[0].Price.PostMarketChange.Fmt
					}
				} else {
					if b.Percentage {
						fmtDiff = priceData.QuoteSummary.Results[0].Price.RegularMarketChangePercent.Fmt
					} else {
						fmtDiff = priceData.QuoteSummary.Results[0].Price.RegularMarketChange.Fmt
					}
				}

				// calculate if price has moved up or down
				var increase bool
				if len(fmtDiff) == 0 {
					increase = true
				} else if string(fmtDiff[0]) == "-" {
					increase = false
				} else {
					increase = true
				}

				decorator := "⬊"
				if increase {
					decorator = "⬈"
				}

				if !b.Arrows {
					decorator = "-"
				}

				if b.Nickname {
					// update nickname instead of activity
					var nickname string
					var activity string

					displayName := b.Header + strings.ToUpper(symbol)

					// format nickname
					nickname = fmt.Sprintf("%s %s $%s", displayName, decorator, fmtPrice)

					// format activity based on trading time
					if priceData.QuoteSummary.Results[0].Price.PostMarketChange == emptyChange {
						activity = fmt.Sprintf("Change: %s%s", activityHeader, fmtDiff)
					} else {
						activity = fmt.Sprintf("AHT: %s%s", activityHeader, fmtDiff)
					}

					// Update nickname in guilds
					for _, g := range guilds {
						err = dg.GuildMemberNickname(g.ID, "@me", nickname)
						if err != nil {
							logger.Errorf("Error updating nickname: %s\n", err)
							continue
						}
						logger.Infof("Set nickname in %s: %s", g.Name, nickname)
						lastUpdate.With(prometheus.Labels{"type": "board", "ticker": b.Name, "guild": g.Name}).SetToCurrentTime()

						// change bot color
						err = setRole(dg, b.ClientID, g.ID, increase)
						if err != nil {
							logger.Errorf("Color roles: %s", err)
						}
					}

					err = dg.UpdateGameStatus(0, b.Name)
					if err != nil {
						logger.Errorf("Unable to set activity: %s\n", err)
					} else {
						logger.Infof("Set activity: %s", activity)
					}

				} else {
					var activity string

					// format activity based on trading time
					if priceData.QuoteSummary.Results[0].Price.PostMarketChange != emptyChange {
						activity = fmt.Sprintf("%s %s AHT %s", symbol, fmtPrice, fmtDiff)
					} else {
						activity = fmt.Sprintf("%s %s %s $%s", symbol, fmtPrice, decorator, fmtDiff)
					}

					err = dg.UpdateGameStatus(0, activity)
					if err != nil {
						logger.Errorf("Unable to set activity: %s\n", err)
					} else {
						logger.Infof("Set activity: %s", activity)
						lastUpdate.With(prometheus.Labels{"type": "board", "ticker": b.Name, "guild": "None"}).SetToCurrentTime()
					}
				}
			}
		}
	}
}

func (b *Board) watchCryptoPrice() {
	var nilCache *redis.Client

	// create a new discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + b.Token)
	if err != nil {
		logger.Errorf("Error creating Discord session: %s\n", err)
		lastUpdate.With(prometheus.Labels{"type": "board", "ticker": b.Name, "guild": "None"}).Set(0)
		return
	}

	// show as online
	err = dg.Open()
	if err != nil {
		logger.Errorf("error opening discord connection: %s\n", err)
		lastUpdate.With(prometheus.Labels{"type": "board", "ticker": b.Name, "guild": "None"}).Set(0)
		return
	}

	// Get guides for bot
	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		logger.Errorf("Error getting guilds: %s\n", err)
		b.Nickname = false
	}
	if len(guilds) == 0 {
		b.Nickname = false
	}

	// perform management operations
	if *managed {
		setName(dg, b.label())
	}

	logger.Infof("Watching board for %s", b.Name)
	ticker := time.NewTicker(time.Duration(b.Frequency) * time.Second)

	// continuously watch
	for {
		for _, symbol := range b.Items {
			select {
			case <-b.Close:
				logger.Infof("Shutting down price watching for %s", b.Name)
				return
			case <-ticker.C:

				logger.Debugf("Fetching crypto price for %s", symbol)

				var priceData utils.GeckoPriceResults
				var fmtPrice string
				var fmtDiff string

				// save the price struct & do something with it
				if rdb == nilCache {
					priceData, err = utils.GetCryptoPrice(symbol)
				} else {
					priceData, err = utils.GetCryptoPriceCache(rdb, ctx, symbol)
				}
				if err != nil {
					logger.Errorf("Unable to fetch stock price for %s: %s", symbol, err)
				}

				var change float64
				var activityHeader string
				var activityFooter string

				if b.Percentage {
					change = priceData.MarketData.PriceChangePercent
					activityHeader = ""
					activityFooter = "%"
				} else {
					change = priceData.MarketData.CurrentPrice.USD
					activityHeader = "$"
					activityFooter = ""
				}

				// Check for cryptos below 1c
				if priceData.MarketData.CurrentPrice.USD < 0.01 {
					fmtPrice = fmt.Sprintf("%.4f", priceData.MarketData.CurrentPrice.USD)
					fmtDiff = fmt.Sprintf("%.4f", change)
				} else if priceData.MarketData.CurrentPrice.USD < 1.0 {
					fmtPrice = fmt.Sprintf("%.3f", priceData.MarketData.CurrentPrice.USD)
					fmtDiff = fmt.Sprintf("%.3f", change)
				} else {
					fmtPrice = fmt.Sprintf("%.2f", priceData.MarketData.CurrentPrice.USD)
					fmtDiff = fmt.Sprintf("%.2f", change)
				}

				// calculate if price has moved up or down
				var increase bool
				if len(fmtDiff) == 0 {
					increase = true
				} else if string(fmtDiff[0]) == "-" {
					increase = false
				} else {
					increase = true
				}

				decorator := "⬊"
				if increase {
					decorator = "⬈"
				}

				if !b.Arrows {
					decorator = "-"
				}

				if b.Nickname {
					// update nickname instead of activity
					var nickname string
					var activity string

					displayName := b.Header + strings.ToUpper(priceData.Symbol)

					// format nickname
					nickname = fmt.Sprintf("%s %s $%s", displayName, decorator, fmtPrice)

					// format activity
					activity = fmt.Sprintf("24hr: %s%s%s", activityHeader, fmtDiff, activityFooter)

					// Update nickname in guilds
					for _, g := range guilds {
						err = dg.GuildMemberNickname(g.ID, "@me", nickname)
						if err != nil {
							logger.Errorf("Error updating nickname: %s\n", err)
							continue
						}
						logger.Infof("Set nickname in %s: %s", g.Name, nickname)
						lastUpdate.With(prometheus.Labels{"type": "board", "ticker": b.Name, "guild": g.Name}).SetToCurrentTime()

						// change bot color
						err = setRole(dg, b.ClientID, g.ID, increase)
						if err != nil {
							logger.Errorf("Color roles: %s", err)
						}
					}

					err = dg.UpdateGameStatus(0, b.Name)
					if err != nil {
						logger.Errorf("Unable to set activity: %s\n", err)
					} else {
						logger.Infof("Set activity: %s", activity)
					}

				} else {

					// format activity
					activity := fmt.Sprintf("%s $%s %s %s", strings.ToUpper(priceData.Symbol), fmtPrice, decorator, fmtDiff)
					err = dg.UpdateGameStatus(0, activity)
					if err != nil {
						logger.Errorf("Unable to set activity: %s\n", err)
					} else {
						logger.Infof("Set activity: %s", activity)
						lastUpdate.With(prometheus.Labels{"type": "board", "ticker": b.Name}).SetToCurrentTime()
					}
				}
			}
		}
	}
}
