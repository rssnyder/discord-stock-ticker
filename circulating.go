package main

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	"github.com/rssnyder/discord-stock-ticker/utils"
)

type Circulating struct {
	Ticker         string   `json:"ticker"`
	Name           string   `json:"name"`
	Nickname       bool     `json:"nickname"`
	Frequency      int      `json:"frequency"`
	CurrencySymbol string   `json:"currency_symbol"`
	Decimals       int      `json:"decimals"`
	Activity       string   `json:"activity"`
	ClientID       string   `json:"client_id"`
	Token          string   `json:"discord_bot_token"`
	close          chan int `json:"-"`
}

// label returns a human readble id for this bot
func (c *Circulating) label() string {
	label := strings.ToLower(fmt.Sprintf("%s-%s", c.Name, c.CurrencySymbol))
	if len(label) > 32 {
		label = label[:32]
	}
	return label
}

// Shutdown sends a signal to shut off the goroutine
func (c *Circulating) Shutdown() {
	c.close <- 1
}

func (c *Circulating) watchCirculating() {

	// create a new discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + c.Token)
	if err != nil {
		logger.Errorf("Creating Discord session (%s): %s", c.ClientID, err)
		lastUpdate.With(prometheus.Labels{"type": "circulating", "ticker": c.Name, "guild": "None"}).Set(0)
		return
	}

	// show as online
	err = dg.Open()
	if err != nil {
		logger.Errorf("Opening discord connection (%s): %s", c.ClientID, err)
		lastUpdate.With(prometheus.Labels{"type": "circulating", "ticker": c.Name, "guild": "None"}).Set(0)
		return
	}

	// Get guides for bot
	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		logger.Errorf("Getting guilds: %s", err)
		c.Nickname = false
	}
	if len(guilds) == 0 {
		c.Nickname = false
	}

	// check for frequency override
	if *frequency != 0 {
		c.Frequency = *frequency
	}

	// Grab custom activity messages
	var custom_activity []string
	itr := 0
	itrSeed := 0.0
	if c.Activity != "" {
		custom_activity = strings.Split(c.Activity, ";")
	}

	// perform management operations
	if *managed {
		setName(dg, c.label())
	}

	logger.Infof("Watching circulating for %s", c.Name)
	ticker := time.NewTicker(time.Duration(c.Frequency) * time.Second)
	var success bool

	c.close = make(chan int, 1)

	// continuously watch
	for {
		select {
		case <-c.close:
			logger.Infof("Shutting down price watching for %s", c.Name)
			return
		case <-ticker.C:
			logger.Debugf("Fetching crypto price for %s", c.Name)

			var priceData utils.GeckoPriceResults
			var fmtPrice string

			// get the coin price data
			if *cache {
				priceData, success, err = utils.GetCryptoPriceCache(rdb, ctx, c.Name)
				if success {
					cacheHits.Inc()
				} else {
					cacheMisses.Inc()
				}
			} else {
				priceData, err = utils.GetCryptoPrice(c.Name)
			}
			if err != nil {
				logger.Errorf("Unable to fetch circulating for %s: %s", c.Name, err)
				continue
			}

			// set currency
			if c.CurrencySymbol == "" {
				c.CurrencySymbol = strings.ToUpper(priceData.Symbol)
			}

			// Check for custom decimal places
			p := message.NewPrinter(language.English)
			switch c.Decimals {
			case 1:
				fmtPrice = p.Sprintf("%.1f %s", priceData.MarketData.CirculatingSupply, c.CurrencySymbol)
			case 2:
				fmtPrice = p.Sprintf("%.2f %s", priceData.MarketData.CirculatingSupply, c.CurrencySymbol)
			case 3:
				fmtPrice = p.Sprintf("%.3f %s", priceData.MarketData.CirculatingSupply, c.CurrencySymbol)
			case 4:
				fmtPrice = p.Sprintf("%.4f %s", priceData.MarketData.CirculatingSupply, c.CurrencySymbol)
			case 5:
				fmtPrice = p.Sprintf("%.5f %s", priceData.MarketData.CirculatingSupply, c.CurrencySymbol)
			case 6:
				fmtPrice = p.Sprintf("%.6f %s", priceData.MarketData.CirculatingSupply, c.CurrencySymbol)
			case 7:
				fmtPrice = p.Sprintf("%.7f %s", priceData.MarketData.CirculatingSupply, c.CurrencySymbol)
			case 8:
				fmtPrice = p.Sprintf("%.8f %s", priceData.MarketData.CirculatingSupply, c.CurrencySymbol)
			case 9:
				fmtPrice = p.Sprintf("%.9f %s", priceData.MarketData.CirculatingSupply, c.CurrencySymbol)
			case 10:
				fmtPrice = p.Sprintf("%.10f %s", priceData.MarketData.CirculatingSupply, c.CurrencySymbol)
			case 11:
				fmtPrice = p.Sprintf("%.11f %s", priceData.MarketData.CirculatingSupply, c.CurrencySymbol)
			default:
				fmtPrice = p.Sprintf("%.2f %s", priceData.MarketData.CirculatingSupply, c.CurrencySymbol)
				switch {
				case priceData.MarketData.CirculatingSupply < 1000000:
					fmtPrice = p.Sprintf("%.2fk %s", priceData.MarketData.CirculatingSupply/1000, c.CurrencySymbol)
				case priceData.MarketData.CirculatingSupply < 1000000000:
					fmtPrice = p.Sprintf("%.2fM %s", priceData.MarketData.CirculatingSupply/1000000, c.CurrencySymbol)
				case priceData.MarketData.CirculatingSupply < 1000000000000:
					fmtPrice = p.Sprintf("%.2fB %s", priceData.MarketData.CirculatingSupply/1000000000, c.CurrencySymbol)
				case priceData.MarketData.CirculatingSupply < 1000000000000000:
					fmtPrice = p.Sprintf("%.2fT %s", priceData.MarketData.CirculatingSupply/1000000000000, c.CurrencySymbol)
				}
			}

			// update nickname instead of activity
			if c.Nickname {
				nickname := fmt.Sprintf("%s %s", c.Ticker, fmtPrice)

				// Update nickname in guilds
				for _, g := range guilds {
					err = dg.GuildMemberNickname(g.ID, "@me", nickname)
					if err != nil {
						logger.Errorf("Updating nickname: %s", err)
						continue
					}
					logger.Debugf("Set nickname in %s: %s", g.Name, nickname)
					lastUpdate.With(prometheus.Labels{"type": "circulating", "ticker": c.Name, "guild": g.Name}).SetToCurrentTime()
					time.Sleep(time.Duration(c.Frequency) * time.Second)
				}

				// Custom activity messages
				if len(custom_activity) > 0 {

					// Display the real activity once per cycle
					if itr == len(custom_activity) {
						itr = 0
						itrSeed = 0.0
					} else if math.Mod(itrSeed, 2.0) == 1.0 {
						c.Activity = custom_activity[itr]
						itr++
						itrSeed++
					} else {
						c.Activity = custom_activity[itr]
						itrSeed++
					}
				}

				// set activity
				err = dg.UpdateWatchStatus(0, c.Activity)
				if err != nil {
					logger.Errorf("Unable to set activity: %s", err)
				} else {
					logger.Debugf("Set activity: %s", c.Activity)
				}

			} else {

				// format activity
				err = dg.UpdateWatchStatus(0, c.Activity)
				if err != nil {
					logger.Errorf("Unable to set activity: %s", err)
				} else {
					logger.Debugf("Set activity: %s", c.Activity)
					lastUpdate.With(prometheus.Labels{"type": "circulating", "ticker": c.Name, "guild": "None"}).SetToCurrentTime()
				}
			}
		}
	}
}
