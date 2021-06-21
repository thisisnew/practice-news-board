package messages

import "practice-news-board-web/models"

type BoardGroup struct {
	Items []models.BoardGroup `json:"items"`
	News  []News              `json:"news"`
	User  User                `json:"user"`
}
