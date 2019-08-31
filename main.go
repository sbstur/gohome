package main

import (
	"fmt"
	//	myC "gohome/a/b/c"
	//	"gohome/c"
	"gohome/user"

	"github.com/AnuchitO/say"
)

/*
func Something(new, word string) {
	fmt.Println("Say something is ", word)
}
*/
func main() {
	u := user.User{Name: "Nong", Email: "n@g.com"}
	fmt.Println(u)
	//c.Show(u)
	//myC.Show(u)

	say.Something("nong", "hello")
}
