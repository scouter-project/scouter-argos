package instance

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
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

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

}
