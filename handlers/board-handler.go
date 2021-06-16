package handlers

import (
	"encoding/json"
	"net/http"
	"practice-news-board-web/messages"
	"practice-news-board-web/processors"
)

func GetBoard(w http.ResponseWriter, r *http.Request) {
	user := GetUser()

	list := processors.GetBoard()

	var isLimit bool
	if len(list) == messages.DAY_WRITE_LIMIT {
		isLimit = true
	}

	news := processors.GetLatestNews()

	var items messages.BoardList
	items.Items = list
	items.IsLimit = isLimit
	items.News = news
	items.User = user

	json.NewEncoder(w).Encode(items)
}
