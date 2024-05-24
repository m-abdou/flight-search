package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (r *Router) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]string{"Status": "Live"})
}
