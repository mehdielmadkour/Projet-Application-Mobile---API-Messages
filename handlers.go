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
func getUser(c *gin.Context) {
	var uid = c.Param("uid")

	for user := range users {
		if users[user].ID == uid {
			c.JSON(http.StatusOK, users[user])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
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
	var uid = c.Param("uid")
	var title = c.Param("title")
	var photoUrl = c.Param("photoUrl")

	var conversation = Conversation{strconv.Itoa(len(conversations)), title, photoUrl, []Message{}}
	conversations = append(conversations, conversation)

	for user := range users {
		if users[user].ID == uid {
			users[user].ConversationIdList = append(users[user].ConversationIdList, conversation.ID)
			break
		}
	}

	c.JSON(http.StatusOK, conversation.ID)
}

func getConversation(c *gin.Context) {
	var id = c.Param("id")

	for conversation := range conversations {
		if conversations[conversation].ID == id {
			c.JSON(http.StatusOK, conversations[conversation])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Conversation not found"})
}

func getConversationList(c *gin.Context) {
	var uid = c.Param("uid")

	for userIndex := range users {
		if users[userIndex].ID == uid {
			var conversationList []User = []User{}
			for conversation := range users[userIndex].ConversationIdList {
				for user := range users {
					if users[user].ID == users[userIndex].ConversationIdList[conversation] {
						conversationList = append(conversationList, users[user])
					}
				}
			}
	
			c.JSON(http.StatusOK, conversationList)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})

}

func postMessage(c *gin.Context) {
	var text = c.Param("text")
	var authorId = c.Param("authorId")
	var conversationId = c.Param("conversationId")

	var conversationIndex = -1

	for conversation := range conversations {
		if conversations[conversation].ID == conversationId {
			conversationIndex = conversation
			break
		}
	}

	if conversationIndex != -1 {
		var message Message
		message.ID = strconv.Itoa(len(conversations[conversationIndex].Messages))
		message.Text = text
		message.AuthorId = authorId

		conversations[conversationIndex].Messages =
			append(conversations[conversationIndex].Messages, message)
	}

}

func inviteFriend(c *gin.Context) {
	var uid = c.Param("uid")
	var conversationId = c.Param("conversationId")

	for user := range users {
		if users[user].ID == uid {
			users[user].ConversationIdList = append(users[user].ConversationIdList, conversationId)
			break
		}
	}
}
