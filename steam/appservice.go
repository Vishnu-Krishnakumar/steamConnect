package steam

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type AppListReponse struct {
	AppList Apps `json:"applist"`
}
type Apps struct {
	Apps []App `json:"apps"`
}
 type App struct {
 	Name string `json:"name"`
 	ID   uint32 `json:"appid"`
 }


func  GetAppList() (apps [] App, err error) {
	const url = "https://api.steampowered.com/ISteamApps/GetAppList/v2/"
	resp, err := http.Get(url)

	if err != nil {
		err = fmt.Errorf("call to steam api failed with: %w", err)
		return
	}
	defer resp.Body.Close()
	var response AppListReponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		err = fmt.Errorf("decoding app list response failed: %w", err)
		return
	}
	apps = response.AppList.Apps
	return
}
