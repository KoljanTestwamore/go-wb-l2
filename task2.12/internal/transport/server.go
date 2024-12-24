package transport

import (
	"encoding/json"
	"log"
	"net/http"
	"server/internal/services"
	"time"
)

type HttpServer struct {
	handler services.Handler
}


type HandleCreateEventBody struct {
	Time time.Time `json:"time"`
	Id uint `json:"id"`
}

func (s *HttpServer) Run() {
	mux := http.NewServeMux()

	s.handler = services.NewCalendarHandler()
	// POST /create_event 

	// POST /update_event 
	
	// POST /delete_event 
	
	// GET /events_for_day 
	
	// GET /events_for_week 
	
	// GET /events_for_month

	mux.HandleFunc("/create_event", s.handleCreateEvent)
}

func Middleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("middleware", r.URL)
			h.ServeHTTP(w, r)
	})
}

func (s *HttpServer) handleCreateEvent(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method does not exits", http.StatusServiceUnavailable)
		return
	}

	var p HandleCreateEventBody

	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		log.Fatalf("Could not decode body %v", r.Body)

		http.Error(w, "Bad request body", http.StatusServiceUnavailable)
		return
	}

	eid, t := s.handler.HandleCreateEvent(p.Id, p.Time)

	w.Header().Set("Content-Type", "application/json")
	
	j, _ := json.Marshal(map[uint]time.Time {
		eid: t,
	})

	code, err := w.Write(j)

	if code != 200 {
		log.Fatal("Response request not 200")
	}
}