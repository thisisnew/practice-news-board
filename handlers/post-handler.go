package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"practice-news-board-web/messages"
	"practice-news-board-web/processors"
)

func GetPost(w http.ResponseWriter, r *http.Request) {
	p := mux.Vars(r)
	id := p["boardId"]

	board := processors.GetBoardDetail(id)

	if board.BoardId == "" {
		json.NewEncoder(w).Encode(map[string]string{"result": "데이터가 없습니다."})
		return
	}

	comments := processors.GetComments(id)

	boardDetail := messages.BoardDetail{
		Board:    board,
		Comments: comments,
	}

	json.NewEncoder(w).Encode(boardDetail)
}
