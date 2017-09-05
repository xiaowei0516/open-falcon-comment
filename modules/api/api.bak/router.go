package main

import (
        "net/http"
	"github.com/gin-gonic/gin"
	"fmt"
	"os"
	"io"
)


func main() {
	router := gin.Default()
        router.POST("/upload", func(c *gin.Context) {
        name := c.PostForm("name")
        fmt.Println(name)
        file, header, err := c.Request.FormFile("upload")
        if err != nil {
            c.String(http.StatusBadRequest, "Bad request")
            return
        }
        filename := header.Filename

        fmt.Println(file, err, filename)

        out, err := os.Create(filename)
        defer out.Close()
        io.Copy(out, file)
        c.String(http.StatusCreated, "upload successful")
    })
    router.Run(":8888")
}
