package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func handleUpload(c *gin.Context) {
	// Retrieve the uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Save the file to a specified destination
	err = c.SaveUploadedFile(file, "uploads/"+file.Filename)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}

func handleUpload(c *gin.Context) {
	// Retrieve the uploaded file
	file, err := c.FormFile("file")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Check the file size
	if file.Size > 8<<20 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "File size exceeds the limit"})
		return
	}
	// Save the file to a specified destination
	err = c.SaveUploadedFile(file, "uploads/"+file.Filename)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
