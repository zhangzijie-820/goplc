package goplc

import (
	"fmt"
	"log"
	"time"
)

type Config struct {
	ENIP_PORT            uint16
	Log                  *log.Logger
	ReconnectionInterval time.Duration

	OnConnected    func()
	OnDisconnected func(err error)
}

func (c *Config) Println(v ...interface{}) {
	if c.Log != nil {
		c.Log.Println(v...)
	}
}

func (c *Config) Printf(fmt string, v ...interface{}) {
	c.Log.Printf(fmt, v...)
}

var defaultConfig *Config

func DefaultConfig() *Config {
	_defaultConfig := &Config{}
	_defaultConfig.ENIP_PORT = 0xAF12
	_defaultConfig.Log = nil
	_defaultConfig.ReconnectionInterval = time.Second * 1

	return _defaultConfig
}

func init() {
	defaultConfig = DefaultConfig()
	fmt.Println(defaultConfig)
}
