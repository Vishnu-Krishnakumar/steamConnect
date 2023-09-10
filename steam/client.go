package steam

import (
	"encoding/json"
	"fmt"
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
type OwnedGamesResponse struct{
	Owned OwnedGames `json:"response"`
}
type OwnedGames struct {
	Games []Game `json:"games"`
	
}

func (c Client) GetOwnedGames (steamId string) (games OwnedGames , err error) {
	url := fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json", c.apiKey, steamId)
	resp, err := http.Get(url)
	
	if err != nil {
		err = fmt.Errorf("call to steam api failed with: %w", err)
		return
	}
	defer resp.Body.Close()
	var response OwnedGamesResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		err = fmt.Errorf("decoding games response failed: %w", err)
		return
	}
	games = response.Owned

	return
}


