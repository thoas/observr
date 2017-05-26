package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Healthcheck(c *gin.Context) error {
	c.JSON(http.StatusOK, gin.H{
		"message": "Ok",
	})

	return nil
}
