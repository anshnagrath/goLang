package main
import "fmt"
func main(){
	f()
	fmt.Println("Returned from f")
}
func f(){
	defer func(){
		if r:=recover(); r!=nil{
			fmt.Println("Recovered i from f",r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
}
func g(i int){
	if i > 3{
		fmt.Println("Panicking!")
		fmt.Sprintf("%v",i)
		//passing string valueto panic
		panic(fmt.Sprintf("%v",i))
	}
	defer fmt.Println("Defer in g",i)
	g(i+1)
}