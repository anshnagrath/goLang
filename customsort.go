package main;
import ("fmt" 
		"sort"
		)
type Person struct{
	Name string
	Age int
}
func(p Person) String() string{
	return fmt.Sprintf("%s:%d",p.Name,p.Age)
}
type ByAge []Person
func (a ByAge)Len() int   		{return len(a)}
func (a ByAge)Swap(i,j int) 	{a[i],a[j]=a[j],a[i]}
func (a ByAge)Less(i,j int)bool {return a[i].Age<a[j].Age}
func main(){
	people:=[]Person{
		{"Bob",315},
		{"John",23},
		{"ascs",245},
}

sort.Sort(ByAge(people))
	fmt.Println(people)	

}
