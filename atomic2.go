package main
import ( "fmt"
		"sync"
		"sync/atomic"	
)
var wg sync.WaitGroup


func main(){
var	increamentor int64 = 0;
	gs:=100
	wg.Add(gs)
	for i:=0; i<gs;i++{
		go func(){
		
			atomic.AddInt64(&increamentor,1)
			// runtime.Gosched()
	
			fmt.Println(atomic.LoadInt64(&increamentor))
			//mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}