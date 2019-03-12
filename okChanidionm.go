package main
import "fmt"
func main(){
	e:= make(chan int)
	o:= make(chan int)
	q:= make(chan bool)
	go send(e,o,q)
	recieve(e,o,q)
}
func send(e,o chan<- int,q chan<- bool){
	for i:=0;i<=100;i++{
		if i%2 == 0{
			o <- i	
		}else{
			o <- i
		}
	}
	close(q)
}
func recieve(e,o <-chan int,q <-chan bool){
	for{
		select{
		case v:= <-e:
				fmt.Println("even",v)
		case v:= <- o:
				fmt.Println("odd",v)
		case i,ok:= <- q:	
			if ok{
				fmt.Println("case",ok,i)
			}else{
				fmt.Println("case !ok",i)
			}
			return		
		}
	}
}