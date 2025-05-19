package server

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayersStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayersStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}
func (s *StubPlayersStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayersStore{
		map[string]int{
			"P1": 20,
			"P2": 10,
		},
		nil,
	}
	server := NewPlayerServer(&store)
	t.Run("return P1 score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/P1", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "20")

	})
	t.Run("return P2 score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/P2", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatusCode(t, response.Code, http.StatusOK)
		assertResponseBody(t, response.Body.String(), "10")

	})
	t.Run("return 404 on missing player", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/P3", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatusCode(t, response.Code, http.StatusNotFound)

	})
}

func TestPOSTScore(t *testing.T) {
	store := StubPlayersStore{
		make(map[string]int),
		make([]string, 0),
	}
	server := NewPlayerServer(&store)

	t.Run("store P1 score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/players/P1", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatusCode(t, response.Code, http.StatusAccepted)
		if len(store.winCalls) != 1 {
			t.Errorf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
		}
	})
}

func TestLeague(t *testing.T) {
	store := StubPlayersStore{}
	server := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got []Player

		err := json.NewDecoder(response.Body).Decode(&got)

		if err != nil {
			t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", response.Body, err)
		}

		assertStatusCode(t, response.Code, http.StatusOK)
	})
}

func assertResponseBody(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
func assertStatusCode(t testing.TB, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
