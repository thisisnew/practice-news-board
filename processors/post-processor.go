package processors

import (
	"practice-news-board-web/database"
	"practice-news-board-web/messages"
	"practice-news-board-web/models"
)

func GetPost(id string) messages.Board {
	db := database.GetDB()

	board := messages.Board{}
	db.Table(models.Board{}.Table()).Find(&board, id)

	return board
}
