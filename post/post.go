package post

import (
	"net/http"

	"atomica-blog/models"

	"atomica-blog/db"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	posts := []models.Post{}
	dbInstance, _ := db.Open()
	dbInstance.Find(&posts)
	c.IndentedJSON(http.StatusOK, posts)
}

func Show(c *gin.Context) {
	id := c.Param("id")
	post := models.Post{}
	dbInstance, _ := db.Open()
	result := dbInstance.First(&post, id)

	if result.Error == nil {
		c.IndentedJSON(http.StatusOK, post)
		return
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Post not found"})
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	dbInstance, _ := db.Open()
	post := models.Post{}
	dbInstance.Delete(&post, id)

	c.IndentedJSON(http.StatusOK, gin.H{"message": "Success"})
}

func Store(c *gin.Context) {
	post := models.Post{}

	if err := c.BindJSON(&post); err != nil {
		return
	}
	dbInstance, _ := db.Open()
	dbInstance.Create(&post)
	c.IndentedJSON(http.StatusCreated, post)
}

func Update(c *gin.Context) {
	post := models.Post{}
	id := c.Param("id")
	dbInstance, _ := db.Open()
	result := dbInstance.First(&post, id)

	if result.Error == nil {
		if err := c.BindJSON(&post); err != nil {
			return
		}
	}

	dbInstance.Save(&post)
	c.IndentedJSON(http.StatusCreated, post)
}
