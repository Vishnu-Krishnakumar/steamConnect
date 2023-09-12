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

// UnmarshalJSON implements the Unmarsheler interface to allow map steams FriendSince to a time.Time
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
	// Converting a int64 unix time stamp to a time.Time format
	f.FriendSince = time.Unix(temp.FriendSince,0)
	return nil
}