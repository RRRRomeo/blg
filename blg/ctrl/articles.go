package ctrl

import (
	"blg/blg/db"
	"blg/blg/dto"
	"blg/blg/model"
	"blg/blg/resp"
	"blg/types"
	"fmt"
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleServicer struct {
	query *ArticleQuery
}

// 新建articlesServicer
func NewArticleServicer(q *ArticleQuery) *ArticleServicer {
	return &ArticleServicer{
		query: q,
	}
}

// 获取model.user
func (s *ArticleServicer) GetAuthor(id int) (*model.User, error) {
	art, err := s.GetArticle(id)
	if err != nil {
		return nil, err
	}

	authorId := art.AuthorId
	dbp := db.GetDB()
	author := &model.User{}

	if err := dbp.Model(&model.User{}).Where("id = ?", authorId).First(author).Error; err != nil {
		return nil, err
	}

	return author, nil
}

// 获取tags([]string)
func (s *ArticleServicer) GetTags(id int) ([]string, error) {
	art, err := s.GetArticle(id)
	if err != nil {
		return nil, err
	}

	tags := []model.ArticleTagRelation{}
	dbp := db.GetDB()
	// 执行数据库查询，获取该篇文章对应的所有标签的记录
	if err := dbp.Model(&model.ArticleTagRelation{}).Where("article_id = ?", art.Id).Find(&tags).Error; err != nil {
		return nil, err
	}
	rettags := make([]string, 0)
	for _, tag := range tags {
		// tagId ==> tagname
		tmp := db.GetDB()
		out := &model.ArticleTag{}
		tmp.Model(&model.ArticleTag{}).Where("id = ?", tag.TagId).First(out)
		rettags = append(rettags, out.TagName)
	}

	return rettags, nil
}

// 获取category(string)
func (s *ArticleServicer) GetCategory(id int) (string, error) {
	art := &model.Article{}
	category := &model.ArticleCategory{}
	dbp := db.GetDB()
	if err := dbp.Model(&model.Article{}).Where("id = ?", id).First(art).Error; err != nil {
		fmt.Printf("find category_id fail:%s\n", err)
		return "", err
	}

	dbpp := db.GetDB()
	if err := dbpp.Model(&model.ArticleCategory{}).Where("id = ?", art.CategoryId).First(category).Error; err != nil {
		return "", err
	}

	return category.CategoryName, nil
}

// 获取model.ArcileBody
func (s *ArticleServicer) GetArticleBody(id int) (*model.ArticleBody, error) {
	art, _ := s.GetArticle(id)
	bodyId := art.BodyId

	dbp := db.GetDB()
	body := &model.ArticleBody{}
	if err := dbp.Model(&model.ArticleBody{}).Where("id = ?", bodyId).First(body).Error; err != nil {
		return nil, err
	}
	return body, nil
}

// 获取model.Article model.ArticleTag
func (s *ArticleServicer) GetSelectArticleCategory(id int) (*model.Article, *[]model.ArticleTag, error) {
	// 返回的是tagsId和categoryId
	art := &model.Article{}
	tags := &[]model.ArticleTag{}
	dbp := db.GetDB()
	if err := dbp.Where("id = ?", id).First(art).Error; err != nil {
		fmt.Printf("find category_id fail:%s\n", err)
		return nil, nil, err
	}

	dbpp := db.GetDB()
	if err := dbpp.Model(&model.ArticleTag{}).Where("article_id = ?", id).Find(tags).Error; err != nil {
		fmt.Printf("find tag_id fail:%s\n", err)
		return nil, nil, err
	}

	return art, tags, nil
}

func (s *ArticleServicer) Publisher(req *types.ReqArticle, user *model.User, dbarticle *model.Article) bool {
	if req.Title == "" || req.Category == "" {
		return false
	}

	dbArticleBody := &model.ArticleBody{
		Content:     req.Content,
		ContentHtml: req.ContentHtml,
	}

	categoryId := 0
	if !checkCategoryExist(req.Category, &categoryId) {
		fmt.Printf("category dont exist!\n")
		return false
	}

	fmt.Printf("categoryid :%d\n", categoryId)
	// insert body into db
	dbp := db.GetDB()
	if err := dbp.Model(&model.ArticleBody{}).Create(&dbArticleBody).Error; err != nil {
		out := fmt.Sprintf("create body into db fail:%s", err)
		fmt.Printf("%s\n", out)
		return false
	}

	if !dto.ArticleDto(dbarticle, req, user.Id, categoryId, req.Tags, dbArticleBody) {
		return false
	}

	dbpp := db.GetDB()
	if err := dbpp.Model(&model.Article{}).Create(dbarticle).Error; err != nil {
		fmt.Printf("publish article fail:%s\n", err)
		return false
	}

	return true
}

func (s *ArticleServicer) Updater(req *types.ReqArticle, user *model.User, articleId int, dbarticle *model.Article) bool {
	if req.Title == "" || req.Category == "" {
		return false
	}

	dbp := db.GetDB()
	if err := dbp.Model(&model.Article{}).Where("id = ?", articleId).First(dbarticle).Error; err != nil {
		fmt.Printf("find article fail:%s\n", err)
		return false
	}

	body := &model.ArticleBody{}
	if err := dbp.Model(&model.ArticleBody{}).Where("id = ?", dbarticle.BodyId).First(body).Error; err != nil {
		fmt.Printf("find article body fail:%s\n", err)
		return false
	}

	// update body
	body.Content = req.Content
	body.ContentHtml = req.ContentHtml

	if err := dbp.Model(&model.ArticleBody{}).Save(body).Error; err != nil {
		fmt.Printf("update article fail:%s\n", err)
		return false
	}

	categoryId := getCategoryIdWithName(req.Category)
	updateArticleTags(dbarticle, req.Tags)

	// update article
	dbarticle.BodyId = int64(body.Id)
	dbarticle.CategoryId = categoryId
	dbarticle.Summary = req.Summary
	dbarticle.Title = req.Title
	dbarticle.ViewCounts += 1
	return true
}

// 获取model.Article
func (s *ArticleServicer) GetArticle(id int) (*model.Article, error) {
	dbp := db.GetDB()
	art := &model.Article{}

	if err := dbp.Where("id = ?", id).First(art).Error; err != nil {
		fmt.Printf("find article using id fail:%s\n", err)
		return nil, err
	}

	return art, nil
}

// 通过id删除文章
func (s *ArticleServicer) DeleteArticle(id int) (bool, error) {
	art, err := s.GetArticle(id)
	if art == nil && err != nil {
		fmt.Printf("get article by id fail:%s\n", err)
		return false, err
	}

	dbp := db.GetDB()
	if err := dbp.Model(&model.Article{}).Delete(art).Error; err != nil {
		fmt.Printf("dele article fail:%s\n", err)
		return false, err
	}

	return true, nil
}

// 分页查询core
func (s *ArticleServicer) ListArticles() ([]model.Article, int64) {
	// 分页查询
	query := s.query
	offset := (query.PageNumber - 1) * query.PageSize
	limit := query.PageSize
	relations := []model.ArticleTagRelation{}
	relationIds := make([]int, 0)

	dbp := db.GetDB()

	// 添加标签筛选条件
	if query.Tag != "" {
		tag := &model.ArticleTag{}
		dbpp := db.GetDB()
		dbpp.Model(&model.ArticleTag{}).Where("tagname = ?", query.Tag).First(tag)
		if tag.Id != 0 {
			// dbp = dbp.Where("tag_id = ?", tag.Id)
			// get articlesId in relation table
			tmp := db.GetDB()

			tmp.Model(&model.ArticleTagRelation{}).Where("tag_id = ?", tag.Id).Find(&relations)
			if tmp.Error != nil {
				fmt.Printf("tag err:%s\n", tmp.Error)
				return nil, 0
			}
		}
	}

	// 添加分类筛选条件
	if query.Category != "" {
		category := &model.ArticleCategory{}
		dbpp := db.GetDB()
		dbpp.Model(&model.ArticleCategory{}).Where("categoryname = ?", query.Category).First(category)

		if category.Id != 0 {
			dbp = dbp.Where("category_id = ?", category.Id)
			if dbp.Error != nil {
				fmt.Printf("category err:%s\n", dbp.Error)
				return nil, 0
			}
		}
	}

	// 添加时间筛选条件
	if query.Year != "" && query.Month != "" {
		dbp = dbp.Where("YEAR(create_date) = ? AND MONTH(create_date) = ?", query.Year, query.Month)
		if dbp.Error != nil {
			fmt.Printf("year month err:%s\n", dbp.Error)
			return nil, 0
		}
	}

	// 添加名称筛选条件
	if query.Name != "" {
		dbp = dbp.Where("name LIKE ?", "%"+query.Name+"%")
	}

	// 执行分页查询
	var articles []model.Article
	if len(relations) == 0 {
		if err := dbp.Offset(offset).Limit(limit).Find(&articles).Error; err != nil {
			fmt.Printf("get offset fail:%s\n", err)
			return nil, 0
		}
	} else {
		for _, v := range relations {
			relationIds = append(relationIds, v.ArticleId)
		}

		if err := dbp.Where("id IN (?)", relationIds).Offset(offset).Limit(limit).Find(&articles).Error; err != nil {
			fmt.Printf("get offset fail:%s\n", err)
			return nil, 0
		}
	}
	// 计算总记录数
	var total int64
	if err := dbp.Model(&model.Article{}).Count(&total).Error; err != nil {
		// ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		fmt.Printf("get count fail:%s\n", err)
		return nil, 0
	}
	return articles, total
}

/**TODO: 以下方法应该是articles_service.go中的service方法, 后面移动到articles_service.go中去*/
// 查询
type ArticleQuery struct {
	Category   string `json:"category,omitempty"`
	Tag        string `json:"tag,omitempty"`
	Year       string `json:"year,omitempty"`
	Month      string `json:"month,omitempty"`
	PageNumber int    `json:"page_number,omitempty"`
	PageSize   int    `json:"page_size,omitempty"`
	Name       string `json:"name,omitempty"` // TODO:?
	Sort       string `json:"sort,omitempty"`
}

// 根据查询条件获取文章
func GetArticles(ctx *gin.Context) {
	articleQuery := &ArticleQuery{} // 定义ArticleVo结构体，根据需要修改字段和类型

	err := ctx.ShouldBindJSON(articleQuery) // 将查询参数绑定到article结构体
	if err != nil {
		// 处理错误
		resp.Fail(ctx, nil, err.Error())
		return
	}

	articleServicer := NewArticleServicer(articleQuery)

	dbarticles, total := articleServicer.ListArticles()

	if dbarticles == nil {
		resp.Fail(ctx, nil, "get articles fail!")
		return
	}

	// TODO: dto
	rsparticles := make([]types.RspArticle, len(dbarticles))

	for i, v := range dbarticles {
		curCategory, _ := articleServicer.GetCategory(v.Id)
		curTags, _ := articleServicer.GetTags(v.Id)
		curAuthor, _ := articleServicer.GetAuthor(v.Id)
		curBody, _ := articleServicer.GetArticleBody(v.Id)

		if !dto.ArticleOtd(&rsparticles[i], &v, curAuthor, curBody, curCategory, curTags) {
			fmt.Printf("dto articles fail!\n")
			resp.Fail(ctx, nil, "dto articles fail")
			return
		}

	}
	// 构建分页结果
	result := gin.H{
		"articles": rsparticles,
		"meta": gin.H{
			"total":      total,
			"page":       articleQuery.PageNumber,
			"page_size":  articleQuery.PageSize,
			"total_page": int(math.Ceil(float64(total) / float64(articleQuery.PageSize))),
		},
	}
	resp.Success(ctx, result, "get articles success!")
}

// 获取Hot标签文章
func GetArticlesHot(ctx *gin.Context) {
	articleQuery := &ArticleQuery{} // 定义ArticleVo结构体，根据需要修改字段和类型

	err := ctx.ShouldBindJSON(articleQuery) // 将查询参数绑定到article结构体
	if err != nil {
		// 处理错误
		resp.Fail(ctx, nil, err.Error())
		return
	}

	articleServicer := NewArticleServicer(articleQuery)
	dbarticles, total := articleServicer.ListArticles()
	if dbarticles == nil {
		resp.Fail(ctx, nil, "get articles fail!")
		return
	}

	rsparticles := make([]types.RspArticle, len(dbarticles))

	for i, v := range dbarticles {
		curCategory, _ := articleServicer.GetCategory(v.Id)
		curTags, _ := articleServicer.GetTags(v.Id)
		curAuthor, _ := articleServicer.GetAuthor(v.Id)
		curBody, _ := articleServicer.GetArticleBody(v.Id)

		if !dto.ArticleOtd(&rsparticles[i], &v, curAuthor, curBody, curCategory, curTags) {
			fmt.Printf("dto articles fail!\n")
			resp.Fail(ctx, nil, "dto articles fail")
			return
		}

	}

	// 构建分页结果
	result := gin.H{
		"articles": rsparticles,
		"meta": gin.H{
			"total":      total,
			"page":       articleQuery.PageNumber,
			"page_size":  articleQuery.PageSize,
			"total_page": int(math.Ceil(float64(total) / float64(articleQuery.PageSize))),
		},
	}
	resp.Success(ctx, result, "get articles success!")
}

// 获取New标签文章
func GetArticlesNew(ctx *gin.Context) {
	articleQuery := &ArticleQuery{} // 定义ArticleVo结构体，根据需要修改字段和类型

	err := ctx.ShouldBindJSON(articleQuery) // 将查询参数绑定到article结构体
	if err != nil {
		// 处理错误
		resp.Fail(ctx, nil, err.Error())
		return
	}

	articleServicer := NewArticleServicer(articleQuery)
	dbarticles, total := articleServicer.ListArticles()
	if dbarticles == nil {
		resp.Fail(ctx, nil, "get articles fail!")
		return
	}

	rsparticles := make([]types.RspArticle, len(dbarticles))

	for i, v := range dbarticles {
		curCategory, _ := articleServicer.GetCategory(v.Id)
		curTags, _ := articleServicer.GetTags(v.Id)
		curAuthor, _ := articleServicer.GetAuthor(v.Id)
		curBody, _ := articleServicer.GetArticleBody(v.Id)

		if !dto.ArticleOtd(&rsparticles[i], &v, curAuthor, curBody, curCategory, curTags) {
			fmt.Printf("dto articles fail!\n")
			resp.Fail(ctx, nil, "dto articles fail")
			return
		}

	}

	// 构建分页结果
	result := gin.H{
		"articles": rsparticles,
		"meta": gin.H{
			"total":      total,
			"page":       articleQuery.PageNumber,
			"page_size":  articleQuery.PageSize,
			"total_page": int(math.Ceil(float64(total) / float64(articleQuery.PageSize))),
		},
	}
	resp.Success(ctx, result, "get articles success!")
}

// 通过文章id获取文章内容
func GetSelectArticleView(ctx *gin.Context) {
	// 获取 URL 参数中的文章 ID
	articleIdStr := ctx.Param("id")
	articleId, _ := strconv.Atoi(articleIdStr)
	articleQuery := &ArticleQuery{}         // 定义ArticleVo结构体，根据需要修改字段和类型
	err := ctx.ShouldBindJSON(articleQuery) // 将查询参数绑定到article结构体
	if err != nil {
		// 处理错误
		resp.Fail(ctx, nil, err.Error())
		return
	}

	articleServicer := NewArticleServicer(articleQuery)
	// 根据文章 ID 查询数据库获取文章详情
	dbarticle, err := articleServicer.GetArticle(articleId)
	if err != nil {
		// 处理查询错误
		resp.Fail(ctx, nil, "search fail!")
		return
	}

	rsparticle := &types.RspArticle{}

	curCategory, _ := articleServicer.GetCategory(dbarticle.Id)
	curTags, _ := articleServicer.GetTags(dbarticle.Id)
	curAuthor, _ := articleServicer.GetAuthor(dbarticle.Id)
	curBody, _ := articleServicer.GetArticleBody(dbarticle.Id)

	if !dto.ArticleOtd(rsparticle, dbarticle, curAuthor, curBody, curCategory, curTags) {
		fmt.Printf("dto articles fail!\n")
		resp.Fail(ctx, nil, "dto articles fail")
		return
	}

	// 返回文章详情
	resp.Success(ctx, gin.H{"artcile": rsparticle}, "get article success!")
}

// 通过文章id获取文章分类
func GetSelectArticleCategory(ctx *gin.Context) {
	// 获取 URL 参数中的文章 ID
	articleIdStr := ctx.Param("id")
	articleId, _ := strconv.Atoi(articleIdStr)
	articleQuery := &ArticleQuery{}         // 定义ArticleVo结构体，根据需要修改字段和类型
	err := ctx.ShouldBindJSON(articleQuery) // 将查询参数绑定到article结构体
	if err != nil {
		// 处理错误
		resp.Fail(ctx, nil, err.Error())
		return
	}

	articleServicer := NewArticleServicer(articleQuery)
	// 根据文章 ID 查询数据库获取文章详情
	art, tags, err := articleServicer.GetSelectArticleCategory(articleId)
	if err != nil {
		// 处理查询错误
		resp.Fail(ctx, nil, "search fail!")
		return
	}
	t := make([]int, 0)
	for _, tag := range *tags {
		t = append(t, tag.Id)
	}
	rsp := &types.RspArticleCategory{
		CategoryId: art.CategoryId,
		TagId:      t,
	}

	// 返回文章详情
	resp.Success(ctx, gin.H{"category": rsp}, "get article success!")
}

// 上传文章
func PublishArticle(ctx *gin.Context) {
	req := &types.ReqArticle{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		resp.Fail(ctx, nil, err.Error())
		return
	}

	user, ok := ctx.Get("user")
	if !ok {
		resp.Fail(ctx, nil, "get user fail")
	}

	// TODO: 校验articles参数合法性,后面转到mid去

	s := NewArticleServicer(nil)
	dbArticle := &model.Article{}
	if req.Id == 0 {
		if !s.Publisher(req, user.(*model.User), dbArticle) {
			resp.Fail(ctx, nil, "publish article fail!")
			return
		}
	} else {
		if !s.Updater(req, user.(*model.User), req.Id, dbArticle) {
			resp.Fail(ctx, nil, "update article fail")
			return
		}
	}

	resp.Success(ctx, nil, "publish or update success!")
}

// 打包文章
func GetListArchives(ctx *gin.Context) {
	// "select year(create_date) as year,month(create_date) as month,count(*) as count from me_article group by year(create_date),month(create_date)"
}

// 通过文章id获取文章
func GetArticleById(ctx *gin.Context) {
	// 获取 URL 参数中的文章 ID
	articleIdStr := ctx.Param("id")
	articleId, _ := strconv.Atoi(articleIdStr)
	articleQuery := &ArticleQuery{}         // 定义ArticleVo结构体，根据需要修改字段和类型
	err := ctx.ShouldBindJSON(articleQuery) // 将查询参数绑定到article结构体
	if err != nil {
		// 处理错误
		resp.Fail(ctx, nil, err.Error())
		return
	}

	articleServicer := NewArticleServicer(articleQuery)
	// 根据文章 ID 查询数据库获取文章详情
	dbarticle, err := articleServicer.GetArticle(articleId)
	if err != nil {
		// 处理查询错误
		resp.Fail(ctx, nil, "search fail!")
		return
	}

	rsparticle := &types.RspArticle{}

	curCategory, _ := articleServicer.GetCategory(dbarticle.Id)
	curTags, _ := articleServicer.GetTags(dbarticle.Id)
	curAuthor, _ := articleServicer.GetAuthor(dbarticle.Id)
	curBody, _ := articleServicer.GetArticleBody(dbarticle.Id)

	if !dto.ArticleOtd(rsparticle, dbarticle, curAuthor, curBody, curCategory, curTags) {
		fmt.Printf("dto articles fail!\n")
		resp.Fail(ctx, nil, "dto articles fail")
		return
	}

	// 返回文章详情
	resp.Success(ctx, gin.H{"artcile": rsparticle}, "get article success!")
}

// 通过文章id获取tags
func GetSelectArticleTag(ctx *gin.Context) {
	// 获取 URL 参数中的文章 ID
	articleIdStr := ctx.Param("id")
	articleId, _ := strconv.Atoi(articleIdStr)
	articleQuery := &ArticleQuery{}         // 定义ArticleVo结构体，根据需要修改字段和类型
	err := ctx.ShouldBindJSON(articleQuery) // 将查询参数绑定到article结构体
	if err != nil {
		// 处理错误
		resp.Fail(ctx, nil, err.Error())
		return
	}

	articleServicer := NewArticleServicer(articleQuery)
	// 根据文章 ID 查询数据库获取文章详情
	art, tags, err := articleServicer.GetSelectArticleCategory(articleId)
	if err != nil {
		// 处理查询错误
		resp.Fail(ctx, nil, "search fail!")
		return
	}
	t := make([]int, 0)
	for _, tag := range *tags {
		t = append(t, tag.Id)
	}
	rsp := &types.RspArticleCategory{
		CategoryId: art.CategoryId,
		TagId:      t,
	}

	// 返回文章详情
	resp.Success(ctx, gin.H{"category": rsp}, "get article success!")
}

// 更新文章
func UpdateArticle(ctx *gin.Context) {
	req := &types.ReqArticle{}
	err := ctx.ShouldBindJSON(req)
	if err != nil {
		resp.Fail(ctx, nil, err.Error())
		return
	}

	user, ok := ctx.Get("user")
	if !ok {
		resp.Fail(ctx, nil, "get user fail")
	}

	// TODO: 校验articles参数合法性,后面转到mid去

	s := NewArticleServicer(nil)
	dbArticle := &model.Article{}
	if !s.Updater(req, user.(*model.User), req.Id, dbArticle) {
		resp.Fail(ctx, nil, "update article fail")
		return
	}

	resp.Success(ctx, nil, "publish or update success!")
}

// 删除文章
func DeleteArticleByID(ctx *gin.Context) {
	// 获取 URL 参数中的文章 ID
	articleIdStr := ctx.Param("id")
	articleId, _ := strconv.Atoi(articleIdStr)
	articleQuery := &ArticleQuery{}         // 定义ArticleVo结构体，根据需要修改字段和类型
	err := ctx.ShouldBindJSON(articleQuery) // 将查询参数绑定到article结构体
	if err != nil {
		// 处理错误
		resp.Fail(ctx, nil, err.Error())
		return
	}

	articleServicer := NewArticleServicer(articleQuery)
	ok, err := articleServicer.DeleteArticle(articleId)
	if !ok && err != nil {
		resp.Fail(ctx, nil, "delete article fail: "+err.Error())
		return
	}
	resp.Success(ctx, nil, "detele article success!")
}

/**=======================TOOLS_FUNC=======================*/

func checkCategoryExist(category string, id *int) bool {
	if category == "" {
		return false
	}
	ca := &model.ArticleCategory{}
	dbp := db.GetDB()
	if err := dbp.Model(&model.ArticleCategory{}).Where("categoryname = ?", category).First(ca).Error; err != nil {
		// category dont exist insert in db
		// TODO: change the func from category ctrl
		ca.CategoryName = category
		ca.Avatar = category
		ca.Description = category

		if err := dbp.Model(&model.ArticleCategory{}).Create(ca).Error; err != nil {
			fmt.Printf("create category fail:%s\n", err)
			return false
		}

	}
	*id = ca.Id

	return true
}

func getCategoryIdWithName(category string) int {
	if category == "" {
		return 0
	}

	dbp := db.GetDB()
	ca := &model.ArticleCategory{}
	if err := dbp.Model(&model.ArticleCategory{}).Where("categoryname = ?", category).First(ca).Error; err != nil {
		fmt.Printf("get category_id fail:%s\n", err)
		return 0
	}

	return ca.Id
}

func updateArticleTags(dbarticle *model.Article, tags []string) bool {
	if dbarticle == nil || tags == nil {
		return false
	}

	dbp := db.GetDB()
	rela := &[]model.ArticleTagRelation{}

	if err := dbp.Model(&model.ArticleTagRelation{}).Where("article_id = ?", dbarticle.Id).Find(rela).Error; err != nil {
		fmt.Printf("get artcile tag fail:%s\n", err)
		return false
	}

	for _, v := range tags {
		for _, r := range *rela {
			tag := &model.ArticleTag{}
			dbp.Model(&model.ArticleTag{}).Where("id = ?", r.TagId).First(tag)
			if v == tag.TagName {
				continue
			}

			curTag := &model.ArticleTag{}
			dbp.Model(&model.ArticleTag{}).Where("tagname = ?", v).First(curTag)
			add := &model.ArticleTagRelation{
				ArticleId: dbarticle.Id,
				TagId:     curTag.Id,
			}
			if err := dbp.Model(&model.ArticleTagRelation{}).Create(add).Error; err != nil {
				fmt.Printf("create tag relation fail:%s\n", err)
				return false
			}
		}
	}
	return true
}
