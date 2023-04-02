package controllers

import (
	"go-simple-crud-1/initializers"
	"go-simple-crud-1/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreatePost(c *gin.Context) {
	var body struct {
		Title   string
		Content string
	}
	c.Bind(&body)

	if body.Title == "" || body.Content == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "badRequest",
			"message": "Invalid parameters",
			"target":  "title, content",
		})
		return
	}

	post := models.Post{Title: body.Title, Content: body.Content}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":       "badRequest",
			"message":    "Cannot create post",
			"target":     "post",
			"innererror": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func GetPosts(c *gin.Context) {
	// TODO: Need pagination
	var posts []models.Post
	result := initializers.DB.Find(&posts)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":       "badRequest",
			"message":    "Cannot find posts",
			"target":     "posts",
			"innererror": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})
}

func GetPost(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "badRequest",
			"message": "Invalid parameters",
			"target":  "id",
		})
		return
	}

	idInInteger, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "badRequest",
			"message": "Invalid parameters",
			"target":  "id",
		})
		return
	}

	var post models.Post
	result := initializers.DB.First(&post, idInInteger)

	if result.Error != nil || result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":       "badRequest",
			"message":    "Cannot find post with ID",
			"target":     "id",
			"innererror": result.Error,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func UpdatePost(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "badRequest",
			"message": "Invalid parameters",
			"target":  "id",
		})
		return
	}

	idInInteger, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "badRequest",
			"message": "Invalid parameters",
			"target":  "id",
		})
		return
	}

	var post models.Post
	result := initializers.DB.First(&post, idInInteger)

	if result.Error != nil || result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":       "badRequest",
			"message":    "Cannot find post with ID",
			"target":     "id",
			"innererror": result.Error,
		})
		return
	}

	var body struct {
		Title   string
		Content string
	}
	c.Bind(&body)

	initializers.DB.Model(&post).Updates(models.Post{Title: body.Title, Content: body.Content})

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func DeletePost(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "badRequest",
			"message": "Invalid parameters",
			"target":  "id",
		})
		return
	}

	idInInteger, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "badRequest",
			"message": "Invalid parameters",
			"target":  "id",
		})
		return
	}

	var post models.Post
	result := initializers.DB.First(&post, idInInteger)

	if result.Error != nil || result.Error == gorm.ErrRecordNotFound {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":       "badRequest",
			"message":    "Cannot find post with ID",
			"target":     "id",
			"innererror": result.Error,
		})
		return
	}

	initializers.DB.Delete(&post) // Auto do soft delete :D

	c.JSON(http.StatusOK, gin.H{
		"message": true,
	})
}
