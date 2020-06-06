package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/halafi/go-things/go-anal/internal/repositories"
	"github.com/halafi/go-things/go-anal/internal/services"
	"github.com/julienschmidt/httprouter"
)

// EventHandler is handler for Events
type EventHandler struct {
	eventService *services.EventService
}

// NewEventHandler creates and returns new EventHandler
func NewEventHandler() *EventHandler {
	return &EventHandler{
		eventService: services.NewEventService(),
	}
}

// Track accepts analytics request and builds event from it
func (h *EventHandler) Track(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != http.MethodGet {
		http.Error(w, "Request method is not GET", http.StatusNotFound)
		return
	}
	event, err := h.eventService.BuildEvent(r)
	if err != nil {
		log.Println(err)
	}

	if event != nil && event.Valid() {
		h.eventService.LogEvent(event)
	}

	w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	w.Header().Set("Content-Type", "image/gif")
	w.Write(createPixel())
}

// Stats retrieves stats for the specified query
func (h *EventHandler) Stats(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	if r.Method != http.MethodGet {
		http.Error(w, "Request method is not GET", http.StatusNotFound)
		return
	}

	urlVals := r.URL.Query()
	query := urlVals.Get("q")

	stats := h.eventService.Stats(
		repositories.Stats(query),
	)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(stats)
}

// This is an image pixel, specifically GIF89a
// Check the references at the end for more info.
func createPixel() []byte {
	return []byte{
		71, 73, 70, 56, 57, 97, 1, 0, 1, 0, 128, 0, 0, 0, 0, 0,
		255, 255, 255, 33, 249, 4, 1, 0, 0, 0, 0, 44, 0, 0, 0, 0,
		1, 0, 1, 0, 0, 2, 1, 68, 0, 59,
	}
}
