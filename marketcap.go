package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/go-redis/redis/v8"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/rssnyder/discord-stock-ticker/utils"
)

type MarketCap struct {
	Ticker         string   `json:"ticker"`
	Name           string   `json:"name"`
	Nickname       bool     `json:"nickname"`
	Frequency      int      `json:"frequency"`
	Color          bool     `json:"color"`
	Decorator      string   `json:"decorator"`
	Currency       string   `json:"currency"`
	CurrencySymbol string   `json:"currency_symbol"`
	Decimals       int      `json:"decimals"`
	Activity       string   `json:"activity"`
	ClientID       string   `json:"client_id"`
	Token          string   `json:"discord_bot_token"`
	Close          chan int `json:"-"`
}

func (m *MarketCap) watchMarketCap() {
	var nilCache *redis.Client
	var exRate float64

	// create a new discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + m.Token)
	if err != nil {
		logger.Errorf("Creating Discord session: %s", err)
		lastUpdate.With(prometheus.Labels{"type": "marketcap", "ticker": m.Name, "guild": "None"}).Set(0)
		return
	}

	// show as online
	err = dg.Open()
	if err != nil {
		logger.Errorf("Opening discord connection: %s", err)
		lastUpdate.With(prometheus.Labels{"type": "marketcap", "ticker": m.Name, "guild": "None"}).Set(0)
		return
	}

	// get bot id
	botUser, err := dg.User("@me")
	if err != nil {
		logger.Errorf("Getting bot id: %s", err)
		lastUpdate.With(prometheus.Labels{"type": "marketcap", "ticker": m.Name, "guild": "None"}).Set(0)
		return
	}

	// Get guides for bot
	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		logger.Errorf("Getting guilds: %s", err)
		m.Nickname = false
	}

	// check for frequency override
	if *frequency != 0 {
		m.Frequency = *frequency
	}

	// If other currency, get rate
	if m.Currency != "USD" {
		exData, err := utils.GetStockPrice(m.Currency + "=X")
		if err != nil {
			logger.Errorf("Unable to fetch exchange rate for %s, default to USD.", m.Currency)
		} else {
			exRate = exData.QuoteSummary.Results[0].Price.RegularMarketPrice.Raw
		}
	}

	// Set arrows if no custom decorator
	var arrows bool
	if m.Decorator == "" {
		arrows = true
	}

	// Grab custom activity messages
	var custom_activity []string
	itr := 0
	itrSeed := 0.0
	if m.Activity != "" {
		custom_activity = strings.Split(m.Activity, ";")
	}

	logger.Infof("Watching marketcap for %s", m.Name)
	ticker := time.NewTicker(time.Duration(m.Frequency) * time.Second)

	// continuously watch
	for {
		select {
		case <-m.Close:
			logger.Infof("Shutting down price watching for %s", m.Name)
			return
		case <-ticker.C:
			logger.Debugf("Fetching crypto price for %s", m.Name)

			var priceData utils.GeckoPriceResults
			var fmtPrice string
			var fmtChange string
			var changeHeader string
			var fmtDiffPercent string

			// get the coin price data
			if rdb == nilCache {
				priceData, err = utils.GetCryptoPrice(m.Name)
			} else {
				priceData, err = utils.GetCryptoPriceCache(rdb, ctx, m.Name)
			}
			if err != nil {
				logger.Errorf("Unable to fetch stock price for %s: %s", m.Name, err)
				continue
			}

			// check for no mc data
			if priceData.MarketData.MarketCap.USD != 0 {
				logger.Debug("Using marketcap")
			} else if priceData.MarketData.CirculatingSupply != 0 {
				logger.Debug("Using circulating supply")
				priceData.MarketData.MarketCap.USD = priceData.MarketData.CirculatingSupply * priceData.MarketData.CurrentPrice.USD
			} else if priceData.MarketData.TotalSupply != 0 {
				logger.Debug("Using total supply")
				priceData.MarketData.MarketCap.USD = priceData.MarketData.TotalSupply * priceData.MarketData.CurrentPrice.USD
			} else {
				logger.Warning("No sources found for marketcap")
			}

			// Check if conversion is needed
			if exRate != 0 {
				priceData.MarketData.MarketCap.USD = exRate * priceData.MarketData.MarketCap.USD
				priceData.MarketData.MarketCapChangeCurrency.USD = exRate * priceData.MarketData.MarketCapChangeCurrency.USD
			}

			// format the price changes
			fmtDiffPercent = fmt.Sprintf("%.2f", priceData.MarketData.PriceChangePercent)
			fmtChange = fmt.Sprintf("%.2f", priceData.MarketData.PriceChangeCurrency.USD)

			// Check for custom decimal places
			p := message.NewPrinter(language.English)
			switch m.Decimals {
			case 1:
				fmtPrice = p.Sprintf("%s%.1f", m.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			case 2:
				fmtPrice = p.Sprintf("%s%.2f", m.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			case 3:
				fmtPrice = p.Sprintf("%s%.3f", m.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			case 4:
				fmtPrice = p.Sprintf("%s%.4f", m.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			case 5:
				fmtPrice = p.Sprintf("%s%.5f", m.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			case 6:
				fmtPrice = p.Sprintf("%s%.6f", m.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			case 7:
				fmtPrice = p.Sprintf("%s%.7f", m.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			case 8:
				fmtPrice = p.Sprintf("%s%.8f", m.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			case 9:
				fmtPrice = p.Sprintf("%s%.9f", m.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			case 10:
				fmtPrice = p.Sprintf("%s%.10f", m.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			case 11:
				fmtPrice = p.Sprintf("%s%.11f", m.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			default:
				fmtPrice = p.Sprintf("%s%.2f", m.CurrencySymbol, priceData.MarketData.MarketCap.USD)
				switch {
				case priceData.MarketData.MarketCap.USD < 1000000:
					fmtPrice = p.Sprintf("%s%.2fk", m.CurrencySymbol, priceData.MarketData.MarketCap.USD/1000)
				case priceData.MarketData.MarketCap.USD < 1000000000:
					fmtPrice = p.Sprintf("%s%.2fM", m.CurrencySymbol, priceData.MarketData.MarketCap.USD/1000000)
				case priceData.MarketData.MarketCap.USD < 1000000000000:
					fmtPrice = p.Sprintf("%s%.2fB", m.CurrencySymbol, priceData.MarketData.MarketCap.USD/1000000000)
				case priceData.MarketData.MarketCap.USD < 1000000000000000:
					fmtPrice = p.Sprintf("%s%.2fT", m.CurrencySymbol, priceData.MarketData.MarketCap.USD/1000000000000)
				}
			}

			// calculate if price has moved up or down
			var increase bool
			if len(fmtChange) == 0 {
				increase = true
			} else if string(fmtChange[0]) == "-" {
				increase = false
			} else {
				increase = true
			}

			// set arrows based on movement
			if arrows {
				m.Decorator = "⬊"
				if increase {
					m.Decorator = "⬈"
				}
			}

			// update nickname instead of activity
			if m.Nickname {
				var displayName string
				var nickname string
				var activity string

				// override coin symbol
				if m.Ticker != "" {
					displayName = m.Ticker
				} else {
					displayName = strings.ToUpper(priceData.Symbol)
				}

				// format nickname
				if displayName == m.Decorator {
					nickname = fmtPrice
				} else {
					nickname = fmt.Sprintf("%s %s %s", displayName, m.Decorator, fmtPrice)
				}

				// format activity
				activity = fmt.Sprintf("%s%s (%s%%)", changeHeader, fmtChange, fmtDiffPercent)

				// Update nickname in guilds
				for _, g := range guilds {
					err = dg.GuildMemberNickname(g.ID, "@me", nickname)
					if err != nil {
						logger.Errorf("Updating nickname: %s", err)
						continue
					}
					logger.Debugf("Set nickname in %s: %s", g.Name, nickname)
					lastUpdate.With(prometheus.Labels{"type": "marketcap", "ticker": m.Name, "guild": g.Name}).SetToCurrentTime()

					// change coin color
					if m.Color {
						var redRole string
						var greeenRole string

						// get the roles for color changing
						roles, err := dg.GuildRoles(g.ID)
						if err != nil {
							logger.Errorf("Getting guilds: %s", err)
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

						// make sure roles exist
						if len(redRole) == 0 || len(greeenRole) == 0 {
							logger.Error("Unable to find roles for color changes")
							continue
						}

						// assign role based on change
						if increase {
							err = dg.GuildMemberRoleRemove(g.ID, botUser.ID, redRole)
							if err != nil {
								logger.Errorf("Unable to remove role: %s", err)
							}
							err = dg.GuildMemberRoleAdd(g.ID, botUser.ID, greeenRole)
							if err != nil {
								logger.Errorf("Unable to set role: %s", err)
							}
						} else {
							err = dg.GuildMemberRoleRemove(g.ID, botUser.ID, greeenRole)
							if err != nil {
								logger.Errorf("Unable to remove role: %s", err)
							}
							err = dg.GuildMemberRoleAdd(g.ID, botUser.ID, redRole)
							if err != nil {
								logger.Errorf("Unable to set role: %s", err)
							}
						}
					}
					time.Sleep(time.Duration(m.Frequency) * time.Second)
				}

				// Custom activity messages
				if len(custom_activity) > 0 {

					// Display the real activity once per cycle
					if itr == len(custom_activity) {
						itr = 0
						itrSeed = 0.0
					} else if math.Mod(itrSeed, 2.0) == 1.0 {
						activity = custom_activity[itr]
						itr++
						itrSeed++
					} else {
						activity = custom_activity[itr]
						itrSeed++
					}
				}

				// set activity
				err = dg.UpdateGameStatus(0, activity)
				if err != nil {
					logger.Errorf("Unable to set activity: %s", err)
				} else {
					logger.Debugf("Set activity: %s", activity)
				}

			} else {

				// format activity
				activity := fmt.Sprintf("%s %s %s%%", fmtPrice, m.Decorator, fmtDiffPercent)
				err = dg.UpdateGameStatus(0, activity)
				if err != nil {
					logger.Errorf("Unable to set activity: %s", err)
				} else {
					logger.Debugf("Set activity: %s", activity)
					lastUpdate.With(prometheus.Labels{"type": "marketcap", "ticker": m.Name, "guild": "None"}).SetToCurrentTime()
				}
			}
		}
	}
}
