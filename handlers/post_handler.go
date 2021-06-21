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
	postId := p["postId"]

	post := processors.GetPost(postId)

	comments := processors.GetComments(postId)

	boardDetail := messages.BoardDetail{
		Post:     post,
		Comments: comments,
	}

	json.NewEncoder(w).Encode(boardDetail)
}
