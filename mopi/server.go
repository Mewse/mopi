package mopi

import (
	"encoding/json"
	"log"
	"net/http"
)

type Server struct {
	Routes map[string]Response
}

type Response struct {
	Code int    `json:"code"`
	Body string `json:"body"`
	Url  string `json:"url"`
}

func NewServer() *Server {
	s := Server{}
	s.Routes = make(map[string]Response)
	return &s
}

func (s *Server) Register(w http.ResponseWriter, r *http.Request) {
	var resp Response
	err := json.NewDecoder(r.Body).Decode(&resp)

	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
	}
	s.Routes[resp.Url] = resp
	log.Print(s.Routes)
	w.WriteHeader(201)
}

func (s *Server) Endpoint(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path
	resp, ok := s.Routes[url]
	if ok {
		out, err := json.Marshal(resp.Body)
		if err != nil {
			log.Printf("Unable to unmarshall body for response %s", url)
			w.WriteHeader(500)
		} else {
			w.Write(out)
			w.WriteHeader(200)
		}

	} else {
		log.Printf("No response found for url %s", url)
		w.WriteHeader(404)
	}
}
