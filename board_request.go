package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var (
	itemSplit = ";"
)

// ImportBoard pulls in bots from the provided db
func (m *Manager) ImportBoard() {

	// query
	rows, err := m.DB.Query("SELECT clientID, token, name, nickname, color, crypto, header, items, frequency FROM boards")
	if err != nil {
		logger.Warningf("Unable to query tokens in db: %s", err)
		return
	}

	// load existing bots from db
	for rows.Next() {
		var importedBoard Board
		var itemsBulk string

		err = rows.Scan(&importedBoard.ClientID, &importedBoard.Token, &importedBoard.Name, &importedBoard.Nickname, &importedBoard.Color, &importedBoard.Crypto, &importedBoard.Header, &itemsBulk, &importedBoard.Frequency)
		if err != nil {
			logger.Errorf("Unable to load token from db: %s", err)
			continue
		}

		importedBoard.Items = strings.Split(itemsBulk, itemSplit)
		if importedBoard.Crypto {
			go importedBoard.watchCryptoPrice()
			m.WatchBoard(&importedBoard)
		} else {
			go importedBoard.watchStockPrice()
			m.WatchBoard(&importedBoard)
		}
		logger.Infof("Loaded board from db: %s", importedBoard.label())
	}
	rows.Close()

	// check all entries have id
	for _, board := range m.WatchingBoard {
		if board.ClientID == "" {
			id, err := getIDToken(board.Token)
			if err != nil {
				logger.Errorf("Unable to get id from token: %s", err)
				continue
			}

			stmt, err := m.DB.Prepare("UPDATE boards SET clientId = ? WHERE token = ?")
			if err != nil {
				logger.Errorf("Unable to prepare id update: %s", err)
				continue
			}

			res, err := stmt.Exec(id, board.Token)
			if err != nil {
				logger.Errorf("Unable to update db: %s", err)
				continue
			}

			_, err = res.LastInsertId()
			if err != nil {
				logger.Errorf("Unable to confirm db update: %s", err)
				continue
			} else {
				logger.Infof("Updated id in db for %s", board.label())
				board.ClientID = id
			}
		}
	}
}

// AddBoard adds a new board to the list of what to watch
func (m *Manager) AddBoard(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to add a board")

	// read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("Error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// unmarshal into struct
	var boardReq Board
	if err := json.Unmarshal(body, &boardReq); err != nil {
		logger.Errorf("Error unmarshalling: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// ensure token is set
	if boardReq.Token == "" {
		logger.Error("Discord token required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// make sure token is valid
	if boardReq.ClientID == "" {
		id, err := getIDToken(boardReq.Token)
		if err != nil {
			logger.Errorf("Unable to authenticate with discord token: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		boardReq.ClientID = id
	}

	// ensure frequency is set
	if boardReq.Frequency <= 0 {
		boardReq.Frequency = 60
	}

	// ensure name is set
	if boardReq.Name == "" {
		logger.Error("Board Name required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// add stock or crypto board
	if boardReq.Crypto {

		// check if already existing
		if _, ok := m.WatchingBoard[boardReq.label()]; ok {
			logger.Error("Error: board already exists")
			w.WriteHeader(http.StatusConflict)
			return
		}

		go boardReq.watchCryptoPrice()
		m.WatchBoard(&boardReq)
	} else {

		// check if already existing
		if _, ok := m.WatchingBoard[boardReq.label()]; ok {
			logger.Error("Error: board already exists")
			w.WriteHeader(http.StatusConflict)
			return
		}

		go boardReq.watchStockPrice()
		m.WatchBoard(&boardReq)
	}

	if *db != "" {
		m.StoreBoard(&boardReq)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(boardReq)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	logger.Infof("Added board: %s\n", boardReq.Name)
}

func (m *Manager) WatchBoard(board *Board) {
	boardCount.Inc()
	id := board.label()
	m.WatchingBoard[id] = board
}

// StoreBoard puts a board into the db
func (m *Manager) StoreBoard(board *Board) {

	// store new entry in db
	stmt, err := m.DB.Prepare("INSERT INTO boards(clientId, token, name, nickname, color, crypto, header, items, frequency) values(?,?,?,?,?,?,?,?,?)")
	if err != nil {
		logger.Warningf("Unable to store board in db %s: %s", board.label(), err)
		return
	}

	res, err := stmt.Exec(board.ClientID, board.Token, board.Name, board.Nickname, board.Color, board.Crypto, board.Header, strings.Join(board.Items, itemSplit), board.Frequency)
	if err != nil {
		logger.Warningf("Unable to store board in db %s: %s", board.label(), err)
		return
	}

	_, err = res.LastInsertId()
	if err != nil {
		logger.Warningf("Unable to store board in db %s: %s", board.label(), err)
		return
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
		logger.Error("Error: no board found")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Error: board not found")
		return
	}

	// send shutdown sign
	m.WatchingBoard[id].Close <- 1
	boardCount.Dec()

	var noDB *sql.DB
	if m.DB != noDB {
		// remove from db
		stmt, err := m.DB.Prepare("DELETE FROM boards WHERE name = ?")
		if err != nil {
			logger.Warningf("Unable to query board in db %s: %s", id, err)
			return
		}

		_, err = stmt.Exec(m.WatchingBoard[id].Name)
		if err != nil {
			logger.Warningf("Unable to query board in db %s: %s", id, err)
			return
		}
	}

	// remove from cache
	delete(m.WatchingBoard, id)

	logger.Infof("Deleted board %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// RestartBoard stops and starts a board
func (m *Manager) RestartBoard(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to restart a board")

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := m.WatchingBoard[id]; !ok {
		logger.Error("Error: no board found")
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Error: board not found")
		return
	}

	// send shutdown sign
	m.WatchingBoard[id].Close <- 1

	// wait twice the update time
	time.Sleep(time.Duration(m.WatchingBoard[id].Frequency) * 2 * time.Second)

	// start the board again
	if m.WatchingBoard[id].Crypto {
		go m.WatchingBoard[id].watchCryptoPrice()
	} else {
		go m.WatchingBoard[id].watchStockPrice()
	}

	logger.Infof("Restarted board %s", id)
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
