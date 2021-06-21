package models

type Board struct {
	BoardId      string
	BoardName    string
	BoardExplain string
	BoardNo      string
	BoardState   uint
}

func (Board) Table() string {
	return "board"
}
