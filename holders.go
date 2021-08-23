package main

import (
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
		logger.Errorf("Error creating Discord session: %s\n", err)
		return
	}

	// show as online
	err = dg.Open()
	if err != nil {
		logger.Errorf("error opening discord connection: %s\n", err)
		return
	}

	// set activity as desc
	if h.Nickname {
		err = dg.UpdateGameStatus(0, h.Activity)
		if err != nil {
			logger.Errorf("Unable to set activity: %s\n", err)
		} else {
			logger.Infof("Set activity")
		}
	}

	// get guides for bot
	guilds, err := dg.UserGuilds(100, "", "")
	if err != nil {
		logger.Errorf("Error getting guilds: %s\n", err)
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
						logger.Errorf("Error updating nickname: %s\n", err)
						continue
					} else {
						logger.Infof("Set nickname in %s: %s\n", g.Name, nickname)
					}
				}
			} else {

				err = dg.UpdateGameStatus(0, nickname)
				if err != nil {
					logger.Errorf("Unable to set activity: %s\n", err)
				} else {
					logger.Infof("Set activity: %s\n", nickname)
				}
			}
		}
	}
}
