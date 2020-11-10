package main

import (
	"fmt"

	"github.com/dezhiShen/music-webadv/internal/configutil"
)

func main() {
	err := configutil.LoadConfig("../../configs", "webadv")
	if err != nil {
		panic(err)
	}
	fmt.Println(configutil.GetConfigString("webadv", "username"))
	fmt.Println(configutil.GetConfigString("webadv", "password"))
	fmt.Println(configutil.GetConfigString("webadv", "url"))
	fmt.Println("hello world !")
}
