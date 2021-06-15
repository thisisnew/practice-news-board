package processors

import (
	"practice-news-board-web/database"
	"practice-news-board-web/messages"
	"practice-news-board-web/models"
)

func GetComments(boardId string) []messages.Comment {
	db := database.GetDB()

	var comments []messages.Comment
	db.Table(models.Comment{}.Table()).Where("board_id = ?", boardId).Find(&comments)

	return comments
}
