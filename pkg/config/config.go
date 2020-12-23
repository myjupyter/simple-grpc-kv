package config

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

type AppConfig struct {
	H string `json:"host"`
	P string `json:"port"`
}

func (ac AppConfig) Host() string {
	return ac.H
}

func (ac AppConfig) Port() string {
	return ac.P
}

type StorageConfig struct {
	SP string `json:"savepath"`
	ST string `json:"savetime"`
}

func (sc StorageConfig) SavePath() string {
	return sc.SP
}

func (sc StorageConfig) SaveTime() string {
	return sc.ST
}

type Config struct {
	AppConfig     `json:"grpc"`
	StorageConfig `json:"storage"`
}

func NewConfigFromFile(configPath string) (Config, error) {
	raw, err := ioutil.ReadFile(configPath)
	if err != nil {
		return Config{}, err
	}
	return NewConfig(raw)
}

func NewConfig(raw []byte) (Config, error) {
	config := Config{}
	if err := json.Unmarshal(raw, &config); err != nil {
		return Config{}, err
	}
	_, err := time.ParseDuration(config.SaveTime())
	if err != nil {
		return Config{}, err
	}
	return config, nil
}
