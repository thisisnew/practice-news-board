package handlers

import (
	"net/http"
	"practice-news-board-web/database"
)

type Board struct {
	BoardId      uint   `json:"boardId"`
	BoardName    string `json:"boardName"`
	BoardExplain string `json:"boardExplain"`
	BoardNo      string `json:"boardNo"`
	BoardState   uint   `json:"boardState"`
}

func GetBoardList(w http.ResponseWriter, r *http.Request) {
	db := database.GetDB()

	var board []Board
	db.Table("company").Find(&board)

	//var items Companies
	//items.Companies = companies
	//items.PageInfo = PageInfo{
	//	TotalRecord: 1,
	//	TotalPage:   1,
	//	Limit:       15,
	//	Page:        1,
	//	PrevPage:    1,
	//	NextPage:    1,
	//}
	//
	//json.NewEncoder(w).Encode(items)
}
