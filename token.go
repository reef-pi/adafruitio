package adafruitio

type Token struct {
	Token string `json:"token"`
}

func (c *Client) ListTokens(u string) ([]Token, error) {
	var tokens []Token
	resp, err := c.get("/" + u + "/tokens")
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return tokens, err
	}
	return tokens, c.decodeJSON(resp, &tokens)
}

func (c *Client) GetToken(u, t string) (Token, error) {
	var token Token
	resp, err := c.get("/" + u + "/tokens/" + t)
	if resp != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return token, err
	}
	return token, c.decodeJSON(resp, &token)
}

func (c *Client) CreateToken(u string, t Token) error {
	resp, err := c.post("/"+u+"/tokens", t)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) UpdateToken(u, t string, tok Token) error {
	resp, err := c.patch("/"+u+"/tokens/"+t, tok)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) ReplaceToken(u, t string, tok Token) error {
	resp, err := c.put("/"+u+"/tokens/"+t, tok)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}

func (c *Client) DeleteToken(u, t string) error {
	resp, err := c.delete("/" + u + "/tokens/" + t)
	if resp != nil {
		defer resp.Body.Close()
	}
	return err
}
