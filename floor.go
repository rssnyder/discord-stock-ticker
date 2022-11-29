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
	Marketplace string   `json:"marketplace"`
	Name        string   `json:"name"`
	Nickname    bool     `json:"nickname"`
	Activity    string   `json:"activity"`
	Frequency   int      `json:"frequency"`
	Color       bool     `json:"color"`
	ClientID    string   `json:"client_id"`
	Token       string   `json:"discord_bot_token"`
	close       chan int `json:"-"`
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
<<<<<<< HEAD
	var increase bool
=======
>>>>>>> feat: add floor colors argument
	for {

		select {
		case <-f.close:
			logger.Infof("Shutting down price watching for %s/%s", f.Marketplace, f.Name)
			return
		case <-ticker.C:
			price, activity, err := utils.GetFloorPrice(f.Marketplace, f.Name)
			if err != nil {
				logger.Errorf("Error getting floor rates: %s\n", err)
				continue
			}

			// calculate if price has moved up or down
			var priceRaw float64
			if priceRaw, err = strconv.ParseFloat(strings.ReplaceAll(price, " SOL", ""), 64); err != nil {
				logger.Errorf("Error with price format for %s", f.Name)
				continue
			}
			if priceRaw > oldPrice {
				increase = true
			} else if priceRaw < oldPrice {
				increase = false
			}

			// change nickname
			if f.Nickname {

				// Update nickname in guilds
				for _, g := range guilds {
					err = dg.GuildMemberNickname(g.ID, "@me", price)
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

					if f.Color {
						// get roles for colors
						var redRole string
						var greeenRole string

						roles, err := dg.GuildRoles(gu.ID)
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
							err = dg.GuildMemberRoleRemove(gu.ID, f.ClientID, redRole)
							if err != nil {
								logger.Errorf("Unable to remove role: %s", err)
							}
							err = dg.GuildMemberRoleAdd(gu.ID, f.ClientID, greeenRole)
							if err != nil {
								logger.Errorf("Unable to set role: %s", err)
							}
						} else {
							err = dg.GuildMemberRoleRemove(gu.ID, f.ClientID, greeenRole)
							if err != nil {
								logger.Errorf("Unable to remove role: %s", err)
							}
							err = dg.GuildMemberRoleAdd(gu.ID, f.ClientID, redRole)
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
						f.Activity = activity
					} else if math.Mod(itrSeed, 2.0) == 1.0 {
						f.Activity = custom_activity[itr]
						itr++
						itrSeed++
					} else {
						f.Activity = custom_activity[itr]
						itrSeed++
					}
				}

				err = dg.UpdateGameStatus(0, f.Activity)
				if err != nil {
					logger.Errorf("Unable to set activity: %s\n", err)
				} else {
					logger.Debugf("Set activity: %s", f.Activity)
				}
			} else {

				err = dg.UpdateWatchStatus(0, price)
				if err != nil {
					logger.Errorf("Unable to set activity: %s\n", err)
				} else {
					logger.Debugf("Set activity: %s\n", price)
					lastUpdate.With(prometheus.Labels{"type": "floor", "ticker": f.Name, "guild": "None"}).SetToCurrentTime()
				}
			}
			oldPrice = priceRaw
		}
	}
}
