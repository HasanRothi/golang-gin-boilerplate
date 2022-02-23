package routes

import (
	"github.com/gin-gonic/gin"
	. "golang-gin-boilerplate/bootstrap"
)

func Setup() {
	Router.GET("/", func(c *gin.Context) {
		Response.Json(c)
	})
}
