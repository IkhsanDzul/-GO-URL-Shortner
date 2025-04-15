package main

func main() {
	r := gin.Default()

	svc := service.NewService()
	handler := handler.NewHandler(svc)

	r.GET("/shorten", handler.ShortenURL)
	r.GET("/:short", handler.ResolveURL)

	r.Run(":3000")
}