package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"github.com/go-chi/chi"

	_ "github.com/heroku/x/hmetrics/onload"
)

func main(){
	listenAddr := os.Getenv("PORT")

	if listenAddr == "" {
		log.Fatal("$PORT must be set")
	}
	if err := http.ListenAndServe(listenAddr, Routes()); err != nil {
		fmt.Printf("F")
	}
}

func Routes() chi.Router {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintln(w, "cp-audio-lib")
	})
	return r
}
