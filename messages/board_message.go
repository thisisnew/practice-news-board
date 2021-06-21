package messages

import "practice-news-board-web/models"

const (
	DAY_WRITE_LIMIT = 5
)

type Board struct {
	BoardId      string `json:"boardId"`
	BoardName    string `json:"boardName"`
	BoardExplain string `json:"boardExplain"`
	BoardNo      string `json:"boardNo"`
	BoardState   uint   `json:"boardState"`
	GroupId      string `json:"groupId"`
}

type BoardList struct {
	Items   []models.Board
	IsLimit bool
	User    User
}

type BoardDetail struct {
	Post     Board
	Comments []Comment
}

type News struct {
	By          string
	Descendants uint
	Id          string
	Kids        []uint
	Parts       []uint
	Score       uint
	Time        uint
	Title       string
	Type        string
}
