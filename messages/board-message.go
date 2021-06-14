package messages

import "practice-news-board-web/models"

const (
	DAY_WRITE_LIMIT = 5
)

type Board struct {
	Item    []models.Board
	IsLimit bool
	News    []News
	User    User
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
