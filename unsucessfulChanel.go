package main
import "fmt"
func main(){
	//if both reciever and sender are not ready then it is a deadllock or no space for sending
c:= make(chan int)
go func(){
	c <- 42
}()
fmt.Println(<- c)

}