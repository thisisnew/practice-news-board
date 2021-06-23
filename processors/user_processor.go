package processors

import (
	"practice-news-board-web/database"
	"practice-news-board-web/messages"
	"practice-news-board-web/models"
)

func GetUser(id string, pw string, userState bool) messages.User {
	db := database.GetDB()

	var user messages.User
	result := db.Table(models.User{}.Table()).
		Where("user_email = ? and user_pw = ? and user_state = ?", id, pw, userState).
		Find(&user)

	if result != nil {
		user.IsLogin = true
	}

	return user
}

func GetUserById(userId string) models.User {
	db := database.GetDB()

	var user models.User
	db.Table(models.User{}.Table()).
		Where("user_id = ?", userId).
		Find(&user)

	return user
}
