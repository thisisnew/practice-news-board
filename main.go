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
	router.HandleFunc("/", handlers.GetBoardList).Methods("GET")
	if err := http.ListenAndServe(config.HOST, util.HttpHandler(router)); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
