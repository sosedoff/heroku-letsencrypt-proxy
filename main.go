package main

import (
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
)

var challenges = map[string]string{}

func getToken(c *gin.Context) {
	value := challenges[c.Param("token")]
	if value == "" {
		c.String(404, "Token does not exist")
		return
	}
	c.String(200, value)
}

func setToken(c *gin.Context) {
	token := c.Param("token")
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}
	challenges[token] = string(data)
	c.String(200, "OK")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5555"
	}

	router := gin.Default()
	router.GET("/:token", getToken)
	router.POST("/:token", setToken)
	router.Run(":" + port)
}
