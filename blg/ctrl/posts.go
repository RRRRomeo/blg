package ctrl

import (
	"blg/blg/db"
	"blg/blg/dto"
	"blg/blg/model"
	"blg/blg/resp"
	"blg/types"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TODO: after the mid auth
func CreatePost(ctx *gin.Context) {
	dbp := db.GetDB()
	user, ok := ctx.Get("user")
	if !ok {
		resp.Fail(ctx, nil, "get user fail")
	}

	reqPost := &types.ReqCreatePost{}
	if err := ctx.ShouldBindJSON(reqPost); err != nil {
		resp.Fail(ctx, nil, "get req create post fail!")
		return
	}

	// dto
	dbPost := &model.Post{}
	if !dto.PostDto(dbPost, reqPost, user.(*model.User)) {
		resp.Fail(ctx, nil, "dto fail!")
		return
	}

	if searcher := dbp.Where("author_id = ?", dbPost.Author_id).Where("title = ?", dbPost.Title).First(dbPost).Error; searcher != nil {
		if searcher == gorm.ErrRecordNotFound {
			// add to db
			if err := dbp.Create(dbPost).Error; err != nil {
				resp.Fail(ctx, nil, "insert post into db fail!"+err.Error())
				return
			}
			resp.Success(ctx, nil, "create post success!")
			return
		}

		// the err is others
		resp.Fail(ctx, nil, "get post from db fail!")
	} else {
		// already exist
		resp.Fail(ctx, nil, "the post already exist!")
	}
}

func UpdatePost(ctx *gin.Context) {
	dbp := db.GetDB()
	user, ok := ctx.Get("user")
	if !ok {
		resp.Fail(ctx, nil, "get user fail")
	}

	reqUpdatePost := &types.ReqUpdatePost{}
	if err := ctx.ShouldBindJSON(reqUpdatePost); err != nil {
		resp.Fail(ctx, nil, "get req create post fail!")
		return
	}

	dbPost := &model.Post{}
	if searcher := dbp.Where("id = ?", reqUpdatePost.PostId).First(dbPost).Error; searcher != nil {
		if searcher == gorm.ErrRecordNotFound {
			resp.Fail(ctx, nil, "post dont exist!")
			return
		}

		// the err is others
		resp.Fail(ctx, nil, "get post from db fail!")
	}

	if dbPost.Author_id != user.(*model.User).Id {
		resp.Fail(ctx, nil, "the post cant be modify!")
		return
	}

	dbPost.Content = reqUpdatePost.Content
	dbPost.Title = reqUpdatePost.Title
	dbPost.Is_public = reqUpdatePost.IsPublic
	dbPost.Updated_at = time.Now()

	if err := dbp.Save(dbPost).Error; err != nil {
		resp.Response(ctx, http.StatusInternalServerError, 500, nil, "update fail:"+err.Error())
		return
	}
	resp.Success(ctx, nil, "update post success!")
}

func GetPosts(c *gin.Context) {

}
