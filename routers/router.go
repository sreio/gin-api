package routers

import (
	"gin-api/pkg/setting"
	v1 "gin-api/routers/api/v1"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

    r.Use(gin.Logger())

    r.Use(gin.Recovery())

    gin.SetMode(setting.RunMode)

    apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/tags", v1.GetTags) //获取标签列表
		apiv1.POST("/tags", v1.AddTag) //新建标签
		apiv1.PUT("/tags/:id", v1.UpdateTag) // 更新标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag) //删除标签
	}

    return r
}