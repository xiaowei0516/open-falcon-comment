package main

import (
        "net/http"
	"github.com/gin-gonic/gin"
	yaag_gin "github.com/masato25/yaag/gin"
	"github.com/masato25/yaag/yaag"
)


func main() {
	routes := gin.Default()
	if true {
		yaag.Init(&yaag.Config{
			On:       true,
			DocTitle: "Gin",
			DocPath:  "apidoc.html",
			BaseUrls: map[string]string{"Production": "/app/v1", "Staging": "/app/v1"},
		})
		routes.Use(yaag_gin.Document())
	}
	routes.GET("/", func(c *gin.Context){
	   c.String(http.StatusOK, "hello")
	})
	routes.GET("/home", func(c *gin.Context){
	   c.String(http.StatusOK, "word")
	})
       routes.Run(":8888")
}
