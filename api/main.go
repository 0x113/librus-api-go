package main

import (
	"net/http"
	"os"

	"github.com/0x113/librus-api-go/api/handler"
	"github.com/0x113/librus-api-go/api/librusclient"

	log "github.com/sirupsen/logrus"
)

func main() {
	// check env variables
	port := os.Getenv("port")
	username := librusclient.LibrusClient.Username
	password := librusclient.LibrusClient.Password
	if port == "" || username == "" || password == "" {
		log.Infof("port: %s", port)
		log.Infof("username: %s", username)
		log.Infof("password: %s", password)
		log.Fatalf("Some environment variables are not set")
	}
	// create session
	if err := librusclient.LibrusClient.CreateSession(); err != nil {
		log.Fatalf("Unable to create session: %v", err)
	}

	// endpoints
	http.HandleFunc("/api/luckynumber", handler.LuckyNumber)

	// server
	log.Infof("Serving http on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Unable to start http server: %v", err)
	}
}
