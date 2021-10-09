package routers

import (
	_ "gin-api/docs"
	"gin-api/middleware/jwt"
	"gin-api/pkg/setting"
	"gin-api/pkg/upload"
	"gin-api/routers/api"
	v1 "gin-api/routers/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()

    r.Use(gin.Logger())

    r.Use(gin.Recovery())

    gin.SetMode(setting.ServerSetting.RunMode)

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/auth", api.GetAuth)

	r.POST("/upload", api.UploadImage)

    apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.GET("/tags", v1.GetTags) //获取标签列表
		apiv1.POST("/tags", v1.AddTag) //新建标签
		apiv1.PUT("/tags/:id", v1.UpdateTag) // 更新标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag) //删除标签

		//获取文章列表
        apiv1.GET("/articles", v1.GetArticles)
        //获取指定文章
        apiv1.GET("/articles/:id", v1.GetArticle)
        //新建文章
        apiv1.POST("/articles", v1.AddArticle)
        //更新指定文章
        apiv1.PUT("/articles/:id", v1.EditArticle)
        //删除指定文章
        apiv1.DELETE("/articles/:id", v1.DeleteArticle)	
	}

    return r
}