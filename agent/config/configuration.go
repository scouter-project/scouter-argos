package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"
	"time"
)

// Configmgr is a object that manages configuration for argos.
type Configmgr struct {
	stopRunning chan bool
}

var confFilePath string
var confFile *os.File
var running = make(chan bool)
var confFileModdTime time.Time
var confFileSize int64

type configType struct {
	Configurations collectorType
}

type collectorType struct {
	CollectorIP string   `json:"collector.ip"`
	Udpport     string   `json:"collector.udp.port"`
	Tcpport     string   `json:"collector.tcp.port"`
	Instances   []dbType `json:"db.instances"`
}

type dbType struct {
	IP        string `json:"db.ip"`
	Port      string `json:"db.port"`
	User      string `json:"db.user"`
	Password  string `json:"db.password"`
	Slowquery string `json:"db.slowquery"`
}

func (conf *Configmgr) load() {
	var config configType
	fileInfo, e := confFile.Stat()
	if e != nil {
		//todo: error handling
	}

	if confFileModdTime != fileInfo.ModTime() || confFileSize != fileInfo.Size() {
		file, err := ioutil.ReadFile(confFilePath)
		if err != nil {
			//todo :
		}
		json.Unmarshal(file, &config)
		confFileSize = fileInfo.Size()
		confFileModdTime = fileInfo.ModTime()
	}

}

func (conf *Configmgr) init() bool {
	f, err := os.Open(confFilePath)
	if err != nil {
		return false
	} else {
		confFile = f
	}
	return true
}

func (conf *Configmgr) Start() {
	go conf.run()
}

func (conf *Configmgr) Stop() {
	conf.stopRunning <- true
}

func (conf *Configmgr) run() {
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

var configure *Configmgr
var once sync.Once

// GetInstance returns configuraton singleton instance
func GetInstance() *Configmgr {
	once.Do(func() {
		configure = &Configmgr{}
		if configure.init() {
			configure.Start()
		}
	})
	return configure
}
