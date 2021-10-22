package client

type AppResponse struct {
	ID     string `json:"id"`
	Handle string `json:"handle"`
	Name   string `json:"name"`
}

func (client *Client) GetApp(teamId string, appId string) (AppResponse, error) {
	var app AppResponse

	err := client.GetJSON(client.corePath("/teams/%s/apps/%s", teamId, appId), &app)

	return app, err
}
