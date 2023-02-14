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

// Floor represents the floor data
type Floor struct {
	Marketplace           string   `json:"marketplace"`
	Name                  string   `json:"name"`
	Nickname              bool     `json:"nickname"`
	Activity              string   `json:"activity"`
	Frequency             int      `json:"frequency"`
	Color                 bool     `json:"color"`
	Decorator             string   `json:"decorator"`
	CollectionStats       bool     `json:"collection_stats"`
	CollectionStatsAppend bool     `json:"collection_stats_append"`
	Currency              string   `json:"currency"`
	ClientID              string   `json:"client_id"`
	Token                 string   `json:"discord_bot_token"`
	close                 chan int `json:"-"`
}

// label returns a human readble id for this bot
func (f *Floor) label() string {
	label := strings.ToLower(fmt.Sprintf("%s-%s", f.Marketplace, f.Name))
	if len(label) > 32 {
		label = label[:32]
	}
	return label
}

// Shutdown sends a signal to shut off the goroutine
func (f *Floor) Shutdown() {
	f.close <- 1
}

// watchFloorPrice gets floor prices and rotates through levels
func (f *Floor) watchFloorPrice() {

	// create a new discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + f.Token)
	if err != nil {
		logger.Errorf("Creating Discord session (%s): %s", f.ClientID, err)
		lastUpdate.With(prometheus.Labels{"type": "floor", "ticker": f.Name, "guild": "None"}).Set(0)
		return
	}

	// show as online
	err = dg.Open()
	if err != nil {
		logger.Errorf("Opening discord connection (%s): %s", f.ClientID, err)
		lastUpdate.With(prometheus.Labels{"type": "floor", "ticker": f.Name, "guild": "None"}).Set(0)
		return
	}

	// Get guides for bot
	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		logger.Errorf("Error getting guilds: %s\n", err)
		f.Nickname = false
	}
	if len(guilds) == 0 {
		f.Nickname = false
	}

	// check for frequency override
	// set to avoid lockout
	if *frequency != 0 {
		f.Frequency = 900
	}

	// Set arrows if no custom decorator
	var arrows bool
	if f.Decorator == "" {
		arrows = true
	} else if f.Decorator == " " { // Set to space to disable
		f.Decorator = ""
	}

	// Grab custom activity messages
	var custom_activity []string
	itr := 0
	itrSeed := 0.0
	if f.Activity != "" {
		custom_activity = strings.Split(f.Activity, ";")
	}

	// perform management operations
	if *managed {
		setName(dg, f.label())
	}

	logger.Infof("Watching floor price for %s/%s", f.Marketplace, f.Name)
	ticker := time.NewTicker(time.Duration(f.Frequency) * time.Second)

	f.close = make(chan int, 1)

	// watch floor price
	var oldPrice float64
	var increase bool
	var priceString string
	for {

		select {
		case <-f.close:
			logger.Infof("Shutting down price watching for %s/%s", f.Marketplace, f.Name)
			return
		case <-ticker.C:
			price, activity, currency, collectionStats, err := utils.GetFloorPrice(f.Marketplace, f.Name)
			if err != nil {
				logger.Errorf("Error getting floor rates: %s\n", err)
				continue
			}

			// Use platform currency if not set.
			if f.Currency == "" {
				f.Currency = currency
			}

			// calculate if price has moved up or down
			if price > oldPrice {
				increase = true
			} else if price < oldPrice {
				increase = false
			}

			// Add arrows to price if requested
			if arrows {
				f.Decorator = "⬊"
				if increase {
					f.Decorator = "⬈"
				}
			}

			// Convert price to string format.
			if f.Currency == "ETH" {
				priceString = fmt.Sprintf("%s Ξ%s", f.Decorator, strconv.FormatFloat(price, 'f', -1, 64))
			} else {
				priceString = fmt.Sprintf("%s %s %s", f.Decorator, strconv.FormatFloat(price, 'f', -1, 64), f.Currency)
			}

			// change nickname
			if f.Nickname {

				// Update nickname in guilds
				for _, g := range guilds {
					err = dg.GuildMemberNickname(g.ID, "@me", priceString)
					if err != nil {
						logger.Errorf("Updating nickname: %s", err)
						continue
					}
					logger.Debugf("Set nickname in %s: %s", g.Name, price)
					lastUpdate.With(prometheus.Labels{"type": "floor", "ticker": f.Name, "guild": g.Name}).SetToCurrentTime()

					if f.Color {
						// change bot color
						err = setRole(dg, f.ClientID, g.ID, increase)
						if err != nil {
							logger.Errorf("Color roles: %s", err)
						}
					}

					time.Sleep(time.Duration(f.Frequency) * time.Second)
				}

				// Add collection stats if requested.
				var activityStats = ""
				var customActivityTemp []string = custom_activity
				if f.CollectionStats {
					if len(collectionStats) > 0 {
						// Append to activity if requested.
						if f.CollectionStatsAppend {
							activityStats = ": " + collectionStats
						} else {
							customActivityTemp = append(customActivityTemp, collectionStats)
						}
					}
				}

				// Custom activity messages
				if len(customActivityTemp) > 0 {

					// Display the real activity once per cycle
					if itr == len(customActivityTemp) {
						itr = 0
						itrSeed = 0.0
						f.Activity = activity + activityStats
					} else if math.Mod(itrSeed, 2.0) == 1.0 { // Display custom activities 2 times per cycle.
						f.Activity = customActivityTemp[itr] + activityStats
						itr++
						itrSeed++
					} else {
						f.Activity = customActivityTemp[itr] + activityStats
						itrSeed++
					}
				} else {
					f.Activity = fmt.Sprintf("%s: %s", activity, collectionStats)
				}

				err = dg.UpdateWatchStatus(0, f.Activity)
				if err != nil {
					logger.Errorf("Unable to set activity: %s\n", err)
				} else {
					logger.Debugf("Set activity: %s", f.Activity)
				}
			} else {
				if f.CollectionStats {
					activity = fmt.Sprintf("%s | %s", priceString, collectionStats)
				} else {
					activity = priceString
				}

				err = dg.UpdateGameStatus(0, activity)
				if err != nil {
					logger.Errorf("Unable to set activity: %s\n", err)
				} else {
					logger.Debugf("Set activity: %s\n", price)
					lastUpdate.With(prometheus.Labels{"type": "floor", "ticker": f.Name, "guild": "None"}).SetToCurrentTime()
				}
			}
			oldPrice = price
		}
	}
}
