package main

import (
	"fmt"
	"github.com/ainaraskaz/twitch/internal/auth"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := auth.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		RedirectURI:  "http://localhost:6969/redirect",
	}

	twitch := auth.NewTwitchAuth(cfg)

	// Gorilla session store (cookie-based)
	store := sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
	handler := auth.NewHandler(twitch, store)

	r := mux.NewRouter()
	r.HandleFunc("/login", handler.Login)
	r.HandleFunc("/redirect", handler.Callback)
	r.HandleFunc("/profile", handler.Profile)

	fmt.Println("Server running at http://localhost:6969")
	http.ListenAndServe(":6969", r)
}
