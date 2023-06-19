package serve

import (
	"blg/blg/ctrl"
	"blg/blg/mid"
	"blg/blg/router"
	"net/http"
	"sync"
	"sync/atomic"

	log "github.com/RRRRomeo/QLog/api"

	"github.com/gin-gonic/gin"
)

type BlgServe struct {
	Lock         sync.Mutex
	GlobalRouter *router.Router
	USERID       atomic.Int64
}

var Serve *BlgServe

func Init() bool {
	Serve = &BlgServe{
		GlobalRouter: router.NewRouter(gin.Default()),
	}
	return true
}

func GetBlgServe() *BlgServe {
	if Serve == nil {
		return nil
	}

	return Serve
}

func EventsHandler() bool {
	Serve.GlobalRouter.Engine.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, oauth-token, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})
	//=====================articles===========================/
	Serve.GlobalRouter.SetRouterGet("/articles", mid.UserAuth(), ctrl.GetArticles)
	Serve.GlobalRouter.SetRouterGet("/articles/hot", mid.UserAuth(), ctrl.GetArticlesHot)
	Serve.GlobalRouter.SetRouterGet("/articles/new", mid.UserAuth(), ctrl.GetArticlesNew)
	Serve.GlobalRouter.SetRouterGet("/articles/:id", mid.UserAuth(), ctrl.GetArticleById)
	Serve.GlobalRouter.SetRouterGet("/articles/view/:id", mid.UserAuth(), ctrl.GetSelectArticleView)
	Serve.GlobalRouter.SetRouterGet("/articles/category/:id", mid.UserAuth(), ctrl.GetSelectArticleCategory)
	Serve.GlobalRouter.SetRouterGet("/articles/tag/:id", mid.UserAuth(), ctrl.GetSelectArticleTag)
	Serve.GlobalRouter.SetRouterPost("/articles/publish", mid.UserAuth(), ctrl.PublishArticle)
	Serve.GlobalRouter.SetRouterPost("/articles/update", mid.UserAuth(), ctrl.UpdateArticle)
	Serve.GlobalRouter.SetRouterGet("/articles/delete/:id", mid.UserAuth(), ctrl.DeleteArticleByID)
	Serve.GlobalRouter.SetRouterGet("/articles/listArchives", mid.UserAuth(), ctrl.GetListArchives)

	//=====================category===========================/
	Serve.GlobalRouter.SetRouterGet("/categories", mid.UserAuth(), ctrl.GetCategories)
	Serve.GlobalRouter.SetRouterGet("/categories/detail", mid.UserAuth(), ctrl.GetCategoriesDetail)
	Serve.GlobalRouter.SetRouterGet("/category/:id", mid.UserAuth(), ctrl.GetSelectCategory)
	Serve.GlobalRouter.SetRouterGet("/category/detail/:id", mid.UserAuth(), ctrl.GetSelectCategoryDetail)
	// Serve.GlobalRouter.SetRouterGet("/category/delete", mid.UserAuth(), ctrl.DeleteCategory)
	// Serve.GlobalRouter.SetRouterPost("/category/upload", mid.UserAuth(), ctrl.UploadCategory)
	// Serve.GlobalRouter.SetRouterPost("/category/update", mid.UserAuth(), ctrl.UpdateCategory)

	// //=====================comment===========================/
	Serve.GlobalRouter.SetRouterGet("/comments/article/:id", mid.UserAuth(), ctrl.GetCommentsByArticle)
	Serve.GlobalRouter.SetRouterPost("/comments/create/change", mid.UserAuth(), ctrl.PublishComment)

	//=====================login=============================/
	Serve.GlobalRouter.SetRouterPost("/api/register", mid.CheckUserExistMiddleware, ctrl.Register)
	Serve.GlobalRouter.SetRouterPost("/api/login", ctrl.Login)
	Serve.GlobalRouter.SetRouterPost("/api/create_post", mid.UserAuth(), ctrl.CreatePost)
	Serve.GlobalRouter.SetRouterPost("/api/update_post", mid.UserAuth(), ctrl.UpdatePost)
	Serve.GlobalRouter.SetRouterGet("/users/currentUser", mid.UserAuth(), ctrl.GetCurrentUser)

	// //=====================tags===============================/
	Serve.GlobalRouter.SetRouterGet("/tags", mid.UserAuth(), ctrl.GetTags)
	Serve.GlobalRouter.SetRouterGet("/tags/detail", mid.UserAuth(), ctrl.GetTagsDetail)
	Serve.GlobalRouter.SetRouterGet("/tags/hot", mid.UserAuth(), ctrl.GetTagsHot)
	Serve.GlobalRouter.SetRouterGet("/tags/:id", mid.UserAuth(), ctrl.GetSelectTags)
	Serve.GlobalRouter.SetRouterGet("/tags/detail/:id", mid.UserAuth(), ctrl.GetSelectTagDetail)

	// //=====================upload=============================/
	// Serve.GlobalRouter.SetRouterPost("/upload", mid.UserAuth(), ctrl.Upload)

	return true
}

func Run(ipport string) bool {
	ser := &http.Server{
		Addr:    ipport,
		Handler: Serve.GlobalRouter.Engine,
	}

	err := ser.ListenAndServe()
	if err != nil {
		log.Debugf("Run serve fail:%s\n", err)
		return false
	}
	return true
}
