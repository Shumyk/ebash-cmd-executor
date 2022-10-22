package config

var config Config

func Load() {
	config.Load()
}

func App() AppConfig {
	return config.AppConfig
}

func Vagrant() VagrantConfig {
	return config.VirtualMachineConfig.VagrantConfig
}
