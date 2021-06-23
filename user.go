package main

type User struct {
	ID                 string   `json:"uid"`
	Username           string   `json:"username"`
	PhotoUrl           string   `json:"photoUrl"`
	FriendIdList       []string `json:"friendList"`
	ConversationIdList []string `json:"conversationList"`
}
