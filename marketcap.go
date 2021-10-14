package main

import (
	"context"
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
	Ticker         string               `json:"ticker"`
	Name           string               `json:"name"`
	Nickname       bool                 `json:"nickname"`
	Frequency      int                  `json:"frequency"`
	Color          bool                 `json:"color"`
	Decorator      string               `json:"decorator"`
	Currency       string               `json:"currency"`
	CurrencySymbol string               `json:"currency_symbol"`
	Decimals       int                  `json:"decimals"`
	Activity       string               `json:"activity"`
	ClientID       string               `json:"client_id"`
	Cache          *redis.Client        `json:"-"`
	Context        context.Context      `json:"-"`
	updated        *prometheus.GaugeVec `json:"-"`
	token          string               `json:"-"`
	close          chan int             `json:"-"`
}

// NewMarketCap saves information about the crypto and starts up a watcher on it
func NewMarketCap(clientID string, ticker string, token string, name string, nickname bool, color bool, decorator string, frequency int, currency string, activity string, decimals int, currencySymbol string, updated *prometheus.GaugeVec, cache *redis.Client, context context.Context) *MarketCap {
	s := &MarketCap{
		Ticker:         ticker,
		Name:           name,
		Nickname:       nickname,
		Color:          color,
		Decorator:      decorator,
		Activity:       activity,
		Decimals:       decimals,
		Frequency:      frequency,
		Currency:       strings.ToUpper(currency),
		CurrencySymbol: currencySymbol,
		ClientID:       clientID,
		Cache:          cache,
		Context:        context,
		updated:        updated,
		token:          token,
		close:          make(chan int, 1),
	}

	// spin off go routine to watch the price
	s.Start()
	return s
}

// Start begins watching a ticker
func (s *MarketCap) Start() {
	go s.watchMarketCap()
}

// Shutdown sends a signal to shut off the goroutine
func (s *MarketCap) Shutdown() {
	s.close <- 1
}

func (s *MarketCap) watchMarketCap() {
	var rdb *redis.Client
	var exRate float64

	// create a new discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + s.token)
	if err != nil {
		logger.Errorf("Creating Discord session: %s", err)
		return
	}

	// show as online
	err = dg.Open()
	if err != nil {
		logger.Errorf("Opening discord connection: %s", err)
		return
	}

	// get bot id
	botUser, err := dg.User("@me")
	if err != nil {
		logger.Errorf("Getting bot id: %s", err)
		return
	}

	// Get guides for bot
	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		logger.Errorf("Getting guilds: %s", err)
		s.Nickname = false
	}

	// check for frequency override
	if *frequency != 0 {
		s.Frequency = *frequency
	}

	// If other currency, get rate
	if s.Currency != "USD" {
		exData, err := utils.GetStockPrice(s.Currency + "=X")
		if err != nil {
			logger.Errorf("Unable to fetch exchange rate for %s, default to USD.", s.Currency)
		} else {
			exRate = exData.QuoteSummary.Results[0].Price.RegularMarketPrice.Raw
		}
	}

	// Set arrows if no custom decorator
	var arrows bool
	if s.Decorator == "" {
		arrows = true
	}

	// Grab custom activity messages
	var custom_activity []string
	itr := 0
	itrSeed := 0.0
	if s.Activity != "" {
		custom_activity = strings.Split(s.Activity, ";")
	}

	// create timer
	ticker := time.NewTicker(time.Duration(s.Frequency) * time.Second)
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
			var fmtChange string
			var changeHeader string
			var fmtDiffPercent string

			// get the coin price data
			if s.Cache == rdb {
				priceData, err = utils.GetCryptoPrice(s.Name)
			} else {
				priceData, err = utils.GetCryptoPriceCache(s.Cache, s.Context, s.Name)
			}
			if err != nil {
				logger.Errorf("Unable to fetch stock price for %s: %s", s.Name, err)
				continue
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
			switch s.Decimals {
			case 1:
				fmtPrice = p.Sprintf("%s%.1f", s.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			case 2:
				fmtPrice = p.Sprintf("%s%.2f", s.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			case 3:
				fmtPrice = p.Sprintf("%s%.3f", s.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			case 4:
				fmtPrice = p.Sprintf("%s%.4f", s.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			case 5:
				fmtPrice = p.Sprintf("%s%.5f", s.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			case 6:
				fmtPrice = p.Sprintf("%s%.6f", s.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			case 7:
				fmtPrice = p.Sprintf("%s%.7f", s.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			case 8:
				fmtPrice = p.Sprintf("%s%.8f", s.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			case 9:
				fmtPrice = p.Sprintf("%s%.9f", s.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			case 10:
				fmtPrice = p.Sprintf("%s%.10f", s.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			case 11:
				fmtPrice = p.Sprintf("%s%.11f", s.CurrencySymbol, priceData.MarketData.MarketCap.USD)
			default:
				fmtPrice = p.Sprintf("%s%.2f", s.CurrencySymbol, priceData.MarketData.MarketCap.USD)
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
				s.Decorator = "⬊"
				if increase {
					s.Decorator = "⬈"
				}
			}

			// update nickname instead of activity
			if s.Nickname {
				var displayName string
				var nickname string
				var activity string

				// override coin symbol
				if s.Ticker != "" {
					displayName = s.Ticker
				} else {
					displayName = strings.ToUpper(priceData.Symbol)
				}

				// format nickname
				nickname = fmt.Sprintf("%s %s %s", displayName, s.Decorator, fmtPrice)

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
					s.updated.With(prometheus.Labels{"type": "marketcap", "ticker": s.Name, "guild": g.Name}).SetToCurrentTime()

					// change coin color
					if s.Color {
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
				activity := fmt.Sprintf("%s %s %s%%", fmtPrice, s.Decorator, fmtDiffPercent)
				err = dg.UpdateGameStatus(0, activity)
				if err != nil {
					logger.Errorf("Unable to set activity: %s", err)
				} else {
					logger.Debugf("Set activity: %s", activity)
					s.updated.With(prometheus.Labels{"type": "marketcap", "ticker": s.Name, "guild": "None"}).SetToCurrentTime()
				}
			}
		}
	}
}
