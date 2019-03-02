package main

import "fmt"

func main (){
	//defering a function is to excecute a function at exit
defer foo()
bar()
}
func foo(){
	fmt.Println("foo")
}
func bar(){
	fmt.Println("bar")
}