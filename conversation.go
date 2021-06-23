package main

type Conversation struct {
	ID       string    `json:"id"`
	Title    string    `json:"title"`
	PhotoUrl string    `json:"photo_url"`
	Messages []Message `json:"messages"`
}