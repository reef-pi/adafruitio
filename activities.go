package adafruitio

import (
	"github.com/google/go-querystring/query"
)

// Owner represent Owner of an activity as defined by Adafruit.io API
type Owner struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

// Activity represent activity response object as defined by Adafruit.io API
type Activity struct {
	UserName  string                 `json:"username"`
	Owner     Owner                  `json:"owner"`
	ID        int                    `json:"id"`
	Action    string                 `json:"action"`
	Model     string                 `json:"model"`
	Data      map[string]interface{} `json:"data"`
	UserID    int                    `json:"user_id"`
	CreatedAt string                 `json:"created_at"`
	UpdatedAt string                 `json:"updated_at"`
}

// ListActivitiesOptions represent list actvity api options as defined by Adafruit.io API
type ListActivitiesOptions struct {
	StartTime string `url:"start_time,omitempty"`
	EndTime   string `url:"end_time,omitempty"`
	Limit     int    `url:"limit,omitempty"`
}

// DeleteActivities deletes a given user's activity
func (c *Client) DeleteActivities(u string) error {
	resp, err := c.delete("/" + u + "/activities")
	if resp != nil {
		resp.Body.Close()
	}
	return err
}

// ListActivities list an user's activity
func (c *Client) ListActivities(u string, o ListActivitiesOptions) ([]Activity, error) {
	var activities []Activity
	v, err := query.Values(o)
	if err != nil {
		return activities, err
	}
	resp, err := c.get("/" + u + "/activities?" + v.Encode())
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return activities, err
	}
	return activities, c.decodeJSON(resp, &activities)
}

// ListActivitiesByType list an user's activity of a spcific type
func (c *Client) ListActivitiesByType(u, t string, o ListActivitiesOptions) ([]Activity, error) {
	var activities []Activity
	v, err := query.Values(o)
	if err != nil {
		return activities, err
	}
	resp, err := c.get("/" + u + "/activities/" + t + "?" + v.Encode())
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return activities, err
	}
	return activities, c.decodeJSON(resp, &activities)
}
