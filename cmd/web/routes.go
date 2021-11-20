package main

import (
	"github.com/bmizerany/pat"
	"github.com/igor-koniukhov/web_s_study/internal/handlers"
	"net/http"
)

// routes defines the application routes
func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	return mux
}
