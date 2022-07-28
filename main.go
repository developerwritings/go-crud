package main

import (
	"fmt"
	"todoapp/db"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	database, err := db.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}

	handler := db.NewDBHandler(database)
	fmt.Println(handler)

	router.Run()

}
