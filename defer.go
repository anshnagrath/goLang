package main
import "fmt"
func main (){
	// defer prints the value evaluated at the the time of excecution
	i:=0
	defer fmt.Println(i)
	i++
	b()
	return
}
func b(){
	for i:=0;i<4;i++{
		defer fmt.Println(i,"last in first out manner")
	}
}