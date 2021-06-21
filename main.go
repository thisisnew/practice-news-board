package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"practice-news-board-web/handlers"
	util "practice-news-board-web/utils"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/auth/login", handlers.Login).Methods("POST")
	router.HandleFunc("/auth/refresh", handlers.Refresh).Methods("GET")

	router.HandleFunc("/board", handlers.GetBoardGroups).Methods("GET")
	router.HandleFunc("/board/{groupId}", handlers.GetBoardList).Methods("GET")
	router.HandleFunc("/board/{groupId}/{postId}", handlers.GetPost).Methods("GET")

	if err := http.ListenAndServe(config.Host, util.HttpHandler(router)); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
