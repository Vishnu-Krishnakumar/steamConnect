package steam

import (
	"fmt"
	"io"
	"net/http"
)

type Client struct {
	apiKey string
}

func NewClient(apiKey string) Client {
	return Client{
		apiKey,
	}
}

func (c Client) GetOwnedGames (steamId string) (body []byte, err error) {
	url := fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json", c.apiKey, steamId)
	resp, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("call to steam api failed with: %w/n", err)
		return
	}
	defer resp.Body.Close()
	body, err = io.ReadAll(resp.Body)
	return
}