package ctrl

import (
	"blg/blg/db"
	"blg/blg/dto"
	"blg/blg/model"
	"blg/blg/resp"
	"blg/types"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCommentsByArticle(c *gin.Context) {
	idstr := c.Param("id")
	id, _ := strconv.Atoi(idstr)

	comments := &[]model.Comment{}
	dbp := db.GetDB()

	if err := dbp.Model(&model.Comment{}).Where("article_id = ?", id).Find(comments).Error; err != nil {
		fmt.Printf("get the comment fail")
		resp.Fail(c, nil, "get comment fail")
		return
	}

	resp.Success(c, gin.H{"comments": comments}, "get comments success!")
}

// Serve.GlobalRouter.SetRouterPost("/comments/article/change", mid.UserAuth(), ctrl.PublishComment)
func PublishComment(c *gin.Context) {
	reqcomment := &types.ReqComment{}
	if err := c.ShouldBindJSON(reqcomment); err != nil {
		fmt.Printf("get comment req fail!\n")
		resp.Fail(c, nil, "get comment req fail!")
		return
	}

	dbcomment := &model.Comment{}
	if !dto.CommentDto(reqcomment, dbcomment) {
		resp.Fail(c, nil, "comment dto fail!")
		return
	}

	dbp := db.GetDB()

	if err := dbp.Model(&model.Comment{}).Create(dbcomment).Error; err != nil {
		fmt.Printf("create comment fail\n")
		resp.Fail(c, nil, "create comment fail!")
		return
	}

	resp.Success(c, nil, "create comment success!")

}
