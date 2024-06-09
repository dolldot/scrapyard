package handlers

import (
	"net/http"

	"github.com/dolldot/scrapyard/bookmarkd/views/pages"
	"github.com/gin-gonic/gin"
)

func HealthCheck(c *gin.Context) {
	c.Status(http.StatusOK)
}

func Landing(c *gin.Context) {
	render(c, http.StatusOK, pages.Landing())
}
