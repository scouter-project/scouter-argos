package main

import (
	"fmt"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/scouter-argos/agent/config"
	"github.com/scouter-argos/agent/manager"
)

var pidfile string
var f *os.File

func main() {
	readArgs(os.Args)
	start()

}

func start() {
	displayLogo()
	manager.ServiceStart()
	writePid()

	for fileExist(pidfile) {
		time.Sleep(1 * time.Second)
	}

	return
}

func stop() {
	if fileExist(pidfile) {
		os.Remove(pidfile)
	} else {
		fmt.Println("Agent is not running.")
	}
}

func displayLogo() {
	fmt.Println("MySQL Agent.")
}

func fileExist(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func writePid() error {
	path, err := os.Getwd()
	if err == nil {
		pidfile = path + "/" + strconv.Itoa(os.Getpid()) + ".scouter"
		f, err = os.Create(pidfile)
		if err != nil {
			return err
		}
		defer f.Close()
	} else {
		return nil
	}
	return nil
}

func usage() {
	fmt.Printf("need arguements")
}

func readArgs(args []string) {
	if len(args) > 2 {
		for _, value := range args {
			argsItem := strings.Split(value, "=")
			if argsItem[0] == "scouter.config" {
				config.GetInstance().ConfFilePath = argsItem[1]
			}
		}
	} else {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		config.GetInstance().ConfFilePath = path.Dir(ex) + "/conf/scouter.json"
	}
}
