package main
import ( "fmt"
		"runtime"
		"sync"
	)
		
var wg  sync.WaitGroup;	
func main(){
	fmt.Println("OS\t\t",runtime.GOOS)
	fmt.Println("ARCH\t\t",runtime.GOARCH)
	fmt.Println("CPU\t\t",runtime.NumCPU())
	fmt.Println("Goroutines\t",runtime.NumGoroutine())
	wg.Add(1)
	go foo()
	bar()
	fmt.Println("Goroutinesafter\t",runtime.NumGoroutine())
	wg.Wait()
}
func foo(){
	
for i:=0;i<10;i++{
	fmt.Println("foo",i)
}
wg.Done()
}
func bar(){
	for j:=0;j<10;j++{
		fmt.Println("bar",j)
	}
	}