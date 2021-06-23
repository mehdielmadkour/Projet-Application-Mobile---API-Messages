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
