package models

type Post struct {
	PostId      string
	PostTitle   string
	PostContent string
	BoardId     string
}

func (Post) Table() string {
	return "post"
}
