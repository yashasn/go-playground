package server

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayersStore struct {
	scores map[string]int
}

func (s *StubPlayersStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayersStore{
		map[string]int{
			"P1": 20,
			"P2": 10,
		},
	}
	server := PlayerServer{&store}
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
