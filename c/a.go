package c

import (
	"fmt"
	"gohome/user"
)

func Show(u user.User) {
	fmt.Println("Outer Show", u)
}
