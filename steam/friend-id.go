package steam

import (
	"encoding/json"
	"fmt"
	"net/http"
)
type Key struct {
	apiKey string
}

func NewKey(apiKey string) Key{
	return Key{
		
		apiKey,
	}
}

type FriendsList struct {
	Friendlist Friends `json:"friendslist"`
}

type Friends struct {
	Friends [] Friend `json:"friends"`
}

func (k Key) GetFriendsList (steamId string) (friends  Friends , err error) {
	url := fmt.Sprintf("http://api.steampowered.com/ISteamUser/GetFriendList/v0001/?key=%s&steamid=%s&relationship=friend", k.apiKey, steamId)
	fList, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("call to steam api failed with: %w", err)
		return
	}
	defer fList.Body.Close()
	var friendsList FriendsList
	err = json.NewDecoder(fList.Body).Decode(&friendsList)
	if err != nil {
		err = fmt.Errorf("decoding friends list response failed: %w", err)
		return
	}
	
	friends = friendsList.Friendlist
	return
}