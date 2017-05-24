package handlers

import "github.com/gin-gonic/gin"

func UserCreateHandler(c *gin.Context) error {
	c.JSON(200, gin.H{
		"message": "Ok",
	})

	return nil
}
