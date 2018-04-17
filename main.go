package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Name string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/*.html")
	r.GET("/", func(c *gin.Context) {
		user := User{
			Name: "konojunya",
		}
		c.HTML(http.StatusOK, "index.html", user)
	})

	r.Run(":4000")
}
