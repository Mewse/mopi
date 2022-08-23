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

func (s *Server) Register(w *http.ResponseWriter, r http.Request) {
	var body []byte
	_, err := r.Body.Read(body)
	if err != nil {
		log.Println(err)
	}
	var resp Response
	json.Unmarshal(body, resp)
	log.Println(resp)
	s.Routes[resp.Url] = resp
}
