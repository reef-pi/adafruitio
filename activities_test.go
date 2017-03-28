package adafruitio

import (
	"testing"
)

func TestDeleteActivities(t *testing.T) {
	m := newMockServer("")
	defer m.Close()

	c := &Client{
		authToken:   "fakeToken",
		apiEndpoint: m.URL(),
	}

	if err := c.DeleteActivities("ac1"); err != nil {
		t.Error(err)
	}
	if m.Method != "DELETE" {
		t.Errorf("Expected: DELETE, Found:%s", m.Method)
	}
	if m.RequestURI != "/ac1/activities" {
		t.Errorf("Expected: /ac1/activities  Found:%s", m.RequestURI)
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
		return
	}
	if len(activities) == 0 {
		t.Errorf("Expected non-zero activities")
	}
	if m.Method != "GET" {
		t.Errorf("Expected GET. Found:%s", m.Method)
	}
	if m.RequestURI != "/example_user/activities?limit=100" {
		t.Errorf("Expected /example_user/activities?limit=100. Found:%s", m.RequestURI)
	}
}

func TestListActivitiesByType(t *testing.T) {
	m := newMockServer("activities.json")
	defer m.Close()
	c := &Client{
		authToken:   "faketoken",
		apiEndpoint: m.URL(),
	}
	opt := ListActivitiesOptions{
		Limit: 100,
	}
	activities, err := c.ListActivitiesByType("example_user", "xtype", opt)
	if err != nil {
		t.Error(err)
		return
	}
	if len(activities) == 0 {
		t.Errorf("Expected non-zero activities")
	}
	if m.Method != "GET" {
		t.Errorf("Expected GET. Found:%s", m.Method)
	}
	if m.RequestURI != "/example_user/activities/xtype?limit=100" {
		t.Errorf("Expected /example_user/activities/xtype?limit=100. Found:%s", m.RequestURI)
	}

}
