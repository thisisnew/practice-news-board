package models

type BoardGroup struct {
	BoardGroupId string
	GroupId      string
}

func (BoardGroup) Table() string {
	return "board_group"
}
