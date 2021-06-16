package processors

import (
	"practice-news-board-web/database"
	"practice-news-board-web/messages"
	"practice-news-board-web/models"
	util "practice-news-board-web/utils"
)

func GetComments(boardId string) []messages.Comment {
	db := database.GetDB()

	var comments []messages.Comment

	var data []models.Comment
	db.Table(models.Comment{}.Table()).Where("board_id = ?", boardId).Find(&data)

	for _, c := range data {
		commenter := GetUserById(c.CommenterId)
		comments = append(comments, messages.Comment{
			CommentId: c.CommentId,
			Commenter: messages.Commenter{
				Rank: commenter.UserRank,
				Name: commenter.UserName,
			},
			Contents:    c.Contents,
			CreateDate:  c.CreateDate.Format(util.TIME_LAYOUT_YYYY_MM_DD_HH_MM_SS),
			UpdateDate:  c.UpdateDate.Format(util.TIME_LAYOUT_YYYY_MM_DD_HH_MM_SS),
			CommentHide: c.CommentHide,
			BoardId:     c.PostId,
		})
	}

	return comments
}
