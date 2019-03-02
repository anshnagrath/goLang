package main

import (
	"fmt"
	"math"
)
type circle struct {
	radius float64
}
type square struct {
	side float64
}
type shape interface {
	area() float64
}
func (c circle) area() float64{
return math.Pi*c.radius*c.radius;
}
func (s square) area() float64{
	return s.side * s.side;
}
func info (s shape){
fmt.Println(s.area())
}
func main(){
	c:=circle{
		radius:34.34,
	}
	s:=square{
		side:343.44,
	}
	info(c)
	info(s)
}
