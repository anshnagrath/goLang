package main
import "fmt"
func main(){
c:=make(chan int)
go foo(c)
bar(c)
fmt.Println("Exit")	
}
func foo(c chan<- int){
//sending to an int
c <- 42

}
func bar(c <-chan int){
	fmt.Println(<-c)
}