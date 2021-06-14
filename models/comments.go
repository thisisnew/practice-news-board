package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	CommentId   string `json:"commentId"`
	Commenter   string `json:"boardName"`
	CommentHide bool   `json:"commentHide"`
	BoardId     string `json:"boardId"`
}

func (Comment) Table() string {
	return "comment"
}
