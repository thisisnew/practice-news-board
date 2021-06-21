package handlers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"net/http"
	"practice-news-board-web/messages"
	"practice-news-board-web/processors"
)

func GetBoardList(w http.ResponseWriter, r *http.Request) {
	var items messages.BoardList

	cookie, err := r.Cookie("board-web-token")
	if err != nil {
		if err == http.ErrNoCookie {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	tokenStr := cookie.Value

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims,
		func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if tkn.Valid {
		items.User = messages.User{
			UserEmail: claims.UserEmail,
		}
	}

	p := mux.Vars(r)
	groupId := p["groupId"]

	boardList := processors.GetBoardList(groupId)

	var isLimit bool
	if len(boardList) == messages.DAY_WRITE_LIMIT {
		isLimit = true
	}

	items.Items = boardList
	items.IsLimit = isLimit

	json.NewEncoder(w).Encode(items)
}
