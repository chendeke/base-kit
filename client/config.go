package client

import (
	"github.com/chendeke/base-kit/sd"
	"github.com/chendeke/config"
	"github.com/chendeke/logs"
	"strings"
)

type Config struct {
	sd.Config
}

func NewConfig(path ...string) *Config {
	cfg := &Config{}
	err := config.Get(path...).Scan(cfg)
	if err != nil {
		logs.Errorw("failed to get the server config from "+strings.Join(path, "."), "error", err.Error())
		return nil
	}
	return cfg
}
