package main

import (
	"github.com/fabioelizandro/manitable/modules/must"
	"github.com/fabioelizandro/manitable/web"
)

func main() {
	router := web.Router()
	must.NoErr(router.Run("localhost:8080"))
}
