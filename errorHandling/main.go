package main

import "fmt"
import "errors"

func f1(arg int) (int, error) {
	if arg == 42 {
		return -1, errors.New("Can't work with 42")
	}

	return arg + 3, nil
}

type argError struct {
	arg  int
	prob string
}

func (e *argError) Error() string {
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}

func f2(arg int) (int, error) {
	if arg == 42 {
		return -1, &argError{arg, "cant work with that"}
	}

	return arg + 3, nil
}

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Printf("f1 failed: ", e)
		} else {
			fmt.Printf("f1 passed: ", r)
		}
	}

	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed: ", e)
		} else {
			fmt.Println("f2 passed: ", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
