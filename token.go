package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/rssnyder/discord-stock-ticker/utils"
)

type Token struct {
	Network   string   `json:"network"`
	Contract  string   `json:"contract"`
	Name      string   `json:"name"`
	Nickname  bool     `json:"nickname"`
	Frequency int      `json:"frequency"`
	Color     bool     `json:"color"`
	Decorator string   `json:"decorator"`
	Decimals  int      `json:"decimals"`
	Activity  string   `json:"activity"`
	Source    string   `json:"source"`
	ClientID  string   `json:"client_id"`
	Token     string   `json:"discord_bot_token"`
	close     chan int `json:"-"`
}

// label returns a human readble id for this bot
func (t *Token) label() string {
	label := strings.ToLower(fmt.Sprintf("%s-%s", t.Network, t.Contract))
	if len(label) > 32 {
		label = label[:32]
	}
	return label
}

// Format nickname
func formatNickname(t *Token, price float64) string {
	if price <= math.Pow10(-t.Decimals) {
		return fmt.Sprintf("%s %s $%.2e", t.Name, t.Decorator, price)
	} else {
		return fmt.Sprintf("%s %s $%."+strconv.Itoa(t.Decimals)+"f", t.Name, t.Decorator, price)
	}
}

// Format activity when not using nickname
func formatActivity(t *Token, price float64) string {
	if price <= math.Pow10(-t.Decimals) {
		return fmt.Sprintf("%s %s $%.2e", t.Name, t.Decorator, price)
	} else {
		return fmt.Sprintf("%s %s $%."+strconv.Itoa(t.Decimals)+"f", t.Name, t.Decorator, price)
	}
}

func (t *Token) Shutdown() {
	t.close <- 1
}

