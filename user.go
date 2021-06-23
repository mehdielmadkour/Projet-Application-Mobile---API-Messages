package main

type User struct {
	ID         string `json:"id"`
	Username   string `json:"username"`
	PhotoUrl   string `json:"photo_url"`
	FriendIdList []string `json:"friends_id"`
	ConversationIdList []string `json:"conversations_id"`
}
