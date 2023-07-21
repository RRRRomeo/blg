package dto

import (
	"blg/blg/model"
	"blg/types"
	"time"
)

// 转换请求数据生成数据类型
func ArticleDto(ar_db *model.Article, req *types.ReqArticle, authorId int64, categoryId int, tags []types.ReqTag, arbody *model.ArticleBody) bool {
	var err error
	ar_db.AuthorId = authorId
	ar_db.BodyId = int64(arbody.Id)
	ar_db.CategoryId = categoryId
	if err != nil {
		return false
	}

	ar_db.CommentCounts = 0
	ar_db.CreateDate = time.Now()
	ar_db.Summary = req.Summary
	ar_db.Title = req.Title
	ar_db.ViewCounts = 0
	ar_db.Weight = 1
	return true
}

// 转换数据类型生成将要发给客户端的响应
func ArticleItemOtd(rsp *types.RspArticleItem, dbo *model.Article, dbu *model.User, dbb *model.ArticleBody, tags []model.ArticleTag) bool {
	rsp.Author = types.RspAuthor{Name: dbu.Account}
	rsp.CommentCounts = dbo.CommentCounts
	rsp.CreateDate = dbo.CreateDate.String()
	rsp.Id = dbo.Id
	rsp.Weight = dbo.Weight
	rsp.Summary = dbo.Summary
	ts := make([]types.RspTags, len(tags))
	for i, v := range tags {
		ts[i].Id = v.Id
		ts[i].TagName = v.TagName
	}

	rsp.Tags = ts
	rsp.Title = dbo.Title
	rsp.ViewCount = dbo.ViewCounts
	return true
}

// 转换数据类型生成将要发给客户端的响应
func ArticleCardOtd(rsp *types.RspArticleCard, dbo *model.Article) bool {
	rsp.Id = dbo.Id
	rsp.Title = dbo.Title
	return true
}

func ArticleDetailOtd(rsp *types.RspArticleDetail, dbo *model.Article, dbu *model.User, dbb *model.ArticleBody, tags []model.ArticleTag, category model.ArticleCategory) bool {
	rsp.Id = dbo.Id
	rsp.Title = dbo.Title
	rsp.CommentCounts = dbo.CommentCounts
	rsp.ViewCount = dbo.ViewCounts
	rsp.Summary = dbo.Summary
	rsp.Author = types.RspAuthor{Name: dbu.Account}
	ts := make([]types.RspTags, len(tags))
	for i, v := range tags {
		ts[i].Id = v.Id
		ts[i].TagName = v.TagName
	}
	rsp.Tags = ts
	rsp.CreateDate = dbo.CreateDate.String()
	rsp.Category = types.RspCategory{Id: category.Id, CategoryName: category.CategoryName}
	rsp.Body = types.RspArticleBody{Content: dbb.Content, ContentHtml: dbb.ContentHtml}
	return true
}
