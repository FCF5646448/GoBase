package routers

import (
	"github.com/fcf/go-gin-example/pkg/setting"
	v1 "github.com/fcf/go-gin-example/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api/v1")
	{
		// 获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		// 新建标签
		apiv1.POST("/tags", v1.AddTag)
		// 更新
		apiv1.PUT("/tags/:id", v1.EditTag)
		// 删除
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
	}

	// r.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "test",
	// 	})
	// })

	return r
}
