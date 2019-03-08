package main
import ( "fmt"
		"sync"
		"runtime"		
)
var wg sync.WaitGroup
var mu sync.Mutex

func main(){
	increamentor:=0;
	gs:=100
	wg.Add(gs)
	for i:=0; i<gs;i++{
		go func(){
			mu.Lock()
			v := increamentor
			runtime.Gosched()
			v++
			increamentor = v
			fmt.Println(increamentor)
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
}