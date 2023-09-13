package steam

import (
	"fmt"
	"net/http"
	"os"
)

func CallGameList(w http.ResponseWriter, req *http.Request) {

	apiKey := os.Getenv("STEAM_API_KEY")
	steamClient := NewClient(apiKey)
	ownedGames, err := steamClient.GetOwnedGames("76561197984646227")
	fmt.Fprintf(w, "hello\n")
	if err != nil {
		fmt.Fprintf(w, "error: %v", err)
	}

	for _, game := range ownedGames.Games {

		name, _ := GetApp(game.AppId)
		fmt.Fprintf(w, "app ID: %d\napp name: %s\n", game.AppId, name)

	}

}