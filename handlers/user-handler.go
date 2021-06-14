package handlers

import (
	"net/http"
	"practice-news-board-web/database"
	"practice-news-board-web/models"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	var user models.User
	result := db.Table(models.User{}.Table()).Where("user_email = ? and user_pw = ? and user_state = ?", "thisisnew@naver.com", "1234", true).Find(&user)

	if result != nil {
		http.Redirect(w, r, "/board", http.StatusOK)
	} else {

	}
}
