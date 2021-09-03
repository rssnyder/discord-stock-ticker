package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// BoardRequest represents the json coming in from the request
type BoardRequest struct {
	Items     []string `json:"items"`
	Token     string   `json:"discord_bot_token"`
	Name      string   `json:"name"`
	Header    string   `json:"header"`
	Nickname  bool     `json:"set_nickname"`
	Crypto    bool     `json:"crypto"`
	Color     bool     `json:"set_color"`
	Frequency int      `json:"frequency"`
}

// AddBoard adds a new board to the list of what to watch
func (m *Manager) AddBoard(w http.ResponseWriter, r *http.Request) {
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
	var boardReq BoardRequest
	if err := json.Unmarshal(body, &boardReq); err != nil {
		logger.Errorf("Error unmarshalling: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error unmarshalling: %v", err)
		return
	}

	// ensure token is set
	if boardReq.Token == "" {
		logger.Error("Discord token required")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: token required")
		return
	}

	// ensure name is set
	if boardReq.Name == "" {
		logger.Error("Board Name required")
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error: name required")
		return
	}

	// add stock or crypto ticker
	if boardReq.Crypto {

		// check if already existing
		if _, ok := m.WatchingBoard[strings.ToUpper(boardReq.Name)]; ok {
			logger.Error("Error: board already exists")
			w.WriteHeader(http.StatusConflict)
			return
		}

		crypto := NewCryptoBoard(boardReq.Items, boardReq.Token, boardReq.Name, boardReq.Header, boardReq.Nickname, boardReq.Color, boardReq.Frequency, m.Cache, m.Context)
		m.addBoard(true, crypto)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(crypto)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}
		return
	}

	// check if already existing
	if _, ok := m.WatchingBoard[strings.ToUpper(boardReq.Name)]; ok {
		logger.Error("Error: board already exists")
		w.WriteHeader(http.StatusConflict)
		return
	}

	stock := NewStockBoard(boardReq.Items, boardReq.Token, boardReq.Name, boardReq.Header, boardReq.Nickname, boardReq.Color, boardReq.Frequency)
	m.addBoard(false, stock)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(stock)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (m *Manager) addBoard(crypto bool, board *Board) {
	boardCount.Inc()
	id := board.Name
	m.WatchingBoard[id] = board

	var noDB *sql.DB
	if m.DB == noDB {
		return
	}

	// query
	stmt, err := m.DB.Prepare("SELECT id FROM tickers WHERE tickerType = 'board' AND name = ?")
	if err != nil {
		logger.Warningf("Unable to query board in db %s: %s", id, err)
		return
	}

	rows, err := stmt.Query(board.Name)
	if err != nil {
		logger.Warningf("Unable to query board in db %s: %s", id, err)
		return
	}

	var existingId int

	for rows.Next() {
		err = rows.Scan(&existingId)
		if err != nil {
			logger.Warningf("Unable to query board in db %s: %s", id, err)
			return
		}
	}
	rows.Close()

	if existingId != 0 {

		// update entry in db
		stmt, err := m.DB.Prepare("update tickers set token = ?, name = ?, nickname = ?, color = ?, crypto = ?, header = ?, items = ?, frequency = ? WHERE id = ?")
		if err != nil {
			logger.Warningf("Unable to update board in db %s: %s", id, err)
			return
		}

		res, err := stmt.Exec(board.token, board.Name, board.Nickname, board.Color, crypto, board.Header, strings.Join(board.Items, ";"), board.Frequency, existingId)
		if err != nil {
			logger.Warningf("Unable to update board in db %s: %s", id, err)
			return
		}

		_, err = res.LastInsertId()
		if err != nil {
			logger.Warningf("Unable to update board in db %s: %s", id, err)
			return
		}

		logger.Infof("Updated board in db %s", id)
	} else {

		// store new entry in db
		stmt, err := m.DB.Prepare("INSERT INTO tickers(tickerType, token, name, nickname, color, crypto, header, items, frequency) values(?,?,?,?,?,?,?,?,?)")
		if err != nil {
			logger.Warningf("Unable to store board in db %s: %s", id, err)
			return
		}

		res, err := stmt.Exec("board", board.token, board.Name, board.Nickname, board.Color, crypto, board.Header, strings.Join(board.Items, ";"), board.Frequency)
		if err != nil {
			logger.Warningf("Unable to store board in db %s: %s", id, err)
			return
		}

		_, err = res.LastInsertId()
		if err != nil {
			logger.Warningf("Unable to store board in db %s: %s", id, err)
			return
		}
	}
}

// DeleteBoard removes a board
func (m *Manager) DeleteBoard(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to delete a board")

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := m.WatchingBoard[id]; !ok {
		logger.Error("Error: no ticker found")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Error: ticker not found")
		return
	}
	// send shutdown sign
	m.WatchingBoard[id].Shutdown()
	boardCount.Dec()

	// remove from cache
	delete(m.WatchingBoard, id)

	logger.Infof("Deleted board %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// GetBoards returns a list of what the manager is watching
func (m *Manager) GetBoards(w http.ResponseWriter, r *http.Request) {
	m.RLock()
	defer m.RUnlock()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m.WatchingBoard); err != nil {
		logger.Errorf("Error serving request: %v", err)
		fmt.Fprintf(w, "Error: %v", err)
	}
}
