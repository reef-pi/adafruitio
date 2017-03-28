package adafruitio

type Dashboard struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Key         string  `json:"key"`
	Blocks      []Block `json:"blocks"`
}

func (c *Client) ListDashboards(u string) ([]Dashboard, error) {
	var dashboards []Dashboard
	resp, err := c.get("/" + u + "/dashboards")
	if err != nil {
		return dashboards, err
	}
	return dashboards, c.decodeJSON(resp, &dashboards)
}

func (c *Client) CreateDashboard(u string, d Dashboard) error {
	_, err := c.post("/"+u+"/dashboards", d)
	return err
}

func (c *Client) DeleteDashboard(u, d string) error {
	_, err := c.delete("/" + u + "/dashboards/" + d)
	return err
}

func (c *Client) GetDashboard(u, d string) (Dashboard, error) {
	var dashboard Dashboard
	resp, err := c.get("/" + u + "/dashboards/" + d)
	if err != nil {
		return dashboard, err
	}
	return dashboard, c.decodeJSON(resp, &dashboard)
}

func (c *Client) UpdateDashboard(u, d string, dashboard Dashboard) error {
	_, err := c.patch("/"+u+"/dashboards/"+d, dashboard)
	return err
}

func (c *Client) ReplaceDashboard(u, d string, dashboard Dashboard) error {
	_, err := c.put("/"+u+"/dashboards/"+d, dashboard)
	return err
}
