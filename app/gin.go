// gin.go: setup for the stand aside gin server to handle requests via the UI.
package main

import (
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
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status":  http.StatusUnprocessableEntity,
				"error":   "Unable to get file from request",
				"message": nil,
			})
			return
		}

		out, err := os.Create("./tmp/upload~" + header.Filename)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"error":   "Internal error - Unable to provision temporary file for write",
				"message": nil,
			})
			return
		}
		defer out.Close()

		_, err = io.Copy(out, file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"error":   "Internal error - Unable to write to temporary file",
				"message": nil,
			})
			return
		}

		isValid, status, msg := VerifyTradesHeaders(header)
		if !isValid {
			c.JSON(status, gin.H{
				"status":  http.StatusOK,
				"error":   msg,
				"message": nil,
			})
		} else {
			c.JSON(status, gin.H{
				"status":  status,
				"error":   nil,
				"message": msg,
			})
		}
	})

	return r
}
