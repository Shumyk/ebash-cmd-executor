package config

var config Config

func Load() {
	config.Load()
}

func App() AppConfig {
	return config.AppConfig
}

func Vms() VirtualMachineConfig {
	return config.VirtualMachineConfig
}

func Vagrant() VagrantConfig {
	return config.VirtualMachineConfig.VagrantConfig
}

func Persistance() PersistanceConfig {
	return config.PersistanceConfig
}
