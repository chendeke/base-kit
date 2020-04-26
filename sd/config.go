package sd

import (
	"github.com/chendeke/base-kit/retry"
	"github.com/chendeke/base-kit/sd/consul"
	"github.com/chendeke/base-kit/sd/direct"
	"github.com/chendeke/base-kit/sd/etcdv3"
	"github.com/chendeke/config"
)

type Config struct {
	Mode   string                    `json:"mode" yaml:"mode" db:"mode"`
	Url    string                    `json:"url" yaml:"url"`
	Retry  *retry.Config             `json:"retry" yaml:"retry" db:"retry"`
	EtcdV3 *etcdv3.Config            `json:"etcd" yaml:"etcd"`
	Consul *consul.Config            `json:"consul" yaml:"consul"`
	Direct map[string]*direct.Config `json:"direct" yaml:"direct" db:"direct"`
}

func NewConfig(path string) *Config {
	cfg := &Config{}
	if err := config.ScanKey(path, cfg); err != nil {
		return nil
	}
	return cfg
}
