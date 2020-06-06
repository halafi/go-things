package routes

import (
	"net/http"

	"github.com/halafi/go-things/go-anal/internal/handlers"
	"github.com/julienschmidt/httprouter"
)

// Routes initializes the routes
func Routes() http.Handler {
	rt := httprouter.New()

	eventHandler := handlers.NewEventHandler()
	rt.GET("/hello", eventHandler.Track)
	rt.GET("/stats", eventHandler.Stats)

	rt.ServeFiles("/dashboard/*filepath", http.Dir("./gogal/web"))

	return rt
}
