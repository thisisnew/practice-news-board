package handlers

import (
	"encoding/json"
	"net/http"
	"practice-news-board-web/messages"
)

func GetBoardList(w http.ResponseWriter, r *http.Request) {
	var items messages.BoardList

	json.NewEncoder(w).Encode(items)
}
