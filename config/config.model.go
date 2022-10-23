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
}

type AppConfig struct {
	Port string `yaml:"port"`
}

type VirtualMachineConfig struct {
	RunOn         string `yaml:"runOn"`
	VagrantConfig `yaml:"vagrant"`
}

type VagrantConfig struct {
	Vagrantfiles []string `yaml:"vagrantfiles"`
	Verbose      bool     `yaml:"verbose"`
	Halt         bool     `yaml:"halt"`
}

func (c *Config) Load() {
	configFile := util.Cautiosly(os.ReadFile("config/application.yaml"))("read config file")
	util.Panically(yaml.Unmarshal(configFile, &c), "unmarshal config file")
	log.Println("successfully loaded configurations")
}
