package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// ImportGas pulls in bots from the provided db
func (m *Manager) ImportGas() {

	// query
	rows, err := m.DB.Query("SELECT clientID, token, nickname, network, frequency FROM gases")
	if err != nil {
		logger.Warningf("Unable to query tokens in db: %s", err)
		return
	}

	// load existing bots from db
	for rows.Next() {
		var importedGas Gas

		err = rows.Scan(&importedGas.ClientID, &importedGas.Token, &importedGas.Nickname, &importedGas.Network, &importedGas.Frequency)
		if err != nil {
			logger.Errorf("Unable to load token from db: %s", err)
			continue
		}

		go importedGas.watchGasPrice()
		m.WatchGas(&importedGas)
		logger.Infof("Loaded gas from db: %s", importedGas.label())
	}
	rows.Close()

	// check all entries have id
	for _, gas := range m.WatchingGas {
		if gas.ClientID == "" {
			id, err := getIDToken(gas.Token)
			if err != nil {
				logger.Errorf("Unable to get id from token: %s", err)
				continue
			}

			stmt, err := m.DB.Prepare("UPDATE gases SET clientId = ? WHERE token = ?")
			if err != nil {
				logger.Errorf("Unable to prepare id update: %s", err)
				continue
			}

			res, err := stmt.Exec(id, gas.Token)
			if err != nil {
				logger.Errorf("Unable to update db: %s", err)
				continue
			}

			_, err = res.LastInsertId()
			if err != nil {
				logger.Errorf("Unable to confirm db update: %s", err)
				continue
			} else {
				logger.Infof("Updated id in db for %s", gas.label())
				gas.ClientID = id
			}
		}
	}
}

// AddTicker adds a new Ticker or crypto to the list of what to watch
func (m *Manager) AddGas(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to add a gas")

	// read body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Errorf("%s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// unmarshal into struct
	var gasReq Gas
	if err := json.Unmarshal(body, &gasReq); err != nil {
		logger.Errorf("Unmarshalling: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// ensure token is set
	if gasReq.Token == "" {
		logger.Error("Discord token required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// make sure token is valid
	if gasReq.ClientID == "" {
		id, err := getIDToken(gasReq.Token)
		if err != nil {
			logger.Errorf("Unable to authenticate with discord token: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		gasReq.ClientID = id
	}

	// ensure frequency is set
	if gasReq.Frequency <= 0 {
		gasReq.Frequency = 60
	}

	// ensure network is set
	if gasReq.Network == "" {
		logger.Error("Network required")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// check if already existing
	if _, ok := m.WatchingGas[gasReq.label()]; ok {
		logger.Error("Network already exists")
		w.WriteHeader(http.StatusConflict)
		return
	}

	go gasReq.watchGasPrice()
	m.WatchGas(&gasReq)

	if *db != "" {
		m.StoreGas(&gasReq)
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(gasReq)
	if err != nil {
		logger.Errorf("Unable to encode gas: %s", err)
		w.WriteHeader(http.StatusBadRequest)
	}
	logger.Infof("Added gas: %s\n", gasReq.Network)
}

func (m *Manager) WatchGas(gas *Gas) {
	gasCount.Inc()
	id := gas.label()
	m.WatchingGas[id] = gas
}

// StoreGas puts a gas into the db
func (m *Manager) StoreGas(gas *Gas) {

	// store new entry in db
	stmt, err := m.DB.Prepare("INSERT INTO gases(clientId, token, nickname, network, frequency) values(?,?,?,?,?)")
	if err != nil {
		logger.Warningf("Unable to store gas in db %s: %s", gas.label(), err)
		return
	}

	res, err := stmt.Exec(gas.ClientID, gas.Token, gas.Nickname, gas.Network, gas.Frequency)
	if err != nil {
		logger.Warningf("Unable to store gas in db %s: %s", gas.label(), err)
		return
	}

	_, err = res.LastInsertId()
	if err != nil {
		logger.Warningf("Unable to store gas in db %s: %s", gas.label(), err)
		return
	}
}

// DeleteGas addds a new gas or crypto to the list of what to watch
func (m *Manager) DeleteGas(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to delete a gas")

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := m.WatchingGas[id]; !ok {
		logger.Errorf("No gas found: %s", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// send shutdown sign
	m.WatchingGas[id].Close <- 1
	gasCount.Dec()

	var noDB *sql.DB
	if m.DB != noDB {
		// remove from db
		stmt, err := m.DB.Prepare("DELETE FROM gases WHERE network = ?")
		if err != nil {
			logger.Warningf("Unable to query holder in db %s: %s", id, err)
			return
		}

		_, err = stmt.Exec(m.WatchingGas[id].Network)
		if err != nil {
			logger.Warningf("Unable to query holder in db %s: %s", id, err)
			return
		}
	}

	// remove from cache
	delete(m.WatchingGas, id)

	logger.Infof("Deleted gas %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// RestartGas stops and starts a gas
func (m *Manager) RestartGas(w http.ResponseWriter, r *http.Request) {
	m.Lock()
	defer m.Unlock()

	logger.Debugf("Got an API request to restart a gas")

	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := m.WatchingGas[id]; !ok {
		logger.Errorf("No gas found: %s", id)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// send shutdown sign
	m.WatchingGas[id].Close <- 1

	// wait twice the update time
	time.Sleep(time.Duration(m.WatchingGas[id].Frequency) * 2 * time.Second)

	// start the gas again
	go m.WatchingGas[id].watchGasPrice()

	logger.Infof("Restarted gas %s", id)
	w.WriteHeader(http.StatusNoContent)
}

// GetGas returns a list of what the manager is watching
func (m *Manager) GetGas(w http.ResponseWriter, r *http.Request) {
	m.RLock()
	defer m.RUnlock()
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(m.WatchingGas); err != nil {
		logger.Errorf("Serving request: %s", err)
	}
}
