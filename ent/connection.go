package ent

type Config struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`

	SSHUser     string `json:"ssh-user"`
	Server      string `json:"server-ip"`
	SSHPassword string `json:"ssh-password"`
}

type SSH struct {
	SSHUser     string `json:"ssh-user"`
	Server      string `json:"server"`
	SSHPassword string `json:"ssh-password"`
}
