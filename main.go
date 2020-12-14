package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/tinypulse-anh/simple-login-go/models"
	"github.com/tinypulse-anh/simple-login-go/routes"
)

func main() {
	router := Router.Init()
	models.Init()
	router.Run()
}
