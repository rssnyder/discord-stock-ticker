package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/rssnyder/discord-stock-ticker/utils"
)

// Holders represents the json for holders
type Holders struct {
	Network   string   `json:"network"`
	Address   string   `json:"address"`
	Activity  string   `json:"activity"`
	Nickname  bool     `json:"nickname"`
	Frequency int      `json:"frequency"`
	ClientID  string   `json:"client_id"`
	Token     string   `json:"discord_bot_token"`
	Close     chan int `json:"-"`
}

// label returns a human readble id for this bot
func (h *Holders) label() string {
	label := strings.ToLower(fmt.Sprintf("%s-%s", h.Network, h.Address))
	if len(label) > 32 {
		label = label[:32]
	}
	return label
}

func (h *Holders) watchHolders() {

	// create a new discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + h.Token)
	if err != nil {
		logger.Errorf("Error creating Discord session: %s\n", err)
		lastUpdate.With(prometheus.Labels{"type": "holders", "ticker": fmt.Sprintf("%s-%s", h.Network, h.Address), "guild": "None"}).Set(0)
		return
	}

	// show as online
	err = dg.Open()
	if err != nil {
		logger.Errorf("error opening discord connection: %s\n", err)
		lastUpdate.With(prometheus.Labels{"type": "holders", "ticker": fmt.Sprintf("%s-%s", h.Network, h.Address), "guild": "None"}).Set(0)
		return
	}

	// set activity as desc
	if h.Nickname {
		err = dg.UpdateGameStatus(0, h.Activity)
		if err != nil {
			logger.Errorf("Unable to set activity: %s\n", err)
		} else {
			logger.Debugf("Set activity")
		}
	}

	// get guides for bot
	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		logger.Errorf("Error getting guilds: %s\n", err)
		h.Nickname = false
	}
	if len(guilds) == 0 {
		h.Nickname = false
	}

	// check for frequency override
	// set to one hour to avoid lockout
	if *frequency != 0 {
		h.Frequency = 3600
	}

	// perform management operations
	if *managed {
		setName(dg, h.label())
	}

	logger.Infof("Watching holders for %s", h.Address)
	ticker := time.NewTicker(time.Duration(h.Frequency) * time.Second)

	for {

		select {
		case <-h.Close:
			logger.Infof("Shutting down price watching for %s", h.Activity)
			return
		case <-ticker.C:

			nickname := utils.GetHolders(h.Network, h.Address)

			if h.Nickname {

				for _, g := range guilds {

					err = dg.GuildMemberNickname(g.ID, "@me", nickname)
					if err != nil {
						logger.Errorf("Error updating nickname: %s\n", err)
						continue
					} else {
						logger.Debugf("Set nickname in %s: %s\n", g.Name, nickname)
					}
					lastUpdate.With(prometheus.Labels{"type": "holders", "ticker": fmt.Sprintf("%s-%s", h.Network, h.Address), "guild": g.Name}).SetToCurrentTime()
					time.Sleep(time.Duration(h.Frequency) * time.Second)
				}
			} else {

				err = dg.UpdateGameStatus(0, nickname)
				if err != nil {
					logger.Errorf("Unable to set activity: %s\n", err)
				} else {
					logger.Debugf("Set activity: %s\n", nickname)
					lastUpdate.With(prometheus.Labels{"type": "holders", "ticker": fmt.Sprintf("%s-%s", h.Network, h.Address), "guild": "None"}).SetToCurrentTime()
				}
			}
		}
	}
}
