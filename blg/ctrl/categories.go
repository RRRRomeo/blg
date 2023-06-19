package ctrl

import (
	"blg/blg/db"
	"blg/blg/model"
	"blg/blg/resp"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CategoryQuery struct {
	PageNumber int    `json:"page_number,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	Sort       string `json:"sort,omitempty"`
}

type CategoryServicer struct {
	query *CategoryQuery
}

func NewCategoryServicer(q *CategoryQuery) *CategoryServicer {
	return &CategoryServicer{
		query: q,
	}
}

type Categories struct {
	model.ArticleCategory
	Articles int `json:"articles,omitempty"`
}

func (c *CategoryServicer) GetCategoriesDetail() *[]Categories {
	dbp := db.GetDB()
	categories := []Categories{}
	if err := dbp.Table("me_category c").
		Select("c.*, COUNT(a.category_id) as articles").
		Joins("LEFT JOIN me_article a ON a.category_id = c.id").
		Group("c.id").
		Scan(&categories).Error; err != nil {
		fmt.Printf("categories get fail:%s\n", err)
		return nil
	}

	return &categories
}

func (c *CategoryServicer) GetCategories() *[]model.ArticleCategory {
	dbp := db.GetDB()
	categories := &[]model.ArticleCategory{}

	if err := dbp.Model(&model.ArticleCategory{}).Find(categories).Error; err != nil {
		fmt.Printf("get categories fail\n")
		return nil
	}

	return categories
}
func (c *CategoryServicer) GetCategory(id int) *model.ArticleCategory {
	dbp := db.GetDB()
	category := &model.ArticleCategory{}

	if err := dbp.Model(&model.ArticleCategory{}).Where("id = ?", id).First(category).Error; err != nil {
		fmt.Printf("get categories fail\n")
		return nil
	}

	return category
}

func (c *CategoryServicer) GetCategoryDetail(id int) *Categories {
	dbp := db.GetDB()
	categories := &Categories{}
	if err := dbp.Table("me_category c").
		Select("c.*, COUNT(a.category_id) as articles").
		Joins("LEFT JOIN me_article a ON a.category_id = c.id").
		Where("c.id = ?", id).
		Group("c.id").
		Scan(categories).Error; err != nil {
		fmt.Printf("categories get fail:%s\n", err)
		return nil
	}

	return categories
}

// Serve.GlobalRouter.SetRouterGet("/categorys/detail/:id", mid.UserAuth(), ctrl.GetSelectCategorysDetail)

func GetCategoriesDetail(c *gin.Context) {
	categoryServicer := NewCategoryServicer(nil)

	categories := categoryServicer.GetCategoriesDetail()
	if categories == nil {
		resp.Fail(c, nil, "get categories fail!")
		return
	}

	resp.Success(c, gin.H{"categories": categories}, "get categories success!")
}

func GetCategories(c *gin.Context) {
	categoryServicer := NewCategoryServicer(nil)
	categories := categoryServicer.GetCategories()

	if categories == nil {
		resp.Fail(c, nil, "get categories fail")
		return
	}

	resp.Success(c, gin.H{"categories": categories}, "get categories success!")
}

// Serve.GlobalRouter.SetRouterGet("/categorys/:id", mid.UserAuth(), ctrl.GetSelectCategory)
func GetSelectCategory(c *gin.Context) {
	// 获取 URL 参数中的文章 ID
	articleIdStr := c.Param("id")
	articleId, _ := strconv.Atoi(articleIdStr)

	categoryServicer := NewCategoryServicer(nil)
	category := categoryServicer.GetCategory(articleId)
	if category == nil {
		fmt.Printf("get category fail!")
		resp.Fail(c, nil, "get category fail")
		return
	}

	resp.Success(c, gin.H{"category": category}, "get category success!")
}

func GetSelectCategoryDetail(c *gin.Context) {
	// 获取 URL 参数中的文章 ID
	articleIdStr := c.Param("id")
	articleId, _ := strconv.Atoi(articleIdStr)

	categoryServicer := NewCategoryServicer(nil)
	category := categoryServicer.GetCategoryDetail(articleId)
	if category == nil {
		fmt.Printf("get category fail!")
		resp.Fail(c, nil, "get category fail")
		return
	}

	resp.Success(c, gin.H{"category": category}, "get category success!")
}
