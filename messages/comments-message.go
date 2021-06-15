package messages

type Comment struct {
	CommentId   string    `json:"commentId"`
	Commenter   Commenter `json:"commenter"`
	Contents    string    `json:"contents"`
	CreateDate  string    `json:"create_date"`
	UpdateDate  string    `json:"update_date"`
	CommentHide bool      `json:"commentHide"`
	BoardId     string    `json:"boardId"`
}

type Commenter struct {
	Rank uint
	Name string
}
