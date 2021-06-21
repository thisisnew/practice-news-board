package processors

import (
	"practice-news-board-web/database"
	"practice-news-board-web/models"
)

func GetBoardGroups(useState bool) []models.BoardGroup {
	db := database.GetDB()

	var boardGroup []models.BoardGroup
	db.Table(models.BoardGroup{}.Table()).Where("use_state = ?", useState).Find(&boardGroup)

	return boardGroup
}
