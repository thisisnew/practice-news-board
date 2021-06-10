package handlers

import (
	"encoding/json"
	"net/http"
	"practice-news-board-web/database"
	"practice-news-board-web/messages"
	"practice-news-board-web/models"
)

func GetBoardList(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	var board []models.Board
	db.Table(models.Board{}.Table()).Find(&board)

	var items messages.Board
	items.Item = board

	json.NewEncoder(w).Encode(items)
}
