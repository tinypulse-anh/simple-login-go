package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/tinypulse-anh/simple-login-go/models"
)

func main() {
	models.Init()
}
