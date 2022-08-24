package mopi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Server struct {
	Routes map[string]Response
}

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
	Url  string      `json:"url"`
}

type ConfigContents struct {
}

func NewServer() *Server {
	s := Server{}
	s.Routes = make(map[string]Response)
	s.LoadStaticConfigurations()
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
		}

	} else {
		log.Printf("No response found for url %s", url)
		w.WriteHeader(404)
	}
}

func (s *Server) LoadStaticConfigurations() {
	files, err := ioutil.ReadDir("configurations/")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if strings.Contains(file.Name(), ".json") {
			fileContents, fileErr := os.Open(fmt.Sprintf("configurations/%s", file.Name()))
			if fileErr != nil {
				log.Printf("File reading error: %s", fileErr)
			}
			var contents []Response
			jsonErr := json.NewDecoder(fileContents).Decode(&contents)
			if jsonErr != nil {
				log.Printf("Json decoding error: %s", jsonErr)
				break
			}
			for _, route := range contents {
				s.Routes[route.Url] = route
				log.Printf("Adding response for %s\n", route.Url)
			}
		}
	}
}
