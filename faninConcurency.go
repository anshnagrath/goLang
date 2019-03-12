package main
import ("fmt"
		"sync")
func main(){
odd := make(chan int)
even := make(chan int)
fanin :=make(chan int)
go send(even,odd)
go recieve(even,odd,fanin)
for v:= range fanin{
	fmt.Println(v)
}
fmt.Println("end")
}
func send(even,odd chan<- int){
for i:=0;i<=100;i++{
	if i%2 == 0{
		even <- i
	}else{
		odd <- i
	}

}
close(even)
close(odd)
}
func recieve(even,odd,fanin chan int){
	var wg sync.WaitGroup
	wg.Add(2)
	go func(){
		for v:= range even {
		fanin <- v
		}
		wg.Done()
	}()
	go func(){
		for v:= range odd {
		fanin <- v
		}
		wg.Done()
	}()
	fmt.Println(fanin,"fanin")
	wg.Wait()
	close(fanin)


}