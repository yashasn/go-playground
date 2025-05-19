package server

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	Store PlayerStore
}

func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case http.MethodGet:
		p.getScore(w, r)
	case http.MethodPost:
		p.storeWin(w, r)
	}

}
func (p *PlayerServer) storeWin(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	p.Store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
	return

}

func (p *PlayerServer) getScore(w http.ResponseWriter, r *http.Request) {

	player := strings.TrimPrefix(r.URL.Path, "/players/")

	score := p.Store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}
func GetPlayerScore(name string) string {
	if name == "P1" {
		return "20"
	}

	if name == "P2" {
		return "10"
	}

	return ""
}
