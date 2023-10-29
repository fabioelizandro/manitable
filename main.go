package main

import (
	"net/http"

	"github.com/fabioelizandro/manitable/modules/must"
	"github.com/fabioelizandro/manitable/templates"
	"github.com/gin-gonic/gin"
)

func main() {
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

	must.NoErr(router.Run("localhost:8080"))
}
