package processors

import (
	"practice-news-board-web/database"
	"practice-news-board-web/messages"
	"practice-news-board-web/models"
)

func GetComments(boardId string) []messages.Comment {
	db := database.GetDB()

	var comments []messages.Comment

	var data []models.Comment
	db.Table(models.Comment{}.Table()).Where("board_id = ?", boardId).Find(&data)

	layout := "2006-01-02 15:04:05"

	for _, c := range data {
		commenter := GetUserById(c.CommenterId)
		comments = append(comments, messages.Comment{
			CommentId: c.CommentId,
			Commenter: messages.Commenter{
				Rank: commenter.UserRank,
				Name: commenter.UserName,
			},
			Contents:    c.Contents,
			CreateDate:  c.CreateDate.Format(layout),
			UpdateDate:  c.UpdateDate.Format(layout),
			CommentHide: c.CommentHide,
			BoardId:     c.BoardId,
		})
	}

	return comments
}
