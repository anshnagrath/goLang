package main

import "fmt"

func main (){

	bar("james")
	foo()
	//return type functions
	s:= test();
	fmt.Println(s)
	//function with more than two returns
	x,y:=two("ccsa","acsdcsa")
	fmt.Println(x,y);
	c:=sum(23,12,12,12,34,23,12,544,6,673,33)
	fmt.Println(c,"Total")
}
func bar(s string){
	fmt.Println("bar called")
}
func foo(){
	fmt.Println("hello there")
}
func test() string {
	return fmt.Sprint("testing the return")
}
func two(x ,y string)(string,bool){
a:= fmt.Sprint("test");
b:=true;
return a,b
}
func sum(x ...int) int{
	fmt.Printf("%T\n",x)
	sum:=0;
	for _,v:=range x{
		sum+=v
	}
	return sum
}