package manager

import (
	"github.com/scouter-argos/agent/config"
	"github.com/scouter-argos/agent/instance"
)

// DBInstances is a slice that has a db instance
var DBInstances []*instance.Instance

// ServiceStart is a function which starts monitoring service.
func ServiceStart() {
	configManager := config.GetInstance()
	configManager.ReadConfig()
	//ip := configManager.Conf.CollectorIP
	configManager.Start()

	instCount := len(configManager.Conf.Instances)

	for i := 0; i < instCount; i++ {
		inst := configManager.Conf.Instances[i]
		inst.StartMonitor()
		DBInstances = append(DBInstances, &inst)

	}

}
