package main

import (
	"fmt"

	"gohello/5-init/lib1"
	"gohello/5-init/lib2"
)

func init() {
	fmt.Println("libmain init")
}

func main() {
	lib1.Func1()
	lib2.Func1()
}
