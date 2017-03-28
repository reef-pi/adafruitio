package adafruitio

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type Handler struct {
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whats up")
	fmt.Println(r.Method)
	fmt.Println(r.URL.Path)
	fmt.Println(r.URL.RawQuery)
	for k, v := range r.Header {
		fmt.Println("Header:", k, "=", v)
	}
}

func TestDeleteActivity(t *testing.T) {
	s := httptest.NewServer(&Handler{})
	defer s.Close()

	c := &Client{
		authToken:   "fakeToken",
		apiEndpoint: s.URL,
	}

	if err := c.DeleteActivities("ac1"); err != nil {
		t.Error(err)
	}
	fmt.Println(s.URL)

}
