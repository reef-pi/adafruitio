package adafruitio

type Token struct {
	Token string `json:"token"`
}

func (c *Client) ListTokens(u string) ([]Token, error) {
	var tokens []Token
	resp, err := c.get("/" + u + "/tokens")
	if err != nil {
		return tokens, err
	}
	return tokens, c.decodeJSON(resp, &tokens)
}

func (c *Client) GetToken(u, t string) (Token, error) {
	var token Token
	resp, err := c.get("/" + u + "/tokens/" + t)
	if err != nil {
		return token, err
	}
	return token, c.decodeJSON(resp, &token)
}

func (c *Client) CreateToken(u string, t Token) error {
	_, err := c.post("/"+u+"/tokens", t)
	return err
}

func (c *Client) UpdateToken(u, t string, tok Token) error {
	_, err := c.patch("/"+u+"/tokens/"+t, tok)
	return err
}

func (c *Client) ReplaceToken(u, t string, tok Token) error {
	_, err := c.put("/"+u+"/tokens/"+t, tok)
	return err
}

func (c *Client) DeleteToken(u, t string) error {
	_, err := c.delete("/" + u + "/tokens/" + t)
	return err
}
