package config

var config *Config = load()

func App() AppConfig {
	return config.AppConfig
}

func Vms() VirtualMachineConfig {
	return config.VirtualMachineConfig
}

func Vagrant() VagrantConfig {
	return config.VirtualMachineConfig.VagrantConfig
}

func Persistence() PersistenceConfig {
	return config.PersistenceConfig
}
