package main

import "fmt"

type triangle struct{
	base float64
	height float64
}
type square struct{
	length float64
}

func (t triangle) getArea() float64{
	return 0.5*t.base*t.height
}

func (t square) getArea() float64{
	return t.length*t.length
}

type shape interface {
	getArea() float64
}

func printArea(s shape){
	fmt.Println(s.getArea());
}

func main(){
	t := triangle{5,4}
	s := square{5}
	printArea(t)
	printArea(s)
}