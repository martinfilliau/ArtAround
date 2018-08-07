package main

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {
	r := gin.Default()
	ctx := context.Background()

	r.Use(CORSMiddleware())

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
