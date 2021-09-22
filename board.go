package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/rssnyder/discord-stock-ticker/utils"
)

type Board struct {
	Items      []string             `json:"items"`
	Name       string               `json:"name"`
	Header     string               `json:"header"`
	Nickname   bool                 `json:"nickname"`
	Color      bool                 `json:"color"`
	Percentage bool                 `json:"percentage"`
	Arrows     bool                 `json:"arrows"`
	Frequency  int                  `json:"frequency"`
	ClientID   string               `json:"client_id"`
	Price      int                  `json:"-"`
	Cache      *redis.Client        `json:"-"`
	Context    context.Context      `json:"-"`
	updated    *prometheus.GaugeVec `json:"-"`
	token      string               `json:"-"`
	close      chan int             `json:"-"`
}

// NewBoard saves information about the board and starts up a watcher on it
func NewStockBoard(clientID string, items []string, token string, name string, header string, nickname bool, color bool, frequency int, updated *prometheus.GaugeVec) *Board {
	b := &Board{
		Items:     items,
		Name:      name,
		Header:    header,
		Nickname:  nickname,
		Color:     color,
		Frequency: frequency,
		ClientID:  clientID,
		updated:   updated,
		token:     token,
		close:     make(chan int, 1),
	}

	// spin off go routine to watch the price
	go b.watchStockPrice()
	return b
}

// NewCrypto saves information about the crypto and starts up a watcher on it
func NewCryptoBoard(clientID string, items []string, token string, name string, header string, nickname bool, color bool, frequency int, updated *prometheus.GaugeVec, cache *redis.Client, context context.Context) *Board {
	b := &Board{
		Items:     items,
		Name:      name,
		Header:    header,
		Nickname:  nickname,
		Color:     color,
		Frequency: frequency,
		ClientID:  clientID,
		Cache:     cache,
		Context:   context,
		updated:   updated,
		token:     token,
		close:     make(chan int, 1),
	}

	// spin off go routine to watch the price
	b.Start()
	return b
}

// Start begins a board
func (b *Board) Start() {
	go b.watchCryptoPrice()
}

// Shutdown sends a signal to shut off the goroutine
func (b *Board) Shutdown() {
	b.close <- 1
}

func (b *Board) watchStockPrice() {

	// create a new discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + b.token)
	if err != nil {
		logger.Errorf("Error creating Discord session: %s\n", err)
		return
	}

	// show as online
	err = dg.Open()
	if err != nil {
		logger.Errorf("error opening discord connection: %s\n", err)
		return
	}

	// get bot id
	botUser, err := dg.User("@me")
	if err != nil {
		logger.Errorf("Error getting bot id: %s\n", err)
		return
	}

	// Get guides for bot
	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		logger.Errorf("Error getting guilds: %s\n", err)
		b.Nickname = false
	}

	ticker := time.NewTicker(time.Duration(b.Frequency) * time.Second)

	// continuously watch
	for {
		for _, symbol := range b.Items {
			select {
			case <-b.close:
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
						b.updated.With(prometheus.Labels{"type": "board", "ticker": b.Name, "guild": g.Name}).SetToCurrentTime()

						if b.Color {
							// get roles for colors
							var redRole string
							var greeenRole string

							roles, err := dg.GuildRoles(g.ID)
							if err != nil {
								logger.Errorf("Error getting guilds: %s\n", err)
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
									logger.Errorf("Unable to remove role: %s\n", err)
								}
								err = dg.GuildMemberRoleAdd(g.ID, botUser.ID, greeenRole)
								if err != nil {
									logger.Errorf("Unable to set role: %s\n", err)
								}
							} else {
								err = dg.GuildMemberRoleRemove(g.ID, botUser.ID, greeenRole)
								if err != nil {
									logger.Errorf("Unable to remove role: %s\n", err)
								}
								err = dg.GuildMemberRoleAdd(g.ID, botUser.ID, redRole)
								if err != nil {
									logger.Errorf("Unable to set role: %s\n", err)
								}
							}
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
						b.updated.With(prometheus.Labels{"type": "board", "ticker": b.Name, "guild": "None"}).SetToCurrentTime()
					}
				}
			}
		}
	}
}

func (b *Board) watchCryptoPrice() {
	var rdb *redis.Client

	// create a new discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + b.token)
	if err != nil {
		logger.Errorf("Error creating Discord session: %s\n", err)
		return
	}

	// show as online
	err = dg.Open()
	if err != nil {
		logger.Errorf("error opening discord connection: %s\n", err)
		return
	}

	// get bot id
	botUser, err := dg.User("@me")
	if err != nil {
		logger.Errorf("Error getting bot id: %s\n", err)
		return
	}

	// Get guides for bot
	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		logger.Errorf("Error getting guilds: %s\n", err)
		b.Nickname = false
	}

	ticker := time.NewTicker(time.Duration(b.Frequency) * time.Second)
	logger.Debugf("Watching crypto price for %s", b.Name)

	// continuously watch
	for {
		for _, symbol := range b.Items {
			select {
			case <-b.close:
				logger.Infof("Shutting down price watching for %s", b.Name)
				return
			case <-ticker.C:

				logger.Debugf("Fetching crypto price for %s", symbol)

				var priceData utils.GeckoPriceResults
				var fmtPrice string
				var fmtDiff string

				// save the price struct & do something with it
				if b.Cache == rdb {
					priceData, err = utils.GetCryptoPrice(symbol)
				} else {
					priceData, err = utils.GetCryptoPriceCache(b.Cache, b.Context, symbol)
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
						b.updated.With(prometheus.Labels{"type": "board", "ticker": b.Name, "guild": g.Name}).SetToCurrentTime()

						if b.Color {
							// get roles for colors
							var redRole string
							var greeenRole string

							roles, err := dg.GuildRoles(g.ID)
							if err != nil {
								logger.Errorf("Error getting guilds: %s\n", err)
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
									logger.Errorf("Unable to remove role: %s\n", err)
								}
								err = dg.GuildMemberRoleAdd(g.ID, botUser.ID, greeenRole)
								if err != nil {
									logger.Errorf("Unable to set role: %s\n", err)
								}
							} else {
								err = dg.GuildMemberRoleRemove(g.ID, botUser.ID, greeenRole)
								if err != nil {
									logger.Errorf("Unable to remove role: %s\n", err)
								}
								err = dg.GuildMemberRoleAdd(g.ID, botUser.ID, redRole)
								if err != nil {
									logger.Errorf("Unable to set role: %s\n", err)
								}
							}
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
						b.updated.With(prometheus.Labels{"type": "board", "ticker": b.Name}).SetToCurrentTime()
					}
				}
			}
		}
	}
}
