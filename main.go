package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("SuperScanner start.")

	r := gin.Default()


	r.Static("/static", "./web/static")
	r.StaticFile("/favicon.ico", "./web/static/images/favorite.ico")
	var htmlFiles []string
	_ = filepath.Walk("web/views", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".html") {
			htmlFiles = append(htmlFiles, path)
		}
		return nil
	})
	log.Println(htmlFiles)
	r.LoadHTMLFiles(htmlFiles...)



	datas := make(map[string]interface{})
	datas["name"] = "Asche"


	r.POST("/add", func(context *gin.Context) {
		fmt.Println(context.Keys)
		fmt.Println(context.Accepted)
		fmt.Println(context.PostForm("foo"))

		context.JSON(200, "done")
	})

	r.GET("/", func(context *gin.Context) {
		context.HTML(200, "home.html", datas)
	})
	r.GET("/timestamp", func(context *gin.Context) {
		context.HTML(200, "timestamp.html", nil)
	})

	fmt.Println("http://localhost:8080/")

	r.Run(fmt.Sprintf(":%d", ServerPort))
}
