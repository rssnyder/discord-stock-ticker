package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// TokenRequest represents the json coming in from the request
type TokenRequest struct {
	Network   string `json:"network"`
	Contract  string `json:"contract"`
	Token     string `json:"discord_bot_token"`
	Name      string `json:"name"`
	Nickname  bool   `json:"set_nickname"`
	Frequency int    `json:"frequency" default:"60"`
	Currency  string `json:"currency" default:"0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"`
	Color     bool   `json:"set_color"`
	Decorator string `json:"decorator" default:"-"`
	Activity  string `json:"activity"`
	Decimals  int    `json:"decimals"`
}

// AddToken adds a new Token or crypto to the list of what to watch
func (m *Manager) AddToken(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to add a ticker")

	// read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error reading body: %v", err)
		return
	}

	// unmarshal into struct
	var tokenReq TokenRequest
	if err := json.Unmarshal(body, &tokenReq); err != nil {
		logger.Errorf("Error unmarshalling: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error unmarshalling: %v", err)
		return
	}

	// ensure token is set
	if tokenReq.Token == "" {
		logger.Error("Discord token required")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: token required")
		return
	}

	// ensure network is set, default to eth
	if tokenReq.Network == "" {
		tokenReq.Network = "ethereum"
	}

	// ensure currency is set, default to USDC
	if tokenReq.Currency == "" {
		// Get contract address for usdc
		switch tokenReq.Network {
		case "ethereum":
			tokenReq.Currency = "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
		case "binance-smart-chain":
			tokenReq.Currency = "0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d"
		case "polygon":
			tokenReq.Currency = "0x2791Bca1f2de4661ED88A30C99A7a9449Aa84174"
		default:
			tokenReq.Currency = "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"
		}
	}

	// ensure freq is set
	if tokenReq.Frequency == 0 {
		tokenReq.Frequency = 60
	}

	// ensure name is set
	if tokenReq.Name == "" {
		logger.Error("Name required for token")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: Name required")
		return
	}

	// check if already existing
	if _, ok := m.WatchingToken[strings.ToUpper(tokenReq.Contract)]; ok {
		logger.Error("Error: ticker already exists")
		w.WriteHeader(http.StatusConflict)
		return
	}

	token := NewToken(tokenReq.Network, tokenReq.Contract, tokenReq.Token, tokenReq.Name, tokenReq.Nickname, tokenReq.Frequency, tokenReq.Currency, tokenReq.Decimals, tokenReq.Activity, tokenReq.Color, tokenReq.Decorator)
	m.addToken(token)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(token)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (m *Manager) addToken(token *Token) {
	tokenCount.Inc()
	m.WatchingToken[fmt.Sprintf("%s-%s", token.Network, token.Contract)] = token
}

// DeleteToken addds a new token or crypto to the list of what to watch
func (m *Manager) DeleteToken(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to delete a ticker")

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := m.WatchingToken[id]; !ok {
		logger.Error("Error: no ticker found")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Error: ticker not found")
		return
	}
	// send shutdown sign
	m.WatchingToken[id].Shutdown()
	tokenCount.Dec()

	// remove from cache
	delete(m.WatchingToken, id)

	logger.Infof("Deleted ticker %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// GetToken returns a list of what the manager is watching
func (m *Manager) GetToken(w http.ResponseWriter, r *http.Request) {
	m.RLock()
	defer m.RUnlock()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m.WatchingToken); err != nil {
		logger.Errorf("Error serving request: %v", err)
		fmt.Fprintf(w, "Error: %v", err)
	}
}
