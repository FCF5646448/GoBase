package main

import (
	"fmt"
	"net/http"

	"github.com/fcf/go-gin-example/pkg/setting"
	"github.com/fcf/go-gin-example/routers"
)

func main() {
	router := routers.InitRouter()
	// router := gin.Default()
	// router.GET("/test", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "test",
	// 	})
	// })

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
