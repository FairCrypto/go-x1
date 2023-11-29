package xenblocks

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gorilla/websocket"
	"sync"
)

var (
	queueSize = 2
)

type Config struct {
	Endpoint string
}

type websocketConfig struct {
	c               *websocket.Conn // Global WebSocket connection
	connectionMutex sync.Mutex      // Mutex to protect the WebSocket connection
}

type workerConfig struct {
	jobChannel   chan WebSocketJob
	errorChannel chan error
}

type Xenblocks struct {
	Config
	websocketConfig
	workerConfig
}

func DefaultConfig() Config {
	return Config{
		Endpoint: "ws://xenblocks.io:6668",
	}
}

type WebSocketJob struct {
	PeerID   string
	BlockID  string
	Hash     string
	TimeDiff string
}

func (x *Xenblocks) EstablishXenBlocksConnection() error {
	x.websocketConfig.connectionMutex.Lock()         // Lock the mutex
	defer x.websocketConfig.connectionMutex.Unlock() // Unlock the mutex when the function exits

	// Check again if the connection is already established
	if x.websocketConfig.c != nil {
		return nil
	}

	log.Info("Establishing XenBlocks connection...", "endpoint", x.Endpoint)

	var err error
	x.websocketConfig.c, _, err = websocket.DefaultDialer.Dial(x.Endpoint, nil)
	if err != nil {
		log.Error("dial error:", "err", err)
		return err
	}
	return nil
}

func (x *Xenblocks) SendDataOverWebSocket(peerID, blockID, hash, timeDiff string) {
	x.jobChannel <- WebSocketJob{PeerID: peerID, BlockID: blockID, Hash: hash, TimeDiff: timeDiff}
}

func (x *Xenblocks) sendDataOverWebSocket(peerID string, blockID string, hash string, timeDiff string) {
	// Check if the connection is established, if not, try to establish it
	if x.c == nil {
		err := x.EstablishXenBlocksConnection()
		if err != nil {
			// Handle connection error
			log.Error("WS error:", "err", err)
			return
		}
	}

	// Prepare the data to be sent
	responseData := map[string]interface{}{
		"peer_id":   peerID,
		"block_id":  blockID,
		"hash":      hash,
		"time_diff": timeDiff,
	}

	// Marshal the response data to JSON
	jsonData, err := json.Marshal(responseData)
	if err != nil {
		log.Error("JSON marshal error:", err)
		return
	}

	// Send the JSON response through the WebSocket
	log.Info("Sending data to XenBlocks...", "data", responseData)
	if err := x.websocketConfig.c.WriteMessage(websocket.TextMessage, jsonData); err != nil {
		log.Error("write error:", "err", err)
		err := x.EstablishXenBlocksConnection()
		if err != nil {
			// Handle connection error
			log.Error("WS error:", "err", err)
		}
	}
}

func (x *Xenblocks) worker() {
	for j := range x.jobChannel {
		x.sendDataOverWebSocket(j.PeerID, j.BlockID, j.Hash, j.TimeDiff)
	}
	log.Warn("worker done?")
}

func (x *Xenblocks) Start() error {
	err := x.EstablishXenBlocksConnection()
	if err != nil {
		return err
	}
	x.jobChannel = make(chan WebSocketJob, queueSize)
	go x.worker()
	return nil
}

func (x *Xenblocks) Stop() {
	close(x.jobChannel)
}
