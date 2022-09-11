package routes

import (
	post "atomica-blog/post"

	"github.com/gin-gonic/gin"
)

func Start() {
	router := gin.Default()
	router.GET("/posts", post.Index)
	router.GET("/post/:id", post.Show)
	router.POST("/post", post.Store)
	router.DELETE("/post/:id", post.Delete)
	router.PATCH("/post/:id", post.Update)

	router.Run("localhost:8080")
}
