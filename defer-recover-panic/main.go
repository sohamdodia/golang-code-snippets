package main

import "fmt"

func defFoo() {
	fmt.Println("deffoo() started")
	if r := recover(); r != nil {
		fmt.Println("WHOA! Program is panicking with value", r)
	}
	fmt.Println("deffoo() done")
}

func normMain() {
	fmt.Println("normMain() started")
	defer defFoo()
	panic("HELP")
	fmt.Println("normmain() done")
}

func main() {
	fmt.Println("main() started")
	normMain()
	fmt.Println("main() done")
}
