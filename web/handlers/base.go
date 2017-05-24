package handlers

import "github.com/gin-gonic/gin"

func HealthcheckHandler(c *gin.Context) error {
	c.JSON(200, gin.H{
		"message": "Ok",
	})

	return nil
}
