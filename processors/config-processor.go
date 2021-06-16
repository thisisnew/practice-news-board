package processors

import (
	"practice-news-board-web/database"
	"practice-news-board-web/messages"
	"practice-news-board-web/models"
)

func GetBoardList() []messages.Board {
	db := database.GetDB()

	var board []messages.Board
	db.Table(models.Board{}.Table()).Find(&board)

	return board
}
