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

	// Set activity as desc
	if g.Nickname {
		err = dg.UpdateListeningStatus(fmt.Sprintf("%s gas price", g.Network))
		if err != nil {
			fmt.Printf("Unable to set activity: \n", err)
		} else {
			fmt.Println("Set activity")
		}
	}

	// Get guides for bot
	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		fmt.Println("Error getting guilds: ", err)
		g.Nickname = false
	}

	changeFrequency := time.Duration(g.Frequency) * time.Second
	var nickname string

	// watch gas price
	for {

		// get gas prices
		gasPrices, err := utils.GetGasPrices(g.Network)
		if err != nil {
			fmt.Printf("Error getting rates: %s\n", err)
			time.Sleep(changeFrequency)
			continue
		}

		nickname = fmt.Sprintf("Standard: %dgwei", gasPrices.Standard)

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
		} else {

			err = dg.UpdateListeningStatus(nickname)
			if err != nil {
				fmt.Printf("Unable to set activity: %s\n", err)
			} else {
				fmt.Printf("Set activity: %s\n", nickname)
			}
		}

		time.Sleep(changeFrequency)

		nickname = fmt.Sprintf("Fast: %dgwei", gasPrices.Fast)

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
		} else {

			err = dg.UpdateListeningStatus(nickname)
			if err != nil {
				fmt.Printf("Unable to set activity: %s\n", err)
			} else {
				fmt.Printf("Set activity: %s\n", nickname)
			}
		}

		time.Sleep(changeFrequency)

		nickname = fmt.Sprintf("Instant: %dgwei", gasPrices.Instant)

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
		} else {

			err = dg.UpdateListeningStatus(nickname)
			if err != nil {
				fmt.Printf("Unable to set activity: %s\n", err)
			} else {
				fmt.Printf("Set activity: %s\n", nickname)
			}
		}

		time.Sleep(changeFrequency)
	}

}
