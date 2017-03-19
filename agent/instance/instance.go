package instance

type Instance struct {
	Type      string `json:"db.type"`
	IP        string `json:"db.ip"`
	Port      string `json:"db.port"`
	User      string `json:"db.user"`
	Password  string `json:"db.password"`
	Slowquery string `json:"db.slowquery"`
}

func (*Instance) StartMonitor() {

}
