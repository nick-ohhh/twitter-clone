package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() //r for router
	//r.SetTrustedProxies([]string{""}) //add an ip for squid proxy
	r.GET("/example", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "helloWorld",
		})
	})
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/tacos", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "areGood",
		})
	})
	r.Run()
}
