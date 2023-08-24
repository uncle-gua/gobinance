package delivery

import (
	"github.com/uncle-gua/gobinance/log"
	"github.com/uncle-gua/wsc"
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
		ws := wsc.New(cfg.Endpoint)
		ws.OnConnected(func() {
			if log.Default.OnConnected {
				log.Default.Log("websocket connected")
			}
		})
		ws.OnConnectError(errHandler)
		ws.OnDisconnected(errHandler)
		ws.OnClose(func(code int, text string) {
			if log.Default.OnClose {
				log.Default.Log("websocket closed, code: %d, message: %s", code, text)
			}
		})
		ws.OnSentError(errHandler)
		ws.OnPingReceived(func(appData string) {
			if log.Default.OnPingReceived {
				log.Default.Log("ping received, data: %s", appData)
			}
		})
		ws.OnPongReceived(func(appData string) {
			if log.Default.OnPongReceived {
				log.Default.Log("pong received, data: %s", appData)
			}
		})
		ws.OnTextMessageReceived(handler)
		ws.OnKeepalive(func() {
			if log.Default.OnKeepalive {
				log.Default.Log("keep alive")
			}
		})
		ws.Connect()
		for range done {
			ws.Close()
			return
		}
	}()
	return
}
