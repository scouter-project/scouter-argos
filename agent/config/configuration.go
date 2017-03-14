package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

type Configuration struct {
	stopRunning chan bool
}

var confFilePath string
var confFile *os.File
var running = make(chan bool)
var confFileModdTime time.Time
var confFileSize int64

type ConfObject struct {
	Configurations Collector
}

type Collector struct {
	IP        string       `json:"collector.ip"`
	Udpport   string       `json:"collector.udp.port"`
	Tcpport   string       `json:"collector.tcp.port"`
	Instances []DBInstance `json:"db.instances"`
}

type DBInstance struct {
	IP        string `json:"db.ip"`
	Port      string `json:"db.port"`
	User      string `json:"db.user"`
	Password  string `json:"db.password"`
	Slowquery string `json:"db.slowquery"`
}

func (conf *Configuration) load() {
	var confobj ConfObject
	fileInfo, e := confFile.Stat()
	if e != nil {
		//todo: error handling
	}

	if confFileModdTime != fileInfo.ModTime() || confFileSize != fileInfo.Size() {
		file, err := ioutil.ReadFile(confFilePath)
		if err != nil {
			//todo :
		}
		json.Unmarshal(file, &confobj)
		confFileSize = fileInfo.Size()
		confFileModdTime = fileInfo.ModTime()
	}

}

func (conf *Configuration) init() bool {
	f, err := os.Open(confFilePath)
	if err != nil {
		return false
	} else {
		confFile = f
	}
	return true
}

func (conf *Configuration) Start() {
	go conf.run()
}

func (conf *Configuration) Stop() {
	conf.stopRunning <- true
}

func (conf *Configuration) run() {
	for {
		conf.load()
		time.Sleep(1 * time.Second)
		select {
		case <-conf.stopRunning:
			break
		default:
			continue
		}
	}
}

var configure *Configuration
var once sync.Once

// GetInstance returns configuraton singleton instance
func GetInstance() *Configuration {
	once.Do(func() {
		configure = &Configuration{}
		if configure.init() {
			configure.Start()
		}
	})
	return configure
}
