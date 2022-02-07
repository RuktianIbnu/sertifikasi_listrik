package main

import (
	"epiket-api/http/router"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func init() {
	if strings.ToLower(os.Getenv("GIN_MODE")) != "release" {
		if err := godotenv.Load("configs/.env"); err != nil {
			log.Fatalln(err)
		}
	}
}

func main() {
	if err := router.Routes().Run(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT"))); err != nil {
		log.Fatalln(err)
	}
}
