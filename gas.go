package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rssnyder/discord-stock-ticker/utils"
)

// Gas represents the gas data
type Gas struct {
	Network   string        `json:"network"`
	Nickname  bool          `json:"set_nickname"`
	Frequency time.Duration `json:"frequency"`
	token     string        `json:"-"`
	close     chan int      `json:"-"`
}

func NewGas(network string, token string, nickname bool, frequency int) *Gas {
	g := &Gas{
		Network:   network,
		Nickname:  nickname,
		Frequency: time.Duration(frequency) * time.Second,
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
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	// show as online
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening discord connection,", err)
		return
	}

	// Get guides for bot
	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		fmt.Println("Error getting guilds: ", err)
		g.Nickname = false
	}

	ticker := time.NewTicker(g.Frequency)
	var nickname string

	// watch gas price
	for {

		select {
		case <-g.close:
			logger.Infof("Shutting down price watching for %s", g.Network)
			return
		case <-ticker.C:
			// get gas prices
			gasPrices, err := utils.GetGasPrices(g.Network)
			if err != nil {
				fmt.Printf("Error getting rates: %s\n", err)
				time.Sleep(g.Frequency)
				continue
			}

			nickname = fmt.Sprintf("⚡ %d 🤔 %d 🐌 %d", gasPrices.Instant, gasPrices.Fast, gasPrices.Standard)

			// change nickname
			if g.Nickname {

				for _, g := range guilds {

					err = dg.GuildMemberNickname(g.ID, "@me", nickname)
					if err != nil {
						fmt.Printf("Error updating nickname: %s\n", err)
						continue
					} else {
						fmt.Printf("Set nickname in %s: %s\n", g.Name, nickname)
					}
				}

				err = dg.UpdateGameStatus(0, "Fast, Avg, Slow")
				if err != nil {
					fmt.Printf("Unable to set activity: %s\n", err)
				} else {
					fmt.Println("Set activity")
				}
			} else {

				err = dg.UpdateGameStatus(0, nickname)
				if err != nil {
					fmt.Printf("Unable to set activity: %s\n", err)
				} else {
					fmt.Printf("Set activity: %s\n", nickname)
				}
			}
		}
	}
}
