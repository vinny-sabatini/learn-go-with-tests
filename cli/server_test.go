package poker

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("returns Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
		assertResponseCode(t, response.Code, http.StatusOK)
	})

	t.Run("returns Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
		assertResponseCode(t, response.Code, http.StatusOK)
	})

	t.Run("return 404 on missing player", func(t *testing.T) {
		request := newGetScoreRequest("Frank")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseCode(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("record a win when POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseCode(t, response.Code, http.StatusAccepted)

		AssertPlayerWin(t, &store, player)
	})
}

func TestLeague(t *testing.T) {
	t.Run("it returns 200 on /league", func(t *testing.T) {
		wantedLeague := []Player{
			{"Vinny", 32},
			{"Dom", 20},
			{"Bill", 19},
		}

		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := getLeagueFromResponse(t, response.Body)
		assertResponseCode(t, response.Code, http.StatusOK)
		assertLeague(t, got, wantedLeague)
		assertContentType(t, response, jsonContentType)
	})
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type %s, got %v", want, response.Result().Header)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("response body is wrong, got '%s', want '%s'", got, want)
	}
}

func assertResponseCode(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("response code is wrong, got '%d', want '%d'", got, want)
	}
}

func assertLeague(t *testing.T, got, want []Player) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()
	league, err := NewLeague(body)

	if err != nil {
		t.Fatalf("Unable to parse response from server '%s' into slice of Player, '%v'", body, err)
	}

	return
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}
