package handlers

import (
	"encoding/json"
	"net/http"
	"practice-news-board-web/messages"
	"practice-news-board-web/processors"
)

func GetBoardList(w http.ResponseWriter, r *http.Request) {
	list := processors.GetBoardList()

	var items messages.BoardList
	items.Items = list

	json.NewEncoder(w).Encode(items)
}
