package main

import "github.com/gin-gonic/gin"

var router *gin.Engine

func main() {

	// Set the router as the default one provided by Gin
	router = gin.Default()

	// Process the templates at the start so that they don't have to be loaded
	// from the disk again. This makes serving HTML pages very fast.
	router.LoadHTMLGlob("templates/*")

	// Handle Index
	router.GET("/", showIndexPage)

	// Handle GET requests at /article/view/some_article_id
	router.GET("/article/view/:article_id", getArticle)

	// Start serving the application
	router.Run()

}
