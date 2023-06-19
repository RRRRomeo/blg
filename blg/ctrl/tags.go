package ctrl

import (
	"blg/blg/db"
	"blg/blg/model"
	"blg/blg/resp"
	"fmt"

	"github.com/gin-gonic/gin"
)

// Serve.GlobalRouter.SetRouterGet("/tags", mid.UserAuth(), ctrl.GetTags)
// Serve.GlobalRouter.SetRouterGet("/tags/detail", mid.UserAuth(), ctrl.GetTagsDetail)
// Serve.GlobalRouter.SetRouterGet("/tags/hot", mid.UserAuth(), ctrl.GetTagsHot)
// Serve.GlobalRouter.SetRouterGet("/tags/:id", mid.UserAuth(), ctrl.GetSelectTags)
// Serve.GlobalRouter.SetRouterGet("/tags/detail/:id", mid.UserAuth(), ctrl.GetSelectTagDetail)
type Tags struct {
	model.ArticleTag
	Articles int `json:"articles,omitempty"`
}

// 获取所有tags
func GetTags(ctx *gin.Context) {
	tags := &[]model.ArticleTag{}
	dbp := db.GetDB()
	if err := dbp.Model(&model.ArticleTag{}).Find(tags).Error; err != nil {
		fmt.Printf("get all tags fail\n")
		resp.Fail(ctx, nil, "get all tags fail")
		return
	}

	resp.Success(ctx, gin.H{"tags": tags}, "get all tags success!")
}

func GetTagsDetail(ctx *gin.Context) {
	tagDetails := &[]Tags{}
	dbp := db.GetDB()
	if err := dbp.Table("me_article_tag at").
		Select("t.*, COUNT(at.tag_id) AS articles").
		Joins("RIGHT JOIN me_tag t ON at.tag_id = t.id").
		Group("t.id").
		Scan(tagDetails).Error; err != nil {
		fmt.Printf("get tag detail fail\n")
		resp.Fail(ctx, nil, "get tag detail fail!")
		return
	}

	resp.Success(ctx, gin.H{"tagsDetail": tagDetails}, "get tagsDetail success!")
}
