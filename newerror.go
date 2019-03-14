package main
import ("errors"
		"fmt"
		"log")
func main(){
	s,err:=squrt(2)
	if err != nil{
		fmt.Println(err)
		
	}else{
		fmt.Println(s)
	}

	n,er :=squrt(-2)
	if er != nil{
		fmt.Println(er)
		log.Fatalln(er)
		fmt.Errorf("nagative",er)
	}else{
		fmt.Println(n)
	}

}		
func squrt (f float64)(float64,error){
	if f>0 {
		return f*f ,nil
	}else{
	return 0,errors.New("Negative number")
	}

}