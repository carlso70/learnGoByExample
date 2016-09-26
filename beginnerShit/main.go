package main

import "fmt"
import "time"

const s string = "hi"

func main() {
	//doSomething()
	//learnVariables()
	//looping()
	//ifElse()
	//switchStatements()
	//arrays()
	//maps()
	//usingRange()
	//usingFunctions()
	/*
		//Anonymous function
		fmt.Println(anonFunctions())
		nextInt := anonFunctions()
		fmt.Println(nextInt())
	*/
	recursion()
}

func doSomething() {
	fmt.Println(s)
}

func learnVariables() {
	var a string = "literal"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// Short for var f string = "short"
	f := "short"
	fmt.Println(f)

	var e int
	fmt.Println(e)

}

func looping() {
	i := 3

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 0; j <= 10; j++ {
		fmt.Println("J: ", j)
	}

	for {
		fmt.Println("Infinite looping")
		break
	}
}

func ifElse() {
	if 7%2 == 0 {
		//fmt.Println("7 % 2 = 0")
	} else if 7%2 < 0 {
		//fmt.Println("7 % 2 != 0")
	} else {
		fmt.Println("Else else elsy else ")
	}
}

func switchStatements() {
	i := 2
	switch i {
	case 1:
		fmt.Println("uno")
	case 2:
		fmt.Println("dos")
	}

	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("uno")
	default:
		fmt.Println("Current time ", t)
	}
}

func arrays() {
	var a [5]int
	fmt.Println("Declared empty arr", a)
	a[4] = 69
	fmt.Println("a[4] = 69 ", a)

	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("b := [5]int{1,2,3,4,5} results in ", b)

	fmt.Println("2 D array")
	var two2 [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			two2[i][j] = i + j
		}
	}
	fmt.Println(two2)
}

func maps() {
	m := make(map[string]int)

	m["Jim"] = 4
	m["Jimmy"] = 5
	m["Jimbo"] = 7

	fmt.Println("map: ", m)

	v := m["Jim"]
	fmt.Println(v)

	delete(m, "Jim")

	fmt.Println("After deleting m[\"Jim\"]", m)

	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println(n)
}

func usingRange() {
	nums := []int{1, 2, 3}
	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	//nums := []int{1, 2, 3}
	for i, num := range nums {
		if num == 3 {
			fmt.Println("Index: ", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}

func usingFunctions() {
	fmt.Println("Plus method ", plus(4, 6))
	fmt.Println("Plus method which is plusPlus(a, b, c int) ", plusPlus(3, 5, 7))
	a, b := doubleReturnVals(4, 12, 14, 6)
	fmt.Println("Plus method with 2 return vals doubleReturnVals(a, b, c, d int)(int, int)\n", a, b)

	//Variadic functions
	fmt.Println("Variadic Functions")
	fmt.Println("variadic method called func variadic(nums ...int)")

	nums := []int{1, 4, 6}
	variadicFunc(nums...)
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func doubleReturnVals(a, b, c, d int) (int, int) {
	return a + b, c + d
}

func variadicFunc(nums ...int) {
	fmt.Println(nums)
	total := 0
	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func anonFunctions() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func recursion() {
	fmt.Println(fact(3))
}

// Recursion function used in recursion
func fact(n int) int {
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}
