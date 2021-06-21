package processors

import (
	"practice-news-board-web/database"
	"practice-news-board-web/messages"
	"practice-news-board-web/models"
)

func GetPost(postId string) messages.Board {
	db := database.GetDB()

	board := messages.Board{}
	db.Table(models.Board{}.Table()).Find(&board, postId)

	return board
}
