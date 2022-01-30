package config

type Server struct {
	Port    string `default:"5555"`
	Address string `default:"0.0.0.0"`
	Path    string `default:"/v1/api"`
	Debug   bool   `default:"false"`
}
