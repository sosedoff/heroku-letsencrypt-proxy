package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Challenge struct {
	Token   string `json:"token"`
	AuthKey string `json:"auth_key"`
}

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
	challenge := Challenge{}

	if err := c.BindJSON(&challenge); err != nil {
		c.String(400, err.Error())
		return
	}

	if challenge.Token == "" || challenge.AuthKey == "" {
		c.String(400, "Invalid data")
		return
	}

	challenges[challenge.Token] = challenge.AuthKey
	c.String(200, "OK")
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5555"
	}

	router := gin.Default()
	router.GET("/:token", getToken)
	router.POST("/", setToken)
	router.Run(":" + port)
}
