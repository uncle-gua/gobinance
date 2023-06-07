package futures

import "github.com/uncle-gua/wsc"

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
		ws.OnTextMessageReceived(handler)
		ws.Connect()
		for range done {
			ws.Close()
			return
		}
	}()
	return
}
