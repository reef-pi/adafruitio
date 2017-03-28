package adafruitio

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestDeleteActivities(t *testing.T) {
	var method, path string
	h := func(w http.ResponseWriter, r *http.Request) {
		method = r.Method
		path = r.RequestURI
	}
	s := httptest.NewServer(http.HandlerFunc(h))
	defer s.Close()

	c := &Client{
		authToken:   "fakeToken",
		apiEndpoint: s.URL,
	}

	if err := c.DeleteActivities("ac1"); err != nil {
		t.Error(err)
	}
	if method != "DELETE" {
		t.Errorf("Expected: DELETE, Found:%s", method)
	}
	if path != "/ac1/activities" {
		t.Errorf("Expected: /ac1/activities  Found:%s", path)
	}
}

func TestListActivities(t *testing.T) {
	m := newMockServer("activities.json")
	defer m.Close()
	c := &Client{
		authToken:   "faketoken",
		apiEndpoint: m.URL(),
	}
	opt := ListActivitiesOptions{
		Limit: 100,
	}
	activities, err := c.ListActivities("example_user", opt)
	if err != nil {
		t.Error(err)
	}
	for _, activity := range activities {
		fmt.Println(activity.Model, activity.Action, activity.CreatedAt)
	}
}
