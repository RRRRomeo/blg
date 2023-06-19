package dto

import (
	"blg/blg/model"
	"blg/types"
	"time"
)

func CommentDto(req *types.ReqComment, dbc *model.Comment) bool {
	if req == nil || dbc == nil {
		return false
	}
	dbc.ArticleId = req.ArticleId
	dbc.AuthorId = req.AuthorId
	dbc.Content = req.Content
	dbc.CreateDate = time.Now()
	dbc.ParentId = req.ParentId
	dbc.ToUid = req.ToUid
	dbc.Level = 1
	return true
}
