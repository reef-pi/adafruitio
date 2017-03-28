package adafruitio

import (
	"github.com/google/go-querystring/query"
)

type Owner struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

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

type ListActivitiesOptions struct {
	StartTime string `url:"start_time,omitempty"`
	EndTime   string `url:"end_time,omitempty"`
	Limit     int    `url:"limit,omitempty"`
}

func (c *Client) DeleteActivities(u string) error {
	_, err := c.delete("/" + u + "/activities")
	return err
}

func (c *Client) ListActivities(u string, o ListActivitiesOptions) ([]Activity, error) {
	var activities []Activity
	v, err := query.Values(o)
	if err != nil {
		return activities, err
	}
	resp, err := c.get("/" + u + "/activities?" + v.Encode())
	if err != nil {
		return activities, err
	}
	return activities, c.decodeJSON(resp, &activities)
}

func (c *Client) ListActivitiesByType(u, t string, o ListActivitiesOptions) ([]Activity, error) {
	var activities []Activity
	v, err := query.Values(o)
	if err != nil {
		return activities, err
	}
	resp, err := c.get("/" + u + "/activities/" + t + "?" + v.Encode())
	if err != nil {
		return activities, err
	}
	return activities, c.decodeJSON(resp, &activities)
}
