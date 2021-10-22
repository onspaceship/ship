package client

type RegistryTokenResponse struct {
	Token string `json:"token"`
}

func (client *Client) GetRegistryToken(team string) (string, error) {
	var token RegistryTokenResponse

	err := client.GetJSON(client.corePath("/teams/%s/token", team), &token)
	if err != nil {
		return "", err
	}

	return token.Token, nil
}