func (t *Token) watchTokenPrice() {

	// create a new discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + t.Token)
	if err != nil {
		logger.Errorf("Creating Discord session (%s): %s", t.ClientID, err)
		lastUpdate.With(prometheus.Labels{"type": "token", "ticker": fmt.Sprintf("%s-%s", t.Network, t.Contract), "guild": "None"}).Set(0)
		return
	}

	// show as online
	err = dg.Open()
	if err != nil {
		logger.Errorf("Opening discord connection (%s): %s", t.ClientID, err)
		lastUpdate.With(prometheus.Labels{"type": "token", "ticker": fmt.Sprintf("%s-%s", t.Network, t.Contract), "guild": "None"}).Set(0)
		return
	}

	// Get guides for bot
	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		logger.Errorf("Error getting guilds: %s\n", err)
		t.Nickname = false
	}
	if len(guilds) == 0 {
		t.Nickname = false
	}

	// check for frequency override
	if *frequency != 0 {
		t.Frequency = *frequency
	}

	// Set arrows if no custom decorator
	var arrows bool
	if t.Decorator == "" {
		arrows = true
	} else if t.Decorator == " " {
		t.Decorator = ""
	}

	// Grab custom activity messages
	var custom_activity []string
	itr := 0
	itrSeed := 0.0
	if t.Activity != "" {
		custom_activity = strings.Split(t.Activity, ";")
	}

	// perform management operations
	if *managed {
		setName(dg, t.label())
	}

	logger.Infof("Watching token price for %s", t.Name)
	ticker := time.NewTicker(time.Duration(t.Frequency) * time.Second)
	var success bool

	t.close = make(chan int, 1)

	// continuously watch
	var oldPrice float64
	var increase bool
	for {
		select {
		case <-t.close:
			logger.Infof("Shutting down price watching for %s", t.Name)
			return
		case <-ticker.C:
			logger.Debugf("Fetching token price for %s", t.Name)
			var priceData string
			var fmtPriceRaw float64
			var fmtPrice float64
			var bnbRate utils.GeckoPriceResults

			switch t.Source {
			case "pancakeswap":
				logger.Debugf("Using %s to get price: %s", t.Source, t.Name)

				// Get price from Ps in BNB
				priceData, err = utils.GetPancakeTokenPrice(t.Contract)
				if err != nil {
					logger.Errorf("Unable to fetch token price from %s: %s", t.Source, t.Name)
					continue
				}

				// get the bnb price
				if *cache {
					bnbRate, success, err = utils.GetCryptoPriceCache(rdb, ctx, "binancecoin")
					if !success || (err != nil) {
						cacheMisses.Inc()
					} else {
						cacheHits.Inc()
					}
				} else {
					bnbRate, err = utils.GetCryptoPrice("binancecoin")
				}
				if err != nil {
					logger.Errorf("Unable to fetch bnb price for %s", t.Name)
					continue
				}

				if fmtPriceRaw, err = strconv.ParseFloat(priceData, 64); err != nil {
					logger.Errorf("Error with price format for %s", t.Name)
					continue
				}
				fmtPrice = bnbRate.MarketData.CurrentPrice.USD * fmtPriceRaw

			case "dexlab":
				logger.Debugf("Using %s to get price: %s", t.Source, t.Name)

				// Get price from dexlab in USDT
				priceData, err = utils.GetDexLabPrice(t.Contract)
				if err != nil {
					logger.Errorf("Unable to fetch token price from %s: %s", t.Source, t.Name)
					continue
				}

				if fmtPrice, err = strconv.ParseFloat(priceData, 64); err != nil {
					logger.Errorf("Error with price format for %s", t.Name)
					continue
				}

			default:
				priceData, err = utils.Get1inchTokenPrice(t.Network, t.Contract)
				if err != nil {
					logger.Errorf("Unable to fetch token price for %s", t.Name)
					continue
				}

				if fmtPriceRaw, err = strconv.ParseFloat(priceData, 64); err != nil {
					logger.Errorf("Error with price format for %s", t.Name)
					continue
				}
				fmtPrice = fmtPriceRaw / 10000000
			}

			// calculate if price has moved up or down
			if fmtPrice > oldPrice {
				increase = true
			} else if fmtPrice < oldPrice {
				increase = false
			}

			if arrows {
				t.Decorator = "⬊"
				if increase {
					t.Decorator = "⬈"
				}
			}

			if t.Nickname {
				// update nickname instead of activity
				var nickname string
				var activity string

				// format nickname & activity
				// Check for custom decimal places
				nickname = formatNickname(t, fmtPrice)

				// Update nickname in guilds
				for _, g := range guilds {
					err = dg.GuildMemberNickname(g.ID, "@me", nickname)
					if err != nil {
						logger.Errorf("Error updating nickname: %s\n", err)
						continue
					}
					logger.Debugf("Set nickname in %s: %s", g.Name, nickname)
					lastUpdate.With(prometheus.Labels{"type": "token", "ticker": fmt.Sprintf("%s-%s", t.Network, t.Contract), "guild": g.Name}).SetToCurrentTime()

					if t.Color {
						// get roles for colors
						var redRole string
						var greeenRole string

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

						if len(redRole) == 0 || len(greeenRole) == 0 {
							logger.Error("Unable to find roles for color changes")
							continue
						}

						// assign role based on change
						if increase {
							err = dg.GuildMemberRoleRemove(g.ID, t.ClientID, redRole)
							if err != nil {
								logger.Errorf("Unable to remove role: %s", err)
							}
							err = dg.GuildMemberRoleAdd(g.ID, t.ClientID, greeenRole)
							if err != nil {
								logger.Errorf("Unable to set role: %s", err)
							}
						} else {
							err = dg.GuildMemberRoleRemove(g.ID, t.ClientID, greeenRole)
							if err != nil {
								logger.Errorf("Unable to remove role: %s", err)
							}
							err = dg.GuildMemberRoleAdd(g.ID, t.ClientID, redRole)
							if err != nil {
								logger.Errorf("Unable to set role: %s", err)
							}
						}
					}
					time.Sleep(time.Duration(t.Frequency) * time.Second)
				}

				activity = ""
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

				err = dg.UpdateWatchStatus(0, activity)
				if err != nil {
					logger.Error("Unable to set activity: ", err)
				} else {
					logger.Debugf("Set activity: %s", activity)
				}

			} else {
				activity := formatActivity(t, fmtPrice)

				err = dg.UpdateWatchStatus(0, activity)
				if err != nil {
					logger.Error("Unable to set activity: ", err)
				} else {
					logger.Debugf("Set activity: %s", activity)
					lastUpdate.With(prometheus.Labels{"type": "token", "ticker": fmt.Sprintf("%s-%s", t.Network, t.Contract), "guild": "None"}).SetToCurrentTime()
				}
			}
			oldPrice = fmtPrice
		}
	}
}
