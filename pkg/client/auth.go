package client

func (client *Client) GetTokenURL() (string, error) {
	resp, err := client.Get(client.corePath("/authentications/new"))
	if err != nil {
		return "", err
	}

	url, err := resp.Location()
	if err != nil {
		return "", err
	}

	return url.String(), nil
}

type TokenResponse struct {
	Token string `json:"token"`
}

func (client *Client) GetToken(code string) (string, error) {
	var token TokenResponse

	err := client.GetJSON(client.corePath("/authentications/%s", code), &token)
	if err != nil {
		return "", err
	}

	return token.Token, nil
}
