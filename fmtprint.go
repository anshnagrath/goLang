package main

import "fmt"

var y = 42
//to create own types
type test int

var b test
func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", y)
	x := 3
	//this will print with a line
	fmt.Println("Block scpoed variable", x)
	//this will print the format but not the line
	fmt.Printf("%b\n", x)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	//to  print in the formof string
	s:= fmt.Sprintf("%#x\n", y)
	fmt.Println(s)
	//%v is the default value  while printing
	fmt.Printf("%v\n",x)
	//to print out the type
	fmt.Printf("%T\n",b)
	//conversion
	x = int(b)
	//t is tab
	S:= fmt.Sprintf("%v\t%v\t",x,y)
	fmt.Println(S)
}
