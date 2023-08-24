package log

import "log"

type Config struct {
	OnConnected    bool
	OnClose        bool
	OnPingReceived bool
	OnPongReceived bool
	OnKeepalive    bool
	Log            func(fmt string, a ...any)
}

var Default = &Config{
	OnConnected:    true,
	OnClose:        true,
	OnPingReceived: false,
	OnPongReceived: false,
	OnKeepalive:    false,
	Log:            log.Printf,
}

func SetConnected(value bool) *Config {
	Default.OnConnected = value
	return Default
}

func SetClose(value bool) *Config {
	Default.OnClose = value
	return Default
}

func SetPingReceived(value bool) *Config {
	Default.OnPingReceived = value
	return Default
}

func SetPongReceived(value bool) *Config {
	Default.OnPongReceived = value
	return Default
}

func SetKeepalive(value bool) *Config {
	Default.OnKeepalive = value
	return Default
}

func SetLog(value func(fmt string, a ...any)) *Config {
	Default.Log = value
	return Default
}
