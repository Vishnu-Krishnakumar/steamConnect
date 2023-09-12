package steam

import (
	"encoding/json"
	"time"
)

type Friend struct {
	SteamId      string    `json:"steamid"`
	Relationship string    `json:"relationship"`
	FriendSince  time.Time `json:"friend_since"`
}

func (f *Friend) UnmarshalJSON(data []byte) error {

	type FriendJson struct {
		SteamId      string    `json:"steamid"`
		Relationship string    `json:"relationship"`
		FriendSince  int64    `json:"friend_since"`
	}
	var temp FriendJson 

	if err := json.Unmarshal(data, &temp); err != nil {       
		return err     
	 }     
	f.SteamId = temp.SteamId
	f.Relationship = temp.Relationship
	f.FriendSince = time.Unix(temp.FriendSince,0)
	return nil
}