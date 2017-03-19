package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/scouter-argos/agent/instance"
)

// Configmgr is a object that manages configuration for argos.
type configManager struct {
	stopRunning  chan bool
	confFile     *os.File
	Conf         *Config
	ConfFilePath string
}

var running = make(chan bool)
var confFileModdTime time.Time
var confFileSize int64

// Config is a struct represents configuration.
type Config struct {
	CollectorIP string              `json:"collector.ip"`
	Udpport     string              `json:"collector.udp.port"`
	Tcpport     string              `json:"collector.tcp.port"`
	Instances   []instance.Instance `json:"db.instances"`
}

func (conf *configManager) load() {
	fileInfo, e := conf.confFile.Stat()
	if e != nil {
		//todo: error handling
	}

	if confFileModdTime != fileInfo.ModTime() || confFileSize != fileInfo.Size() {
		file, err := ioutil.ReadFile(conf.ConfFilePath)
		if err != nil {
			//todo :
		}
		json.Unmarshal(file, conf.Conf)
		//fmt.Printf("server ip: %s", config.Configurations.CollectorIP)
		confFileSize = fileInfo.Size()
		confFileModdTime = fileInfo.ModTime()
	}

}

func (conf *configManager) ReadConfig() {
	if conf.initialize() {
		conf.load()
	}
}

func (conf *configManager) initialize() bool {
	f, err := os.Open(conf.ConfFilePath)
	if err != nil {
		return false
	}
	conf.confFile = f
	return true
}

// Start is a method for reading configuraton.
func (conf *configManager) Start() {
	go run(conf)
}

// Stop is a method for stopping read configuration.
func (conf *configManager) Stop() {
	conf.stopRunning <- true
}

func run(conf *configManager) {
	for {
		time.Sleep(1 * time.Second)
		conf.load()
		select {
		case <-conf.stopRunning:
			break
		default:
			continue
		}
	}
}

var configure *configManager
var once sync.Once

// GetInstance returns configuraton singleton instance
func GetInstance() *configManager {
	once.Do(func() {
		configure = &configManager{}
		configure.Conf = &Config{}

	})
	return configure
}
