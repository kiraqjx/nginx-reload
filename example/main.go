package main

import (
	"os"

	"github.com/kiraqjx/nginx-reload/pkg/dispenser"
	"github.com/kiraqjx/nginx-reload/pkg/entity"
	"github.com/kiraqjx/nginx-reload/pkg/vo"
	"gopkg.in/yaml.v3"
)

func main() {
	// load config file
	file, err := os.ReadFile("../config/config.yaml")
	if err != nil {
		panic(err)
	}
	config := vo.Config{}
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}
	dispenser, err := dispenser.NewDispenser(&MemoryStore{}, config.NginxTemplate, config.SshConfigs, true)
	if err != nil {
		panic(err)
	}
	err = dispenser.Do()
	if err != nil {
		panic(err)
	}
}

type MemoryStore struct {
}

func (ms *MemoryStore) AllConfig() []entity.ProxyConfig {
	return []entity.ProxyConfig{
		{
			Id:         "mysql1",
			Datasource: "127.0.0.1:3306",
			Port:       "3306",
		},
		{
			Id:         "mysql2",
			Datasource: "127.0.0.1:3306",
			Port:       "3306",
		},
	}
}
