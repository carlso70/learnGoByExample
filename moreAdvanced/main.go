package main

import "fmt"

type person struct {
	name string
	age  int
}

type rect struct {
	height, width int
}

func main() {
	fmt.Println("vim-go")
	//pointers()
	//structs()
	structsWithMethods()
}

/*
func pointers() {
	i := 1
	fmt.Println("i: ", i)

	zeroval(i)
	fmt.Println("zeroval(ival int) doesnt edit value zeroval: ", i)

	zeroptr(&i)
	fmt.Println("zeroptr(ival *int) does edit it the value zeroptr: ", i)
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(ival *int) {
	*ival = 0
}

func structs() {
	s := person{"bob", 69}
	fmt.Println(s)

	sp := person{name: "Jim", age: 14}
	fmt.Println(sp)

	// Edit s with using pointers
	sx := &s
	sx.age = 45
	fmt.Println(*sx)
}
*/

func structsWithMethods() {
	r := rect{width: 10, height: 15}

	fmt.Println("Perimeter of rect{ 10, 15 } = ", r.perim())
	fmt.Println("Area of rect{ 10, 15 } = ", r.area())

	rp := &r
	fmt.Println("Perimeter of rect{ 10, 15 } = ", rp.perim())
	fmt.Println("Area of rect{ 10, 15 } = ", rp.area())

}

func (r rect) perim() int {
	return 2*r.width + 2*r.height
}

func (r *rect) area() int {
	return r.height * r.width
}
