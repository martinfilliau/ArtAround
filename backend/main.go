package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

func main() {
	r := gin.Default()
	ctx := context.Background()

	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
		panic(err)
	}

	r.GET("/ping", func(c *gin.Context) {
		info, code, err := client.Ping("http://127.0.0.1:9200").Do(ctx)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message":       "error",
				"elasticsearch": err,
			})
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"elasticsearch": gin.H{
				"code":    code,
				"version": info.Version.Number,
			},
		})
	})
	r.Run()
}
