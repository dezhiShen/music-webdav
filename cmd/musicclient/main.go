package main

import (
	"fmt"

	"github.com/dezhiShen/music-webadv/internal/configutil"
	"github.com/dezhiShen/music-webadv/internal/music"
)

func main() {
	err := configutil.LoadConfigs("../../configs", []string{"webadv", "local"})
	if err != nil {
		panic(err)
	}
	syncPlayList := configutil.GetConfigBool("local", "sync-playlist")
	playList, err := music.LoadAllPlayList(syncPlayList)
	if err != nil {
		panic(err)
	}
	//遍历现在的列表
	for _, v := range playList {
		fmt.Println(v.Id)
	}
	fmt.Println(configutil.GetConfigString("webadv", "username"))
	fmt.Println(configutil.GetConfigString("webadv", "password"))
	fmt.Println(configutil.GetConfigString("webadv", "url"))
	fmt.Println(configutil.GetConfigString("webadv", "workspace"))

	fmt.Println("hello world !")
}
