package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type mockUser struct {
	Id    int    `json:"id"`
	Email string `json:"email" binding:"required"`
}

var mockedUsers = []mockUser{{Id: 1}, {Id: 2}, {Id: 3}}

// TODO Mock implementation
func PostUserController(c *gin.Context) {
	var user mockUser
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": err.Error()})
		return
	}
	user.Id = len(mockedUsers) + 1
	c.JSON(http.StatusCreated, gin.H{"data": user, "error": nil})
}

// TODO Mock implementation
func GetUsersController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": mockedUsers, "error": nil})
}

// TODO Mock implementation
func GetUserByIdController(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": fmt.Sprintf("id must be an integer. Provided %s", paramId)})
		return
	}
	if id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": fmt.Sprintf("id must be a positive integer. Provided %d", id)})
		return
	}
	if id > len(mockedUsers) {
		c.JSON(http.StatusNotFound, gin.H{"data": nil, "error": fmt.Sprintf("user with id: %d not found", id)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": mockedUsers[id-1], "error": nil})
}

// TODO Mock implementation
func PutUserController(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": fmt.Sprintf("id must be an integer. Provided %s", paramId)})
		return
	}
	if id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": fmt.Sprintf("id must be a positive integer. Provided %d", id)})
		return
	}
	if id > len(mockedUsers) {
		c.JSON(http.StatusNotFound, gin.H{"data": nil, "error": fmt.Sprintf("user with id: %d not found", id)})
		return
	}
	var user mockUser
	if err = c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": mockedUsers[id-1], "error": nil})
}

// TODO Mock implementation
func DeleteUserController(c *gin.Context) {
	paramId := c.Param("id")
	id, err := strconv.Atoi(paramId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": fmt.Sprintf("id must be an integer. Provided %s", paramId)})
		return
	}
	if id < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"data": nil, "error": fmt.Sprintf("id must be a positive integer. Provided %d", id)})
		return
	}
	if id > len(mockedUsers) {
		c.JSON(http.StatusNotFound, gin.H{"data": nil, "error": fmt.Sprintf("user with id: %d not found", id)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": mockedUsers[id-1], "error": nil})
}
