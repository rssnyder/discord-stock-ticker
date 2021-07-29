package main

import (
	"fmt"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/rssnyder/discord-stock-ticker/utils"
)

// Holders represents the json for holders
type Holders struct {
	Network   string        `json:"network"`
	Address   string        `json:"address"`
	Activity  string        `json:"activity"`
	Nickname  bool          `json:"set_nickname"`
	Frequency time.Duration `json:"frequency"`
	token     string        `json:"-"`
	close     chan int      `json:"-"`
}

// NewHolders saves information about the stock and starts up a watcher on it
func NewHolders(network string, address string, activity string, token string, nickname bool, frequency int) *Holders {
	h := &Holders{
		Network:   network,
		Address:   address,
		Activity:  activity,
		Nickname:  nickname,
		Frequency: time.Duration(frequency) * time.Second,
		token:     token,
		close:     make(chan int, 1),
	}

	// spin off go routine to watch the price
	go h.watchHolders()
	return h
}

// Shutdown sends a signal to shut off the goroutine
func (h *Holders) Shutdown() {
	h.close <- 1
}

func (h *Holders) watchHolders() {

	// create a new discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + h.token)
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

	// set activity as desc
	if h.Nickname {
		err = dg.UpdateGameStatus(0, h.Activity)
		if err != nil {
			fmt.Printf("Unable to set activity: %s\n", err)
		} else {
			fmt.Println("Set activity")
		}
	}

	// get guides for bot
	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		fmt.Println("Error getting guilds: ", err)
		h.Nickname = false
	}

	ticker := time.NewTicker(h.Frequency)
	var nickname string

	for {

		select {
		case <-h.close:
			logger.Infof("Shutting down price watching for %s", h.Activity)
			return
		case <-ticker.C:

			nickname = utils.GetHolders(h.Network, h.Address)

			if h.Nickname {

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
