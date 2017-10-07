package adafruitio

type FeedData struct {
	First map[string]string `json:"first"`
	Last  map[string]string `json:"last"`
	Count int               `json:"count"`
}

type FeedDetails struct {
	SharedWith []map[string]string `json:"shared_with"`
	Data       FeedData            `json:"data"`
}

type Feed struct {
	ID            int                    `json:"id,omitempty"`
	Name          string                 `json:"name"`
	Key           string                 `json:"key,omitempty"`
	Group         map[string]interface{} `json:"group,omitempty"`
	Groups        []Group                `json:"groups,omitempty"`
	Description   string                 `json:"description,omitempty"`
	Details       FeedDetails            `json:"details,omitempty"`
	UnitType      string                 `json:"unit_type,omitempty"`
	UnitSymbol    string                 `json:"unit_symbol,omitempty"`
	History       bool                   `json:"history,omitempty"`
	Visibility    string                 `json:"visibility,omitempty"`
	License       string                 `json:"license,omitempty"`
	Enabled       bool                   `json:"enabled,omitempty"`
	LastValue     string                 `json:"last_value,omitempty"`
	Status        string                 `json:"status,omitempty"`
	StatusNotify  bool                   `json:"status_notify,omitempty"`
	StatusTimeout int                    `json:"status_timeout,omitempty"`
	CreatedAt     string                 `json:"created_at,omitempty"`
	UpdatedAt     string                 `json:"updated_at,omitempty"`
}

func (c *Client) ListFeeds(u string) ([]Feed, error) {
	var feeds []Feed
	resp, err := c.get("/" + u + "/feeds")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return feeds, err
	}
	return feeds, c.decodeJSON(resp, &feeds)
}

func (c *Client) CreateFeed(u string, f Feed) error {
	resp, err := c.post("/"+u+"/feeds", f)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) DeleteFeed(u, f string) error {
	resp, err := c.delete("/" + u + "/feeds/" + f)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) GetFeed(u, f string) (Feed, error) {
	var feed Feed
	resp, err := c.get("/" + u + "/feeds/" + f)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return feed, err
	}
	return feed, c.decodeJSON(resp, &feed)
}

func (c *Client) UpdateFeed(u, f string, feed Feed) error {
	resp, err := c.patch("/"+u+"/feeds/"+f, feed)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) ReplaceFeed(u, f string, feed Feed) error {
	resp, err := c.put("/"+u+"/feeds/"+f, feed)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) GetFeedDetails(u, f string) (Feed, error) {
	var feed Feed
	resp, err := c.get("/" + u + "/feeds/" + f + "/details")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return feed, err
	}
	return feed, c.decodeJSON(resp, &feed)
}

func (c *Client) AddFeedToGroup(u, g, f string) error {
	resp, err := c.post("/"+u+"/groups/"+g+"/add?feed_key="+f, nil)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) RemoveFeedFromGroup(u, g, f string) error {
	resp, err := c.post("/"+u+"/groups/"+g+"/remove?feed_key="+f, nil)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) ListFeedsFromGroup(u, g string) ([]Feed, error) {
	var feeds []Feed
	resp, err := c.get("/" + u + "/groups/" + g + "/feeds")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return feeds, err
	}
	return feeds, c.decodeJSON(resp, &feeds)
}

func (c *Client) CreateFeedInGroup(u, g string, feed Feed) error {
	resp, err := c.post("/"+u+"/groups/"+g+"/feeds", feed)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}
