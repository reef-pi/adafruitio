package adafruitio

import (
	"github.com/google/go-querystring/query"
)

type ListDataOptions struct {
	StartTime  string `url:"start_time,omitempty"`
	EndTime    string `url:"end_time,omitempty"`
	Limit      int    `url:"limit,omitempty"`
	Resolution int    `url:"resolution,omitempty"`
}

type Data struct {
	ID           int         `json:"id,omitempty"`
	Value        interface{} `json:"value"`
	Feed         int         `json:"feed_id,omitempty"`
	Group        int         `json:"group_id,omitempty"`
	Expiration   string      `json:"expiration,omitempty"`
	Lattitude    float32     `json:"lat,omitempty"`
	Longitude    float32     `json:"lon,omitempty"`
	Element      float32     `json:"ele,omitempty"`
	CompletedAt  string      `json:"completed_at,omitempty"`
	CreatedAt    string      `json:"created_at,omitempty"`
	UpdatedAt    string      `json:"updated_at,omitempty"`
	CreatedEPOCH string      `json:"epoch,omitempty"`
}

func (c *Client) ListData(u, f string, o ListDataOptions) ([]Data, error) {
	var dataList []Data
	v, err := query.Values(o)
	if err != nil {
		return dataList, err
	}
	resp, err := c.get("/" + u + "/feeds/" + f + "/data?" + v.Encode())
	if resp != nil {
		defer resp.Body.Close()
	}

	if err != nil {
		return dataList, err
	}
	return dataList, c.decodeJSON(resp, &dataList)
}

func (c *Client) SubmitData(u, f string, d Data) error {
	resp, err := c.post("/"+u+"/feeds/"+f+"/data", d)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) SubmitDataInBatch(u, f string, d []Data) error {
	resp, err := c.post("/"+u+"/feeds/"+f+"/data/batch", d)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) PreviousDataInQueue(u, f string) (Data, error) {
	var d Data
	resp, err := c.get("/" + u + "/feeds/" + f + "/data/previous")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return d, err
	}
	return d, c.decodeJSON(resp, &d)
}

func (c *Client) NextDataInQueue(u, f string) (Data, error) {
	var d Data
	resp, err := c.get("/" + u + "/feeds/" + f + "/data/next")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return d, err
	}
	return d, c.decodeJSON(resp, &d)
}

func (c *Client) LastDataInQueue(u, f string) (Data, error) {
	var d Data
	resp, err := c.get("/" + u + "/feeds/" + f + "/data/last")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return d, err
	}
	return d, c.decodeJSON(resp, &d)
}

func (c *Client) DeleteData(u, f, id string) error {
	resp, err := c.delete("/" + u + "/feeds/" + f + "/data/" + id)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) GetData(u, f, id string) (Data, error) {
	var d Data
	resp, err := c.get("/" + u + "/feeds/" + f + "/data/" + id)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return d, err
	}
	return d, c.decodeJSON(resp, &d)
}

func (c *Client) UpdateData(u, f, id string, d Data) error {
	resp, err := c.patch("/"+u+"/feeds/"+f+"/data/"+id, d)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) ReplaceData(u, f, id string, d Data) error {
	resp, err := c.put("/"+u+"/feeds/"+f+"/data/"+id, d)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) ListDataFromGroup(u, g, f string) ([]Data, error) {
	var dataList []Data
	resp, err := c.get("/" + u + "/groups/" + g + "/feeds/" + f + "/data")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return dataList, err
	}
	return dataList, c.decodeJSON(resp, &dataList)
}

func (c *Client) SubmitDataInGroup(u, g, f string, d Data) error {
	resp, err := c.post("/"+u+"/groups/"+g+"/feeds/"+f+"/data", d)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) SubmitDataInGroupInBatch(u, g, f string, d []Data) error {
	resp, err := c.post("/"+u+"/groups/"+g+"/feeds/"+f+"/data/batch", d)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}
