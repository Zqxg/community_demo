package main

import (
	"community_demo/conf"
	"community_demo/routers"
)

func main() {
	conf.Init()
	r := routers.NewRouter()
	r.Run(conf.HttpPort)
}
