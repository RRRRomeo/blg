package dto

import (
	"blg/blg/db"
	"blg/blg/model"
	"blg/tools/common"
	"blg/types"
	"time"
)

// 转换请求数据生成数据类型
func ArticleDto(ar_db *model.Article, req *types.ReqArticle, authorId int64, category string, tags []string, arbody *model.ArticleBody) bool {
	var err error
	ar_db.AuthorId = authorId
	ar_db.BodyId = int64(arbody.Id)
	ar_db.CategoryId, err = common.FindCategoryId(db.GetDB(), category)
	if err != nil {
		return false
	}

	ar_db.CommentCounts = 0
	ar_db.CreateDate = time.Now()
	ar_db.Summary = req.Summary
	ar_db.Title = req.Title
	ar_db.ViewCount = 0
	return true
}

// 转换数据类型生成将要发给客户端的响应
func ArticleOtd(rsp *types.RspArticle, dbo *model.Article, dbu *model.User, dbb *model.ArticleBody, category string, tags []string) bool {
	rsp.Author = dbu.Account
	rsp.Body.Content = dbb.Content
	rsp.Body.ContentHtml = dbb.ContentHtml
	rsp.Category = category
	rsp.CommentCounts = dbo.CommentCounts
	rsp.CreateDate = dbo.CreateDate
	rsp.Id = dbo.Id
	rsp.Summary = dbo.Summary
	rsp.Tags = tags
	rsp.Title = dbo.Title
	rsp.ViewCount = dbo.ViewCount
	return true
}
