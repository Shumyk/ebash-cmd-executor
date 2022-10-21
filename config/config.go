package config

var config Config

func Load() {
	config.Load()
}

func GetApp() App {
	return config.App
}

func GetVagrant() Vagrant {
	return config.VirtualMachine.Vagrant
}
