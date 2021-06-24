package main

type Message struct {
	ID       string `json:"id"`
	Text     string `json:"text"`
	AuthorId string `json:"user"`
	Timestamp string `json:"timestamp"`
}
