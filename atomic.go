package main
import ("fmt"
		"runtime"
		"sync"
		"sync/atomic"
		)
var wg sync.WaitGroup		
func main(){
fmt.Println("CPU's",runtime.NumCPU())
fmt.Println("go routine",runtime.NumGoroutine())
var counter int64
const gs =100
wg.Add(gs)

for i:=0;i<100;i++{
	go func(){

	atomic.AddInt64(&counter,1)
	fmt.Println("counter\t",atomic.LoadInt64(&counter))
	runtime.Gosched()



	wg.Done()
	//fmt.Println("go routine",runtime.NumGoroutine())
	}()

}
wg.Wait()
fmt.Println("go routine",runtime.NumGoroutine())
fmt.Println("count",counter)
}