package main
import "fmt"
func main(){
	c:=make(chan int)
	for i:=0;i<10;i++{
		go func(){
			for i:=0;i<10;i++{
				c  <- i
			}
		}()
	}
	recieve(c)


	}
	func recieve(c chan int){
		for v:=range c {
			fmt.Println(v)
		}
	}