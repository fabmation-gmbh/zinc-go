package zinc

// clientOpt is a Client option function.
type clientOpt func(*Client)

// SetZincServer sets the Zinc server URL.
func SetZincServer(url string) clientOpt {
	return clientOpt(func(c *Client) {
		c.c.SetHostURL(url + "/api")
	})
}

// SetBasicAuth configures the client to authenticate with Basic Auth.
func SetBasicAuth(username, password string) clientOpt {
	return clientOpt(func(c *Client) {
		c.c.SetBasicAuth(username, password)
	})
}
