package main

import (
	"fmt"
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
	Frequency   int      `json:"frequency"`
	ClientID    string   `json:"client_id"`
	Token       string   `json:"discord_bot_token"`
	Close       chan int `json:"-"`
}

// label returns a human readble id for this bot
func (f *Floor) label() string {
	return strings.ToLower(fmt.Sprintf("%s-%s", f.Marketplace, f.Name))
}

// watchFloorPrice gets floor prices and rotates through levels
func (f *Floor) watchFloorPrice() {

	// create a new discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + f.Token)
	if err != nil {
		logger.Errorf("Error creating Discord session: %s\n", err)
		lastUpdate.With(prometheus.Labels{"type": "floor", "ticker": f.Name, "guild": "None"}).Set(0)
		return
	}

	// show as online
	err = dg.Open()
	if err != nil {
		logger.Errorf("error opening discord connection: %s\n", err)
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
		f.Frequency = 600
	}

	// perform management operations
	if *managed {
		setName(dg, f.label())
	}

	logger.Infof("Watching floor price for %s/%s", f.Marketplace, f.Name)
	ticker := time.NewTicker(time.Duration(f.Frequency) * time.Second)

	// watch floor price
	for {

		select {
		case <-f.Close:
			logger.Infof("Shutting down price watching for %s/%s", f.Marketplace, f.Name)
			return
		case <-ticker.C:
			price, err := utils.GetFloorPrice(f.Marketplace, f.Name)
			if err != nil {
				logger.Errorf("Error getting floor rates: %s\n", err)
				continue
			}

			// change nickname
			if f.Nickname {

				for _, gu := range guilds {

					err = dg.GuildMemberNickname(gu.ID, "@me", price)
					if err != nil {
						logger.Errorf("Error updating nickname: %s\n", err)
						continue
					} else {
						logger.Debugf("Set nickname in %s: %s\n", gu.Name, price)
					}
					lastUpdate.With(prometheus.Labels{"type": "floor", "ticker": f.Name, "guild": gu.Name}).SetToCurrentTime()
					time.Sleep(time.Duration(f.Frequency) * time.Second)
				}

				err = dg.UpdateGameStatus(0, f.Name)
				if err != nil {
					logger.Errorf("Unable to set activity: %s\n", err)
				} else {
					logger.Debugf("Set activity: %s", f.Name)
				}
			} else {

				err = dg.UpdateGameStatus(0, price)
				if err != nil {
					logger.Errorf("Unable to set activity: %s\n", err)
				} else {
					logger.Debugf("Set activity: %s\n", price)
					lastUpdate.With(prometheus.Labels{"type": "floor", "ticker": f.Name, "guild": "None"}).SetToCurrentTime()
				}
			}
		}
	}
}
