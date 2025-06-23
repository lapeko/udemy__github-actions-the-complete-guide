package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetHealthController(c *gin.Context) {
	c.String(http.StatusOK, "OK")
}
