package adafruitio

type Trigger struct {
	Name string `json:"name"`
}

func (c *Client) ListTriggers(u string) ([]Trigger, error) {
	var triggers []Trigger
	resp, err := c.get("/" + u + "/triggers")
	if err != nil {
		return triggers, err
	}
	return triggers, c.decodeJSON(resp, &triggers)
}

func (c *Client) GetTrigger(u, t string) (Trigger, error) {
	var trigger Trigger
	resp, err := c.get("/" + u + "/triggers/" + t)
	if err != nil {
		return trigger, err
	}
	return trigger, c.decodeJSON(resp, &trigger)
}

func (c *Client) CreateTrigger(u string, t Trigger) error {
	_, err := c.post("/"+u+"/triggers", t)
	return err
}

func (c *Client) UpdateTrigger(u, t string, tok Trigger) error {
	_, err := c.patch("/"+u+"/triggers/"+t, tok)
	return err
}

func (c *Client) ReplaceTrigger(u, t string, tok Trigger) error {
	_, err := c.put("/"+u+"/triggers/"+t, tok)
	return err
}

func (c *Client) DeleteTrigger(u, t string) error {
	_, err := c.delete("/" + u + "/triggers/" + t)
	return err
}
