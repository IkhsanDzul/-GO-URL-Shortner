package main

import (
	"url-shortner/shortner"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	svc := shortner.NewService()
	handler := shortner.NewHandler(svc)

	r.GET("/shorten", handler.ShortenURL)
	r.GET("/:short", handler.ResolveURL)

	r.Run(":3000")
}