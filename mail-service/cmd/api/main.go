package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Config struct {
	Mailer Mail
}

const webPort = "80"

func main() {
	app := Config{}

	log.Println("Starting mail service on port", webPort)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func createMail() {
	m := Mail{
		Domain: os.Getenv("MAIL_DOMAIN"),
		Port:   587,
		From:   "
		To:     "
		Auth: smtp.PlainAuth(
	}
}
