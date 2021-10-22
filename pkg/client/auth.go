package client

type UserResponse struct {
	ID     string         `json:"id"`
	Name   string         `json:"name"`
	Handle string         `json:"handle"`
	Email  string         `json:"email"`
	Teams  []TeamResponse `json:"teams"`
}

type TeamResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func (client *Client) GetUser() (UserResponse, error) {
	var user UserResponse

	err := client.GetJSON(client.corePath("/user"), &user)

	return user, err
}

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
