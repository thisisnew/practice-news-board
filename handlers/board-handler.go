package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"practice-news-board-web/messages"
	"practice-news-board-web/processors"
)

func GetBoardList(w http.ResponseWriter, r *http.Request) {
	user := GetUser()

	list := processors.GetBoardList()

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

func GetBoardDetail(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	id := p["boardId"]

	board := processors.GetBoardDetail(id)

	if board.BoardId == "" {
		json.NewEncoder(w).Encode(map[string]string{"result": "데이터가 없습니다."})
		return
	}

	comments := processors.GetComments()

	boardDetail := messages.BoardDetail{
		Board:    board,
		Comments: comments,
	}

	json.NewEncoder(w).Encode(boardDetail)
}
