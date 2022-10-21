package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App            `yaml:"app"`
	VirtualMachine `yaml:"vms"`
}

type App struct {
	Port string `yaml:"port"`
}

type VirtualMachine struct {
	Vagrant `yaml:"vagrant"`
}

type Vagrant struct {
	Vagrantfiles []string `yaml:"vagrantfiles"`
	Verbose      bool     `yaml:"verbose"`
}

func (c *Config) Load() {
	configFile, err := ioutil.ReadFile("config/application.yaml")
	logPanically(err, "read")

	err = yaml.Unmarshal(configFile, &c)
	logPanically(err, "decode")

	log.Println("successfully loaded configurations")
}

func logPanically(err error, action string) {
	if err != nil {
		log.Panicf("could not %v config file: %v", action, err)
	}
}
