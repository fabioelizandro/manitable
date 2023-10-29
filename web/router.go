package web

import (
	"net/http"

	"github.com/fabioelizandro/manitable/web/templates"
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*.gohtml")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.gohtml", templates.Index{
			Layout: templates.Layout{
				Title: "My little title",
			},
			PageContent: "Here's the page content",
		})
	})

	return router
}
