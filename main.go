package main

import (
	"Test1/mapping"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	mapping.UrlMapping(router).Run(":8080")
}
