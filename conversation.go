package main

type Conversation struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	PhotoUrl string    `json:"photoUrl"`
	Messages []Message `json:"messages"`
}
