package main

import (
	"net/http"
	"strconv"
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

func getFriendList(c *gin.Context) {
	var uid = c.Param("uid")


	for userIndex := range users {
		if users[userIndex].ID == uid {
			var friendList []User = []User{}
			for friend := range users[userIndex].FriendIdList {
				for user := range users {
					if users[user].ID == users[userIndex].FriendIdList[friend] {
						friendList = append(friendList, users[user])
					}
				}
			}

			c.JSON(http.StatusOK, friendList)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}

func createConversation(c *gin.Context) {
	var title = c.Param("title")
	var photoUrl = c.Param("photoUrl")

	var conversation Conversation
	conversation.ID = strconv.Itoa(len(conversations))
	conversation.Title = title
	conversation.PhotoUrl = photoUrl
	conversation.Messages = []Message{}

}
