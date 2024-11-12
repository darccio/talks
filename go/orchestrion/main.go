package main

import (
	"fmt"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func serveHTTP() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Pong!")
	})
	http.ListenAndServe("127.0.0.1:9090", r)
}

func generateRequests() {
	for {
		_, err := http.Get("http://127.0.0.1:9090/ping")
		if err != nil {
			slog.Error(err.Error())
		}
		amount := rand.Int() % 300
		time.Sleep(time.Duration(amount) * time.Millisecond)
	}
}

func main() {
	go serveHTTP()
	generateRequests()
}
