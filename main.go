package main

import (
	"fmt"
	"weather/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("file can't found", err)
		return
	}
	router.Routing()
}
