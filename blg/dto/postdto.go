package dto

import (
	"blg/blg/model"
	"blg/types"
	"time"
)

func PostDto(Postd *model.Post, postr *types.ReqCreatePost, curuser *model.User) bool {
	if Postd == nil || postr == nil {
		return false
	}

	Postd.Author_id = curuser.Id
	Postd.Created_at = time.Now()
	Postd.Is_public = postr.IsPublic
	Postd.Content = postr.Content
	Postd.Title = postr.Title
	Postd.Updated_at = time.Now()

	return true
}
