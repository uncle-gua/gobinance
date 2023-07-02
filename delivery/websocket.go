package delivery

import "github.com/uncle-gua/wsc"

// WsHandler handle raw websocket message
type WsHandler func(message []byte)

// ErrHandler handles errors
type ErrHandler func(err error)

// InfoHandler handles informations
type InfoHandler func(format string, a ...any)

// WsConfig webservice configuration
type WsConfig struct {
	Endpoint string
}

func newWsConfig(endpoint string) *WsConfig {
	return &WsConfig{
		Endpoint: endpoint,
	}
}

var wsServe = func(cfg *WsConfig, handler WsHandler, errHandler ErrHandler, infoHandler InfoHandler) (done chan struct{}, err error) {
	done = make(chan struct{})

	go func() {
		ws := wsc.New(cfg.Endpoint)
		ws.OnConnected(func() {
			infoHandler("websocket connected")
		})
		ws.OnConnectError(errHandler)
		ws.OnDisconnected(errHandler)
		ws.OnClose(func(code int, text string) {
			infoHandler("websocket closed, code: %d, message: %s", code, text)
		})
		ws.OnSentError(errHandler)
		ws.OnPingReceived(func(appData string) {
			infoHandler(appData)
		})
		ws.OnPongReceived(func(appData string) {
			infoHandler(appData)
		})
		ws.OnTextMessageReceived(handler)
		ws.OnKeepalive(func() {
			infoHandler("websocket keepalive")
		})
		ws.Connect()
		for range done {
			ws.Close()
			return
		}
	}()
	return
}
