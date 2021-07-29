package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/rssnyder/discord-stock-ticker/utils"
)

type Matic struct {
	Contract  string        `json:"contract"`
	Name      string        `json:"name"`
	Nickname  bool          `json:"nickname"`
	Frequency time.Duration `json:"frequency"`
	Currency  string        `json:"currency"`
	Decimals  int           `json:"decimals"`
	token     string        `json:"-"`
	close     chan int      `json:"-"`
}

// NewMatic saves information about the stock and starts up a watcher on it
func NewMatic(contract string, token string, name string, nickname bool, frequency int, currency string, decimals int) *Matic {
	m := &Matic{
		Contract:  contract,
		Name:      name,
		Nickname:  nickname,
		Frequency: time.Duration(frequency) * time.Second,
		Currency:  currency,
		token:     token,
		close:     make(chan int, 1),
	}

	// spin off go routine to watch the price
	go m.watchMaticPrice()
	return m
}

// Shutdown sends a signal to shut off the goroutine
func (m *Matic) Shutdown() {
	m.close <- 1
}

func (m *Matic) watchMaticPrice() {

	// create a new discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + m.token)
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
		m.Nickname = false
	}

	logger.Infof("Watching token price for %s", m.Name)
	ticker := time.NewTicker(m.Frequency)

	// continuously watch
	for {
		select {
		case <-m.close:
			logger.Infof("Shutting down price watching for %s", m.Name)
			return
		case <-ticker.C:
			logger.Infof("Fetching stock price for %s", m.Name)

			// save the price struct & do something with it
			priceData, err := utils.GetMaticPrice(m.Contract, m.Currency)
			if err != nil {
				logger.Errorf("Unable to fetch stock price for %s", m.Name)
			}

			var fmtPriceRaw float64

			if fmtPriceRaw, err = strconv.ParseFloat(priceData.Totokenamount, 64); err != nil {
				logger.Errorf("Error with price format for %s", m.Name)
			}

			fmtPrice := fmtPriceRaw / 10000000

			if m.Nickname {
				// update nickname instead of activity
				var nickname string
				var activity string

				// format nickname & activity
				// Check for custom decimal places
				switch m.Decimals {
				case 1:
					nickname = fmt.Sprintf("%s - $%.1f", m.Name, fmtPrice)
				case 2:
					nickname = fmt.Sprintf("%s - $%.2f", m.Name, fmtPrice)
				case 3:
					nickname = fmt.Sprintf("%s - $%.3f", m.Name, fmtPrice)
				case 4:
					nickname = fmt.Sprintf("%s - $%.4f", m.Name, fmtPrice)
				case 5:
					nickname = fmt.Sprintf("%s - $%.5f", m.Name, fmtPrice)
				case 6:
					nickname = fmt.Sprintf("%s - $%.6f", m.Name, fmtPrice)
				case 7:
					nickname = fmt.Sprintf("%s - $%.7f", m.Name, fmtPrice)
				case 8:
					nickname = fmt.Sprintf("%s - $%.8f", m.Name, fmtPrice)
				case 9:
					nickname = fmt.Sprintf("%s - $%.9f", m.Name, fmtPrice)
				case 10:
					nickname = fmt.Sprintf("%s - $%.10f", m.Name, fmtPrice)
				case 11:
					nickname = fmt.Sprintf("%s - $%.11f", m.Name, fmtPrice)
				default:
					nickname = fmt.Sprintf("%s - $%.4f", m.Name, fmtPrice)
				}

				activity = "Using USDC on 1inch"

				// Update nickname in guilds
				for _, g := range guilds {
					err = dg.GuildMemberNickname(g.ID, "@me", nickname)
					if err != nil {
						fmt.Println("Error updating nickname: ", err)
						continue
					}
					logger.Infof("Set nickname in %s: %s", g.Name, nickname)
				}

				err = dg.UpdateGameStatus(0, activity)
				if err != nil {
					logger.Error("Unable to set activity: ", err)
				} else {
					logger.Infof("Set activity: %s", activity)
				}

			} else {
				activity := fmt.Sprintf("%s - $%.2f", m.Name, fmtPrice)

				err = dg.UpdateGameStatus(0, activity)
				if err != nil {
					logger.Error("Unable to set activity: ", err)
				} else {
					logger.Infof("Set activity: %s", activity)
				}
			}
		}
	}
}
