package main

import "fmt"

func main() {
	// use _ to throw away any returns
	n, err := fmt.Println("Hello we are inside main")
	fmt.Println(n);
	fmt.Println(err)
	foo()
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			fmt.Println("loop terminated",i)
		}
	}
	fmt.Println("exited from foo")
}
func foo() {
	fmt.Println("inside foo")
}
