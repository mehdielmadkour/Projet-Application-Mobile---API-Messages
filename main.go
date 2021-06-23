package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/users/", getUsers)

	r.POST("/createUser/:uid/:username/:photoUrl", createUser)
	r.POST("/addFriend/:uid/:friendId", addFriend)

	r.Run(":80")
}
