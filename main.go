package main

import (
	"booking-app/Desktop/SteamConnect/steam"
	"fmt"
	"os"
)


func main(){

  apiKey := os.Getenv("steam_api_key")
  steamClient := steam.NewClient(apiKey)
  response, err := steamClient.GetOwnedGames("76561197984646227")
  if err != nil{
	 fmt.Printf("error: %v", err)
  }
  fmt.Printf("%s", response)
}

