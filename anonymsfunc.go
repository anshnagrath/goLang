package main

import "fmt"

func main (){
	//function returned from a function
	s2 := bar()
	fmt.Printf("%T\n",s2);
	//returned s2 as a function and then we called s2 to generate a output of type int
	s3:= s2()
	fmt.Println(s3);
//or we can do this
	fmt.Println(bar()())



	s1 := foo()
	fmt.Println(s1)
	func(s int){
		sum:=22;
		sum+=s;
		fmt.Println(sum);
	}(42)
	f:=func (s string)  {
		fmt.Println(s)
	}
	f("test")
		}
func  foo() string{
	s := "test"
	return s
}
func bar() func() int {
	return func() int{
		return 431
	}
}