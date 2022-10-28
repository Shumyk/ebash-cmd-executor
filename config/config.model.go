package config

import (
	"ebash/cmd-executor/util"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	AppConfig            `yaml:"app"`
	VirtualMachineConfig `yaml:"vms"`
	PersistenceConfig    `yaml:"persistence"`
}

type AppConfig struct {
	Port string `yaml:"port"`
}

type VirtualMachineConfig struct {
	RunOn           string `yaml:"runOn"`
	SessionPoolSize int    `yaml:"sessionPoolSize"`
	VagrantConfig   `yaml:"vagrant"`
}

type VagrantConfig struct {
	Vagrantfiles []string `yaml:"vagrantfiles"`
	Verbose      bool     `yaml:"verbose"`
	Halt         bool     `yaml:"halt"`
}

type PersistenceConfig struct {
	Enabled bool `yaml:"enabled"`
}

func load() *Config {
	configFile := util.Cautiosly(os.ReadFile("config/application.yaml"))("read config file")
	config := &Config{}
	util.Panically(yaml.Unmarshal(configFile, config), "unmarshal config file")
	log.Println("successfully loaded configurations")
	return config
}
