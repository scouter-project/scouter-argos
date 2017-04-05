package instance

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/scouter-argos/agent/monsql"
)

type Instance struct {
	Type      string `json:"db.type"`
	IP        string `json:"db.ip"`
	Port      string `json:"db.port"`
	User      string `json:"db.user"`
	Password  string `json:"db.password"`
	Slowquery string `json:"db.slowquery"`
}

func (inst *Instance) StartMonitor() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", inst.User, inst.Password, inst.IP, inst.Port, "mysql")

	db, err := sql.Open("mysql", dsn) //// Open doesn't open a connection. Validate DSN data
	if err != nil {
		panic(err.Error()) //todo : change error handler instead of panic.
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		//Write log.
		return
	}

	go inst.run()
}

func (inst *Instance) run() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", inst.User, inst.Password, inst.IP, inst.Port, "mysql")
	db, err := sql.Open("mysql", dsn) //// Open doesn't open a connection. Validate DSN data
	if err != nil {
		return
	}
	for {
		resultRow, err := db.Query(monsql.GlobalStatusSQL)
		defer resultRow.Close()
		if err != nil {
			return
		}

		var statusName string
		var statusValue int64
		for resultRow.Next() {
			// get RawBytes from data
			err = resultRow.Scan(&statusName, &statusValue)
			if err != nil {
				panic(err.Error()) // proper error handling instead of panic in your app
			}

		}
	}
}
