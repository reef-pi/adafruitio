package adafruitio

type Permission struct {
	ID         string `json:"id,omitempty"`
	UserID     string `json:"user_id,omitempty"`
	Scope      string `json:"scope,omitempty"`
	ScopeValue string `json:"scope_value,omitempty"`
	Model      string `json:"model,omitempty"`
	ObjectID   string `json:"object_id,omitempty"`
	CreatedAt  string `json:"created_at,omitempty"`
	UpdatedAt  string `json:"updated_at,omitempty"`
}

func (c *Client) ListPermissions(u, t string) ([]Permission, error) {
	var permissions []Permission
	resp, err := c.get("/" + u + "/type/" + t + "/acl")
	if err != nil {
		return permissions, err
	}
	return permissions, c.decodeJSON(resp, &permissions)
}

func (c *Client) CreatePermission(u, t string, p Permission) error {
	_, err := c.post("/"+u+"/type/"+t+"/acl", p)
	return err
}

func (c *Client) UpdatePermission(u, t, a string, p Permission) error {
	_, err := c.patch("/"+u+"/type/"+t+"/acl/"+a, p)
	return err
}

func (c *Client) ReplacePermission(u, t, a string, p Permission) error {
	_, err := c.put("/"+u+"/type/"+t+"/acl/"+a, p)
	return err
}

func (c *Client) DeletePermission(u, t, a string) error {
	_, err := c.delete("/" + u + "/type/" + t + "/acl/" + a)
	return err
}
