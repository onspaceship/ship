package client

type UserResponse struct {
	ID     string         `json:"id"`
	Name   string         `json:"name"`
	Handle string         `json:"handle"`
	Email  string         `json:"email"`
	Teams  []TeamResponse `json:"teams"`
}

type TeamResponse struct {
	ID     string `json:"id"`
	Handle string `json:"handle"`
	Name   string `json:"name"`
}

func (client *Client) GetUser() (UserResponse, error) {
	var user UserResponse

	err := client.GetJSON(client.corePath("/user"), &user)

	return user, err
}
