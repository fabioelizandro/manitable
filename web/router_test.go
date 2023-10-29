package web_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/fabioelizandro/manitable/modules/must"
	"github.com/fabioelizandro/manitable/web"
	"github.com/stretchr/testify/assert"
)

func TestRouterSuite(t *testing.T) {
	t.Run("home page", func(t *testing.T) {
		router := web.Router()

		res := httptest.NewRecorder()
		router.ServeHTTP(res, must.Return(http.NewRequest("GET", "/", nil)))
		doc := must.Return(goquery.NewDocumentFromReader(res.Body))

		assert.Equal(t, 200, res.Code)
		assert.Equal(t, "My little title", doc.Find("title").Text())
		assert.Contains(t, doc.Find(".page-content").Text(), "Here's the page content")
	})
}
