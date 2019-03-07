package main
import ("fmt"
		"runtime"
		"sync"
		)
var wg sync.WaitGroup		
func main(){
fmt.Println("CPU's",runtime.NumCPU())
fmt.Println("go routine",runtime.NumGoroutine())
counter:=0
const gs =100
wg.Add(gs)
var mu sync.Mutex
for i:=0;i<100;i++{
	go func(){
	mu.Lock()
	v := counter	
	runtime.Gosched()
	v++
	counter = v
	mu.Unlock()
	wg.Done()
	//fmt.Println("go routine",runtime.NumGoroutine())
	}()

}
wg.Wait()
fmt.Println("go routine",runtime.NumGoroutine())
fmt.Println("count",counter)
}