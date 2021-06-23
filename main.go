package main

import "github.com/gin-gonic/gin"


func main() {
	r := gin.Default()

	r.GET("/users/", getUsers)

	r.Run(":80")
}
