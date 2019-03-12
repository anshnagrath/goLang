package main
import (

	"fmt"
	"runtime"
	"time"
	"context"
)
func main(){
ctx,cancel:=context.WithCancel(context.Background())
fmt.Println("error check 1",ctx.Err())
fmt.Println("number of go routine",runtime.NumGoroutine())
go func(){
	n:=0
	for{
		select{
		case <-ctx.Done():
				return
			default : 
					n++
					time.Sleep(time.Millisecond * 200)
					fmt.Println("working",n)

		}
	}
}()
time.Sleep(time.Millisecond * 200)
fmt.Println("error check 2",ctx.Err())
fmt.Println("number of go routine2",runtime.NumGoroutine())
fmt.Println("about to cancel")
cancel()
fmt.Println(time.Second*2)
fmt.Println("canceled")
fmt.Println("number of go routine3",runtime.NumGoroutine())
}