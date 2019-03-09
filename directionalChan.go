package main 
import "fmt"
func main(){
	//we can only send the values to a channel
	c:= make(chan<- int,2)

	fmt.Println(c)
	// we can only recieve values from a channel
	g:= make(<-chan int,2) 
	
	fmt.Println(<-g)
}