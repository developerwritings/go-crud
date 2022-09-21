package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type User struct {
	username string
	password string
}

func main() {
	fmt.Println("EMI Calculator")
	router := gin.Default()
	userGroup := router.Group("/v1/user")
	userGroup.POST("/login", userLogin)
	userGroup.POST("/signup", userRegister)
	router.Run()
}

func userLogin(c *gin.Context) {
	var user User
	err := c.BindJSON(user)
	if err != nil {
		fmt.Println("unable to marshal Data")
	}
	fmt.Println(user)

}

func userRegister(c *gin.Context) {
	var user User
	err := c.BindJSON(user)
	if err != nil {
		fmt.Println("unable to marshal Data")
	}
	fmt.Println(user)
}
