package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {}

type Config struct {
	R *gin.Engine
}

func NewHandler(c *Config){
	// h:=&Handler{}
	g:=c.R.Group("/account")
	g.GET("/",func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"message":"hello",
		})
	})
}