package metrics

import (
	"github.com/chendeke/config"
	"github.com/chendeke/logs"
	"strings"
)

type Config struct {
	Enable     bool   `json:"enable" yaml:"Enable" default:"true"`
	Department string `json:"department" yaml:"department" default:"location"`
	Project    string `json:"project" yaml:"project"`
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
