package xenblocks

import (
	"encoding/json"
	"github.com/ethereum/go-ethereum/log"
	"github.com/gorilla/websocket"
	"github.com/recws-org/recws"
)

var (
	queueSize = 2
)

type Config struct {
	Endpoint string
	Enabled  bool
}

type workerConfig struct {
	jobChannel   chan WebSocketJob
	errorChannel chan error
}

type Xenblocks struct {
	Config
	ws recws.RecConn
	workerConfig
}

func DefaultConfig() Config {
	return Config{
		Endpoint: "",
	}
}

type WebSocketJob struct {
	PeerID   string
	BlockID  string
	Hash     string
	TimeDiff string
}

func (x *Xenblocks) Send(peerID, blockID, hash, timeDiff string) {
	x.jobChannel <- WebSocketJob{PeerID: peerID, BlockID: blockID, Hash: hash, TimeDiff: timeDiff}
}

func (x *Xenblocks) sendDataOverWebSocket(peerID string, blockID string, hash string, timeDiff string) {
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
	log.Info("Sending data to XenBlocks", "data", responseData)
	if err := x.ws.WriteMessage(websocket.TextMessage, jsonData); err != nil {
		log.Error("Failed to send peer info to xenblocks", "err", err)
	}
}

func (x *Xenblocks) worker() {
	for j := range x.jobChannel {
		x.sendDataOverWebSocket(j.PeerID, j.BlockID, j.Hash, j.TimeDiff)
	}
}

func (x *Xenblocks) establishConnection() {
	log.Info("Establishing connection to XenBlocks")
	x.ws = recws.RecConn{
		NonVerbose: true,
	}
	x.ws.Dial(x.Endpoint, nil)
}

func (x *Xenblocks) Start() {
	x.Enabled = true
	x.establishConnection()
	x.jobChannel = make(chan WebSocketJob, queueSize)
	go x.worker()
}

func (x *Xenblocks) Stop() {
	close(x.jobChannel)
}
