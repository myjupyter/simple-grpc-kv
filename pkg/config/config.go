package config

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type ServiceConfig struct {
	ServiceHost string `json:"host"`
	ServicePort string `json:"port"`
}

func (sc ServiceConfig) Host() string {
	return sc.ServiceHost
}

func (sc ServiceConfig) Port() string {
	return sc.ServicePort
}

type StorageConfig struct {
	SavePath string `json:"save_path"`
	SaveTime string `json:"save_time"`
	time     time.Duration
}

func (sc StorageConfig) Path() string {
	return sc.SavePath
}

func (sc StorageConfig) Time() time.Duration {
	return sc.time
}

type Config struct {
	GRPC    ServiceConfig `json:"grpc"`
	HTTP    ServiceConfig `json:"http,omitempty"`
	Storage StorageConfig `json:"storage"`
}

func NewConfigFromFile(configPath string) (*Config, error) {
	raw, err := ioutil.ReadFile(configPath)
	if err != nil {
		return nil, err
	}
	return NewConfig(raw)
}

func NewConfig(raw []byte) (*Config, error) {
	config := new(Config)
	if err := json.Unmarshal(raw, config); err != nil {
		return nil, err
	}
	t, err := time.ParseDuration(config.Storage.SaveTime)
	if err != nil {
		return nil, err
	}
	config.Storage.time = t
	return config, nil
}
