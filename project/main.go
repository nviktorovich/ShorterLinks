package main

import (
	"LinksShortner/project/LinkEnv"
	"fmt"
)

func main() {

	l := LinkEnv.NewLink("https://gobyexample.com/random-numbers")
	l.WriteToBD()

	s := LinkEnv.SearchInDB("https://2")
	fmt.Println(s)
}
