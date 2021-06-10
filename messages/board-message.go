package messages

import "practice-news-board-web/models"

type Board struct {
	Item []models.Board
	News News
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
