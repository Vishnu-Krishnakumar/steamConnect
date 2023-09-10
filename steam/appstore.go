package steam

import "fmt"

var appMap = make(map[uint32]string)

func GetApp(appId uint32) (name string, err error) {

	if name, ok := appMap[appId]; ok {

		return name,nil
	}
	apps , err := GetAppList()
	if err != nil {
		err = fmt.Errorf("call for app list failed with: %w", err)
		return
	}
	for _, app := range apps{
		if appId == app.ID{
			name = app.Name
		}
		appMap[app.ID] = app.Name
	}
	return
}