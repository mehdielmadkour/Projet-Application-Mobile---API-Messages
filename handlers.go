package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var users []User = []User{}
var conversations []Conversation = []Conversation{}

func getUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func createUser(c *gin.Context) {
	var uid = c.Param("uid")
	var username = c.Param("username")
	var photoUrl = c.Param("photoUrl")

	for user := range users {
		if users[user].ID == uid {
			return
		}
	}

	var user = User{uid, username, photoUrl, []string{}, []string{}}

	users = append(users, user)

}

func addFriend(c *gin.Context) {
	var uid = c.Param("uid")
	var friendId = c.Param("friendId")

	for user := range users {
		if users[user].ID == uid {
			users[user].FriendIdList = append(users[user].FriendIdList, friendId)
			return
		}
	}
}
