package adafruitio

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
)

type mockServer struct {
	s          *httptest.Server
	file       string
	Method     string
	RequestURI string
}

func (m *mockServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	m.Method = r.Method
	m.RequestURI = r.RequestURI
	if m.file == "" {
		return
	}
	wd, _ := os.Getwd()
	fi, err := os.Open(wd + "/testdata/" + m.file)
	if err != nil {
		return
	}
	io.Copy(w, fi)
}

func (m *mockServer) URL() string {
	return m.s.URL
}

func newMockServer(f string) *mockServer {
	m := &mockServer{
		file: f,
	}
	m.s = httptest.NewServer(m)
	return m
}
func (m *mockServer) Close() {
	m.s.Close()
	return
}
