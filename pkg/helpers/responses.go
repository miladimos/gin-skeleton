package helpers

import (
	"gin/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusOK(c *gin.Context, message string, data map[string]interface{}) {
	c.JSON(http.StatusOK, structs.Response{
		Status:       true,
		Message:      "Success",
		OtherMessage: message,
		Data:         data,
	})
}
func StatusInternalServerError(c *gin.Context, message string, data map[string]interface{}) {
	c.JSON(http.StatusInternalServerError, structs.Response{
		Status:       false,
		Message:      "Error",
		OtherMessage: message,
		Data:         data,
	})
}
func StatusNotFound(c *gin.Context, message string, data map[string]interface{}) {
	c.JSON(http.StatusNotFound, structs.Response{
		Status:       false,
		Message:      "Error",
		OtherMessage: message,
		Data:         data,
	})
}
