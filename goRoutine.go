package main
import ("fmt"
		"sync"	
		"runtime"
		)
func main(){
	var wg sync.WaitGroup
	wg.Add(2)
	go func(){
		fmt.Println("go Routine 1")
	wg.Done()	
	}()
	go func(){
	fmt.Println("go Routine 2")	
	wg.Done()
		}()
	fmt.Println("Number of go routine",runtime.NumGoroutine())	
	fmt.Println("Number of CPU",runtime.NumCPU())
}