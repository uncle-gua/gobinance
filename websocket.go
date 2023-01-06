package binance

import (
	"net/http"
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

var wsServe = func(cfg *WsConfig, handler WsHandler, errHandler ErrHandler) (doneC, stopC chan struct{}, err error) {
	Dialer := websocket.Dialer{
		Proxy:             http.ProxyFromEnvironment,
		HandshakeTimeout:  45 * time.Second,
		EnableCompression: false,
	}

	c, _, err := Dialer.Dial(cfg.Endpoint, nil)
	if err != nil {
		return nil, nil, err
	}
	c.SetReadLimit(655350)
	doneC = make(chan struct{})
	stopC = make(chan struct{})
	go func() {
		// This function will exit either on error from
		// websocket.Conn.ReadMessage or when the stopC channel is
		// closed by the client.
		defer close(doneC)
		// Wait for the stopC channel to be closed.  We do that in a
		// separate goroutine because ReadMessage is a blocking
		// operation.
		stop := false
		go func() {
			select {
			case <-stopC:
				stop = true
			case <-doneC:
			}
			c.Close()
		}()
		if WebsocketKeepalive {
			ticker := time.NewTicker(WebsocketTimeout)

			go func() {
				defer ticker.Stop()
				for {
					if stop {
						return
					}
					deadline := time.Now().Add(10 * time.Second)
					err := c.WriteControl(websocket.PingMessage, []byte{}, deadline)
					if err != nil {
						errHandler(err)
					}
					<-ticker.C
				}
			}()
		}
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err) {
					if stop {
						return
					}

					errHandler(err)
					c, _, err = Dialer.Dial(cfg.Endpoint, nil)
					if err != nil {
						errHandler(err)
					}
					continue
				}

				errHandler(err)
				continue
			}
			handler(message)
		}
	}()
	return
}
