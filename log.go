package main
import ("os"
		"fmt"
		"log")
func main(){
f,err:= os.Create("log.txt")
if err != nil {
	fmt.Println(err)
}
defer f.Close()
log.SetOutput(f)
f2,err := os.Open("nocsad.txt")
if(err) !=nil {
	log.Println("err happened",err)
}
defer f2.Close()
fmt.Println("check the log ")
}
