package handlers

import (
	"practice-news-board-web/messages"
	"practice-news-board-web/processors"
)

func GetUser() messages.User {
	return processors.GetUser()
}
