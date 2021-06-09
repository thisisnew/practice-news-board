package main

import (
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

	http.HandleFunc("/", handlers.GetBoardList)

	if err := http.ListenAndServe(config.HOST, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
