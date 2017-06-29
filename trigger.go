package adafruitio

type Trigger struct {
	Name string `json:"name"`
}

func (c *Client) ListTriggers(u string) ([]Trigger, error) {
	var triggers []Trigger
	resp, err := c.get("/" + u + "/triggers")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return triggers, err
	}
	return triggers, c.decodeJSON(resp, &triggers)
}

func (c *Client) GetTrigger(u, t string) (Trigger, error) {
	var trigger Trigger
	resp, err := c.get("/" + u + "/triggers/" + t)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return trigger, err
	}
	return trigger, c.decodeJSON(resp, &trigger)
}

func (c *Client) CreateTrigger(u string, t Trigger) error {
	resp, err := c.post("/"+u+"/triggers", t)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) UpdateTrigger(u, t string, tok Trigger) error {
	resp, err := c.patch("/"+u+"/triggers/"+t, tok)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) ReplaceTrigger(u, t string, tok Trigger) error {
	resp, err := c.put("/"+u+"/triggers/"+t, tok)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) DeleteTrigger(u, t string) error {
	resp, err := c.delete("/" + u + "/triggers/" + t)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}
