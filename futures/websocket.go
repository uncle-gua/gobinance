package futures

import (
	"github.com/uncle-gua/wsc"
)

// WsHandler handle raw websocket message
type WsHandler func(message []byte)

// ErrHandler handles errors
type ErrHandler func(err error)

type Logger struct {
	OnConnected    bool
	OnClose        bool
	OnPingReceived bool
	OnPongReceived bool
	OnKeepalive    bool
	Log            func(format string, a ...any)
}

// WsConfig webservice configuration
type WsConfig struct {
	Endpoint string
}

func newWsConfig(endpoint string) *WsConfig {
	return &WsConfig{
		Endpoint: endpoint,
	}
}

var wsServe = func(cfg *WsConfig, handler WsHandler, errHandler ErrHandler, logger *Logger) (done chan struct{}, err error) {
	done = make(chan struct{})

	go func() {
		ws := wsc.New(cfg.Endpoint)
		ws.OnConnected(func() {
			if logger != nil && logger.OnConnected {
				logger.Log("websocket connected")
			}
		})
		ws.OnConnectError(errHandler)
		ws.OnDisconnected(errHandler)
		ws.OnClose(func(code int, text string) {
			if logger != nil && logger.OnClose {
				logger.Log("websocket closed, code: %d, message: %s", code, text)
			}
		})
		ws.OnSentError(errHandler)
		ws.OnPingReceived(func(appData string) {
			if logger != nil && logger.OnPingReceived {
				logger.Log("ping received, data: %s", appData)
			}
		})
		ws.OnPongReceived(func(appData string) {
			if logger != nil && logger.OnPingReceived {
				logger.Log("pong received, data: %s", appData)
			}
		})
		ws.OnTextMessageReceived(handler)
		ws.OnKeepalive(func() {
			if logger != nil && logger.OnPingReceived {
				logger.Log("keep alive")
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
