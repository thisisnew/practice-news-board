package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"practice-news-board-web/database"
	"practice-news-board-web/messages"
	"practice-news-board-web/models"
)

const (
	LIMIT = 5
)

func GetBoardList(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	var board []models.Board
	result := db.Table(models.Board{}.Table()).Find(&board)

	var isLimit bool
	if result.RowsAffected == LIMIT {
		isLimit = true
	}

	resp, err := http.Get("https://hacker-news.firebaseio.com/v0/item/126809.json?print=pretty")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var news messages.News
	json.Unmarshal([]byte(data), &news)

	var items messages.Board
	items.Item = board
	items.IsLimit = isLimit
	items.News = news

	json.NewEncoder(w).Encode(items)
}
