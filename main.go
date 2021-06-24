package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	//Users
	r.GET("/users/", getUsers)
	r.POST("/createUser/:uid/:username/:photoUrl", createUser)
	r.GET("/user/:uid", getUser)

	//Friends
	r.POST("/addFriend/:uid/:friendId", addFriend)
	r.GET("/friendList/:uid", getFriendList)

	//Conversations
	r.POST("/createConversation/:uid/:title/:photoUrl", createConversation)
	r.GET("/conversation/:id", getConversation)
	r.GET("/conversationList/:uid", getConversationList)
	r.POST("/postMessage/:conversationId/:authorId/:text", postMessage)
	r.POST("/inviteFriend/:conversationId/:friendId", inviteFriend)

	r.Run(":80")
}
