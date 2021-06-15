package messages

import "time"

type Comment struct {
	CommentId   string    `json:"commentId"`
	Commenter   string    `json:"commenter"`
	Contents    string    `json:"contents"`
	CreateDate  time.Time `json:"create_date"`
	UpdateDate  time.Time `json:"update_date"`
	CommentHide bool      `json:"commentHide"`
	BoardId     string    `json:"boardId"`
}
