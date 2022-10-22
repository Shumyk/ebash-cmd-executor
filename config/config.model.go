package config

import (
	"io/ioutil"
	"log"

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
	configFile, err := ioutil.ReadFile("config/application.yaml")
	logPanically(err, "read")

	err = yaml.Unmarshal(configFile, &c)
	logPanically(err, "unmarshal")

	log.Println("successfully loaded configurations")
}

func logPanically(err error, action string) {
	if err != nil {
		log.Panicf("could not %v config file: %v", action, err)
	}
}
