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
 
  if err != nil{
	 fmt.Printf("error: %v", err)
  }

   for _, game := range ownedGames.Games{
   
    name, _ := steam.GetApp(game.AppId)
    fmt.Printf("app ID: %d\n app name: %s\n", game, name)
    
  }
  
}

