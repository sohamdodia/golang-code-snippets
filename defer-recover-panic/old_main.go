package main

import "fmt"

// func accessElement(a []int, index int) int {
// 	if len(a) > index {
// 		return a[index]
// 	}
//  	panic("Out of bound condition")
// }

// func main() {
// 	a := []int{1, 2, 3}
// 	fmt.Println(accessElement(a, 2))
// 	fmt.Println(accessElement(a, 3))
// }

func defFooStart() {
	fmt.Println("defFooStart () executed")
}

func defFooEnd() {
	fmt.Println("defFooEnd() executed")
}

func defMain() {
	fmt.Println("defMain() executed")
}

func foo() {
	fmt.Println("foo() executed")
	defer defFooStart()
	panic("panic from foo()")
	defer defFooEnd()
	fmt.Println("foo() done")
}

func a() {
	fmt.Println("main() started")
	defer defMain()
	foo()
	fmt.Println("main() done")
}
