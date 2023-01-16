package main

import (
	"fmt"
	"github.com/Moonlight-Zhao/go-project-example/cotroller"
	"github.com/Moonlight-Zhao/go-project-example/repository"
	"gopkg.in/gin-gonic/gin.v1"
	"net/http"
	"os"
)

func main() {
	if err := Init("./data/"); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := cotroller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	r.POST("/community/page/post/", func(context *gin.Context) {
		var item repository.PostQuery

		// Call BindJSON to bind the received JSON to new Post item
		if err := context.BindJSON(&item); err != nil {
			return
		}
		fmt.Println(item)
		//// Add the new album to the slice.
		//albums = append(albums, newAlbum)
		data := cotroller.InsertPost(&item)
		context.IndentedJSON(http.StatusCreated, data)
		//context.JSON(200, data)
	})
	err := r.Run()
	if err != nil {
		return
	}
}

func Init(filePath string) error {
	if err := repository.Init(filePath); err != nil {
		return err
	}
	return nil
}
