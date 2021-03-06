package http

import (
	"errors"
	"github.com/elastic/beats/v7/libbeat/outputs/codec"
	"net/http"
)

type config struct {
	URL                   string                 `config:"url"`
	Codec                 codec.Config           `config:"codec"`
	OnlyFields            bool                   `config:"only_fields"`
	MaxRetries            int                    `config:"max_retries"`
	Compression           bool                   `config:"compression"`
	KeepAlive             bool                   `config:"keep_alive"`
	MaxIdleConns          int                    `config:"max_idle_conns"`
	IdleConnTimeout       int                    `config:"idle_conn_timeout"`
	ResponseHeaderTimeout int                    `config:"response_header_timeout"`
	Username              string                 `config:"username"`
	Password              string                 `config:"password"`
	AddFields             map[string]interface{} `config:"add_fields"`
}

var (
	defaultConfig = config{
		URL:                   "http://127.0.0.1:8090/message",
		OnlyFields:            false,
		MaxRetries:            -1,
		Compression:           false,
		KeepAlive:             true,
		MaxIdleConns:          1,
		IdleConnTimeout:       0,
		ResponseHeaderTimeout: 100,
		Username:              "",
		Password:              "",
		AddFields:             make(map[string]interface{}, 0),
	}
)

func (c *config) Validate() error {
	_, err := http.NewRequest("POST", c.URL, nil)
	if err != nil {
		return err
	}
	if c.MaxIdleConns < 1 {
		return errors.New("max_idle_conns can't be <1")
	}
	if c.IdleConnTimeout < 0 {
		return errors.New("idle_conn_timeout can't be <0")
	}
	if c.ResponseHeaderTimeout < 1 {
		return errors.New("response_header_timeout can't be <1")
	}
	if c.Username != "" && c.Password == "" {
		return errors.New("password can't be empty")
	}
	return nil
}
