package config

import (
	sqlserver "MANCHESTER/API-GLOBAL-LINX/serversql"
	"net/http"
)

//Config ...
type Config struct {
	APIPort int `yaml:"api port"`
	Linx    SQL `yaml:"linx"`
	App     SQL `yaml:"fin"`
}

//Auth ...
type Auth struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}

//SQL ...
type SQL struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Auth `yaml:",inline"`

	Db string `yaml:"db"`
}
type linx sqlserver.SQLStr

//Controller ...
type Controller struct {
	//conf config.Config
	linx *linx

	server *http.Server
}
