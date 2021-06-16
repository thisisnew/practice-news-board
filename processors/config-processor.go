package processors

import (
	"practice-news-board-web/database"
	"practice-news-board-web/models"
)

func GetBoardList() []models.Board {
	db := database.GetDB()

	var board []models.Board
	db.Table(models.Board{}.Table()).Find(&board)

	return board
}
