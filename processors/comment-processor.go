package processors

import (
	"practice-news-board-web/database"
	"practice-news-board-web/messages"
	"practice-news-board-web/models"
)

func GetComments() []messages.Comment {
	db := database.GetDB()

	var comments []messages.Comment
	db.Table(models.Comment{}.Table()).Find(&comments)

	return comments
}
