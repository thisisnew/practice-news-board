package models

import (
	"time"
)

type Comment struct {
	CommentId   string
	Commenter   string
	CommentHide bool
	Contents    string
	CreateDate  time.Time
	UpdateDate  time.Time
	BoardId     string
}

func (Comment) Table() string {
	return "comment"
}
