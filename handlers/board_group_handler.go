package handlers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"practice-news-board-web/messages"
	"practice-news-board-web/processors"
)

func GetBoardGroups(w http.ResponseWriter, r *http.Request) {
	var items messages.BoardGroup

	boardGroups := processors.GetBoardGroups(true)

	for _, bg := range boardGroups {
		board := processors.GetBoardList(bg.GroupId)
		for _, b := range board {
			items.Items = append(items.Items, b)
		}
	}
	items.News = processors.GetLatestNews()

	cookie, err := r.Cookie("board-web-token")
	if err != nil {
		json.NewEncoder(w).Encode(items)
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

	json.NewEncoder(w).Encode(items)
}

func InsertBoardGroup(w http.ResponseWriter, r *http.Request) {
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

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	///////////////////////////
}

func UpdateBoardGroup(w http.ResponseWriter, r *http.Request) {
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

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	///////////////////////////
}

func DeleteBoardGroup(w http.ResponseWriter, r *http.Request) {
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

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	///////////////////////////
}
