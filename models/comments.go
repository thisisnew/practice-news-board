package models

import (
	"time"
)

type Comment struct {
	CommentId   string
	CommentHide bool
	Contents    string
	CreateDate  time.Time
	UpdateDate  time.Time
	BoardId     string
	CommenterId string
}

func (Comment) Table() string {
	return "comment"
}
