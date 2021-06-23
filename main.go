package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	//Users
	r.GET("/users/", getUsers)
	r.POST("/createUser/:uid/:username/:photoUrl", createUser)

	//Friends
	r.POST("/addFriend/:uid/:friendId", addFriend)
	r.GET("/friendList/:uid", getFriendList)

	//Conversations
	r.POST("/createConversation/:title/:photoUrl", createConversation)

	r.Run(":80")
}
