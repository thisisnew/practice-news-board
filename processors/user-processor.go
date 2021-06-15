package processors

import (
	"practice-news-board-web/database"
	"practice-news-board-web/messages"
	"practice-news-board-web/models"
)

func GetUser() messages.User {
	db := database.GetDB()

	var user messages.User
	result := db.Table(models.User{}.Table()).Where("user_email = ? and user_pw = ? and user_state = ?", "thisisnew@naver.com", "1234", true).Find(&user)

	if result != nil {
		user.IsLogin = true
	}

	return user
}

func GetUserById(userId string) models.User {
	db := database.GetDB()

	var user models.User
	db.Table(models.User{}.Table()).Where("user_id = ?", userId).Find(&user)

	return user
}
