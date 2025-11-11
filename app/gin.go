// gin.go: setup for the stand aside gin server to handle requests via the UI.
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewGin() *gin.Engine {
	r := gin.Default()

	r.Use(cors.Default())

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"response": "pong"})
	})

	r.POST("/upload/trades-csv", func(c *gin.Context) {
		file, header, err := c.Request.FormFile("data")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read file"})
			return
		}

		out, err := os.Create("./tmp/upload~" + header.Filename)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Internal error writing to disk"})
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Internal error writing to disk"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Succesfully uploaded csv file to path: %s", "./tmp/upload~"+header.Filename),
		})
	})

	return r
}
