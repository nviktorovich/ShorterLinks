package main

import (
	"LinksShortner/project/Configuration"
	"LinksShortner/project/ServerEnv"
)

func main() {
	Configuration.ReadConfig()
	ServerEnv.RunServer()

}
