package config

import (
	"MANCHESTER/TopBrands-app/api/mysql"
	"MANCHESTER/TopBrands-app/api/sqlserver"

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
type app mysql.SQLStr

//Controller ...
type Controller struct {
	//conf config.Config
	linx *linx
	app  *app

	server *http.Server
}
