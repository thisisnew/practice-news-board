package messages

type Comment struct {
	CommentId   string `json:"commentId"`
	Commenter   string `json:"boardName"`
	CommentHide bool   `json:"commentHide"`
	BoardId     string `json:"boardId"`
}
