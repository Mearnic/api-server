package main

import (
	"api-server/src/network"
)

func main() {
	e := network.New()
	servers := network.InitApi(e)
	err := servers.Start(":1323")
	servers.Logger.Fatal(err)
}
