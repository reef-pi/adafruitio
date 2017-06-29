package adafruitio

type User struct {
	ID        int    `json:"id"`
	Name      string `json:"username"`
	Color     string `json:"color"`
	TimeZone  string `json:"time_zone"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (c *Client) GetUser() (User, error) {
	var u User
	resp, err := c.get("/user")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return u, err
	}
	return u, c.decodeJSON(resp, &u)
}
