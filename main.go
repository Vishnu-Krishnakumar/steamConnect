package main

import (
	"booking-app/Desktop/SteamConnect/steam"
	"fmt"
	"os"
)


func main(){

  apiKey := os.Getenv("STEAM_API_KEY")
  steamClient := steam.NewClient(apiKey)
  ownedGames, err := steamClient.GetOwnedGames("76561197984646227")
  friends := steam.NewKey(apiKey)
  friendsList, err2 := friends.GetFriendsList("76561197984646227")

  if err != nil{
	 fmt.Printf("error: %v", err)
  }

   for _, game := range ownedGames.Games{
   
    name, _ := steam.GetApp(game.AppId)
    fmt.Printf("app ID: %d\napp name: %s\n", game.AppId, name)
    
  }
  if err2 != nil{
    fmt.Printf("error: %v", err2)
   }
 
     for _, friend := range friendsList.Friends{
    
      
      fmt.Printf("friend: %v %s %v \n",friend.SteamId,friend.Relationship,friend.FriendSince )
     
    }
  
}

