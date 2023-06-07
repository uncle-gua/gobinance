package futures

import (
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WsHandler handle raw websocket message
type WsHandler func(message []byte)

// ErrHandler handles errors
type ErrHandler func(err error)

// WsConfig webservice configuration
type WsConfig struct {
	Endpoint string
}

func newWsConfig(endpoint string) *WsConfig {
	return &WsConfig{
		Endpoint: endpoint,
	}
}

var wsServe = func(cfg *WsConfig, handler WsHandler, errHandler ErrHandler) (done chan struct{}, err error) {
	done = make(chan struct{})

	go func() {
		ws := &Wsc{
			Config: &Config{
				WriteWait:         10 * time.Second,
				MaxMessageSize:    512,
				MinRecTime:        2 * time.Second,
				MaxRecTime:        60 * time.Second,
				RecFactor:         1.5,
				MessageBufferSize: 256,
			},
			WebSocket: &WebSocket{
				Url:           cfg.Endpoint,
				Dialer:        websocket.DefaultDialer,
				RequestHeader: http.Header{},
				isConnected:   false,
				connMu:        &sync.RWMutex{},
				sendMu:        &sync.Mutex{},
			},
		}
		ws.OnTextMessageReceived(handler)
		ws.Connect()
		for range done {
			ws.Close()
			return
		}
	}()
	return
}
