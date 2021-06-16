package models

type Config struct {
	ConfigId string
	BoardId  string
}

func (Config) Table() string {
	return "config"
}
