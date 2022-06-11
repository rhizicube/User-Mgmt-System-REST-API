package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type users struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	City        string `json:"city"`
	PhoneNumber int64  `json:"phnumber"`
}

//getUserInfo will provide all the users details - GET
func getUserInfo(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, userData)
}

// search the userDetail by ID - GET
func getUserInfoById(c *gin.Context) {
	id := c.Param("id")

	for _, a := range userData {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No user found with given id"})
}

// add new user in userData - POST
func addNewUser(c *gin.Context) {
	var newUser users

	// if body is empty
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "message": "Please Send some data"})
		return
	}
	//body contains - {}
	if newUser.isEmpty() {
		c.JSON(http.StatusNoContent, gin.H{"error": "204 No Content success", "message": "No data inside json"})
		return
	}

	// generate unique id  and convert that id into string
	// after that newUser into userData
	if newUser.ID == "" {
		rand.Seed(time.Now().UnixNano())
		newUser.ID = strconv.Itoa(rand.Intn(100))
	}

	userData = append(userData, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

func (u *users) isEmpty() bool {
	return u.Name == ""
}

//update any user details - PUT
func updateUserInfo(c *gin.Context) {
	id := c.Param("id")
	var newUserDetails users

	if err := c.ShouldBindJSON(&newUserDetails); err != nil {
		c.JSON(422, gin.H{"error": err.Error(), "message": "invalid request Body"})
		return
	}

	for index, a := range userData {
		if a.ID == id {
			userData[index].Name = newUserDetails.Name
			userData[index].City = newUserDetails.City
			userData[index].PhoneNumber = newUserDetails.PhoneNumber
			c.IndentedJSON(http.StatusCreated, newUserDetails)
			return
		}
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"error":   true,
		"message": "invalid user id",
	})

}

//remove a user from userData by ID - DELETE
func deleteUserInfo(c *gin.Context) {
	id := c.Param("id")

	for i, a := range userData {
		if a.ID == id {
			userData = append(userData[:i], userData[i+1:]...)
			c.JSON(200, gin.H{"error": false})
			return
		}
	}
	c.JSON(404, gin.H{
		"error":   true,
		"message": "invalid user id",
	})
}

var userData = []users{
	{ID: "1", Name: "Utkarsh", City: "Hyderabad", PhoneNumber: 9998887776},
	{ID: "2", Name: "Udit", City: "Noida", PhoneNumber: 989898987},
	{ID: "3", Name: "Lalit", City: "Bangalore", PhoneNumber: 9998887772},
}

func main() {
	fmt.Println("web-service-gin")
	router := gin.Default()
	router.GET("/users", getUserInfo)
	router.GET("/user/:id", getUserInfoById)
	router.POST("/user", addNewUser)
	router.PUT("/user/:id", updateUserInfo)
	router.DELETE("/user/:id", deleteUserInfo)
	router.Run("localhost:8000")
}
