package models

type BoardGroup struct {
	BoardGroupId string
	GroupId      string
	Board        Board
}

func (BoardGroup) Table() string {
	return "board_group"
}
