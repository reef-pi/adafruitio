package adafruitio

type BlockFeed struct {
	ID    string `json:"id"`
	Feed  Feed   `json:"feed"`
	Group Group  `json:"group"`
}

type Block struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	Key         string      `json:"key"`
	VisualType  string      `json:"visual_type"`
	Column      int         `json:"column"`
	Row         int         `json:"row"`
	SizeX       int         `json:"size_x"`
	SizeY       int         `json:"size_y"`
	Feeds       []BlockFeed `json:"block_feeds"`
}

func (c *Client) ListBlocks(u, d string) ([]Block, error) {
	var blocks []Block
	resp, err := c.get("/" + u + "/dashboards/" + d + "/blocks")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return blocks, err
	}
	return blocks, c.decodeJSON(resp, &blocks)
}

func (c *Client) GetBlock(u, d, b string) (Block, error) {
	var block Block
	resp, err := c.get("/" + u + "/dashboards/" + d + "/blocks/" + b)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return block, err
	}
	return block, c.decodeJSON(resp, &block)
}

func (c *Client) CreateBlock(u, d string, b Block) error {
	resp, err := c.post("/"+u+"/dashboards/"+d+"/blocks", b)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) UpdateBlock(u, d, b string, block Block) error {
	resp, err := c.patch("/"+u+"/dashboards/"+d+"/blocks/"+b, block)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) ReplaceBlock(u, d, b string, block Block) error {
	resp, err := c.put("/"+u+"/dashboards/"+d+"/blocks/"+b, block)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) DeleteBlock(u, d, b string) error {
	resp, err := c.delete("/" + u + "/dashboards/" + d + "/blocks/" + b)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}
