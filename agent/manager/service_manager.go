package manager

import "github.com/scouter-argos/agent/config"

// ServiceStart is a function which starts monitoring service.
func ServiceStart() {
	configManager := config.GetInstance()
	configManager.ReadConfig()
	//ip := configManager.Conf.CollectorIP
	configManager.Start()
}
