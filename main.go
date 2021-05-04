package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"path/filepath"
	"strconv"
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



	data := make(map[string]interface{})
	data["name"] = "Asche"

	r.GET("/", func(context *gin.Context) {
		context.HTML(200, "home.html", data)
	})
	r.GET("/loading", func(context *gin.Context) {
		context.HTML(200, "loading.html", nil)
	})
	r.GET("/result", ResultController)
	r.GET("/ipResult", IPResultController)
	r.GET("/portResult", PortResultController)
	r.GET("/domainResult", DomainResultController)

	r.GET("/disSet", DisSetController)
	r.GET("/getAll", GetAllSlaveController)
	r.GET("/delete", DeleteSalveController)
	r.POST("/add", AddSlaveController)


	fmt.Println("http://localhost:" + strconv.Itoa(ServerPort))

	r.Run(fmt.Sprintf(":%d", ServerPort))
}
