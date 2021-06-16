package models

type BoardGroup struct {
	ConfigId string
	BoardId  string
}

func (BoardGroup) Table() string {
	return "board_group"
}
