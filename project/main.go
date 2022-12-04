package main

import (
	"LinksShortner/project/ServerEnv"
)

func main() {

	ServerEnv.RunServer()
	//a := LinkEnv.DBCheckQuery("original", "https://www.postgresql.org/docs/current/tutorial-fk.html")
	//fmt.Print(a)
}
