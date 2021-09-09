package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rssnyder/discord-stock-ticker/utils"
)

// Gas represents the gas data
type Gas struct {
	Network   string   `json:"network"`
	Nickname  bool     `json:"set_nickname"`
	Frequency int      `json:"frequency"`
	ClientID  string   `json:"client_id"`
	token     string   `json:"-"`
	close     chan int `json:"-"`
}

func NewGas(clientID string, network string, token string, nickname bool, frequency int) *Gas {
	g := &Gas{
		Network:   network,
		Nickname:  nickname,
		Frequency: frequency,
		ClientID:  clientID,
		token:     token,
		close:     make(chan int, 1),
	}

	// spin off go routine to watch the prices
	go g.watchGasPrice()
	return g
}

// Shutdown sends a signal to shut off the goroutine
func (g *Gas) Shutdown() {
	g.close <- 1
}

// watchGasPrice gets gas prices and rotates through levels
func (g *Gas) watchGasPrice() {

	// create a new discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + g.token)
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

	// Get guides for bot
	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		logger.Errorf("Error getting guilds: %s\n", err)
		g.Nickname = false
	}

	ticker := time.NewTicker(time.Duration(g.Frequency) * time.Second)
	var nickname string

	// watch gas price
	for {

		select {
		case <-g.close:
			logger.Infof("Shutting down price watching for %s\n", g.Network)
			return
		case <-ticker.C:
			// get gas prices
			gasPrices, err := utils.GetGasPrices(g.Network)
			if err != nil {
				logger.Errorf("Error getting rates: %s\n", err)
				continue
			}

			nickname = fmt.Sprintf("⚡ %d 🤔 %d 🐌 %d", gasPrices.Instant, gasPrices.Fast, gasPrices.Standard)

			// change nickname
			if g.Nickname {

				for _, g := range guilds {

					err = dg.GuildMemberNickname(g.ID, "@me", nickname)
					if err != nil {
						logger.Errorf("Error updating nickname: %s\n", err)
						continue
					} else {
						logger.Debugf("Set nickname in %s: %s\n", g.Name, nickname)
					}
				}

				err = dg.UpdateGameStatus(0, "Fast, Avg, Slow")
				if err != nil {
					logger.Errorf("Unable to set activity: %s\n", err)
				} else {
					logger.Debugf("Set activity")
				}
			} else {

				err = dg.UpdateGameStatus(0, nickname)
				if err != nil {
					logger.Errorf("Unable to set activity: %s\n", err)
				} else {
					logger.Debugf("Set activity: %s\n", nickname)
				}
			}
		}
	}
}
