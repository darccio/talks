package main

import (
	"fmt"
	"log/slog"
	"math/rand/v2"
	"net/http"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

var log *slog.Logger

func serveHTTP() {
	r := chi.NewRouter()
	r.Use(logger)
	r.Use(middleware.Recoverer)
	r.Get("/ping", ping)
	http.ListenAndServe("127.0.0.1:9090", r)
}

func generateRequests() {
	for {
		_, err := http.Get("http://127.0.0.1:9090/ping")
		if err != nil {
			log.Error("failed to make request", slog.Attr{
				Key:   "error",
				Value: slog.StringValue(err.Error()),
			})
		}
		amount := rand.Int() % 300
		time.Sleep(time.Duration(amount) * time.Millisecond)
	}
}

func main() {
	go serveHTTP()
	generateRequests()
}

func logger(next http.Handler) http.Handler {
	log = slog.New(
		slog.NewJSONHandler(os.Stdout, nil),
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Info("request received")
		next.ServeHTTP(w, r)
	})
}

//dd:span span.name:ping
func ping(w http.ResponseWriter, r *http.Request) {
	if span, ok := tracer.SpanFromContext(r.Context()); ok {
		span.SetTag("my_route", r.URL.Path)
	}
	fmt.Fprint(w, "Pong!")
}
