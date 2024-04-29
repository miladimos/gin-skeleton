package controllers

import (
	"gin-skeleton/config"
	"gin-skeleton/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func TagList(c *gin.Context) {

	pageStr := c.Param("page")
	page, _ := strconv.Atoi(pageStr)
	offset := (page - 1) * 10

	var tags []models.Tag
	config.DB.Limit(10).Offset(offset).Find(&tags)

	c.JSON(200, gin.H{"tags": tags})
}

func TagCreate(c *gin.Context) {

	var body struct {
		Title string
	}

	c.Bind(&body)

	tag := models.Tag{Title: body.Title}

	result := config.DB.Create(&tag)
	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{"message": "tag created"})
}

func TagShow(c *gin.Context) {

	id := c.Param("id")

	var tag models.Tag
	config.DB.First(&tag, id)

	c.JSON(200, gin.H{"tag": tag})
}

func TagUpdate(c *gin.Context) {

	var body struct {
		Title string
	}

	c.Bind(&body)

	id := c.Param("id")

	var tag models.Tag
	config.DB.First(&tag, id)

	config.DB.Model(&tag).Updates(models.Tag{
		Title: body.Title,
	})

	c.JSON(200, gin.H{"tag": "tag updated"})
}

func TagDelete(c *gin.Context) {
	id := c.Param("id")

	config.DB.Delete(&models.Tag{}, id)

	c.JSON(200, gin.H{"tag": "tag deleted"})

}
