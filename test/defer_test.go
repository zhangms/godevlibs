package test

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	deferCall()
}

func deferCall() {
	defer println("print1")
	defer println("print2")
	defer println("print3")
	panic("exception")
}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

//1. 值传递，defer后改变不会影响defer内实参
//2. defer 参数中有function的，会先求值
//10, 1, 2, 3
//20, 0, 2, 2
//2 , 0, 2, 2
//1 , 1, 3, 4

func TestDefer2(t *testing.T) {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func TestDefer3(t *testing.T) {
	fmt.Println("--->", deferChangeReturnValue(10))
	fmt.Println("--->", deferChangeReturnValue(30))
	fmt.Println("--->", deferChangeReturnValue(60))
}

func deferChangeReturnValue(input int) (output int) {
	defer func() {
		if output > 100 {
			output = 100
		}
	}()
	output = input * 2
	return
}

func TestDefer4(t *testing.T) {
	fmt.Println("--->", deferCantChangeReturnValue(10))
	fmt.Println("--->", deferCantChangeReturnValue(30))
	fmt.Println("--->", deferCantChangeReturnValue(60))
}

func deferCantChangeReturnValue(input int) int {
	output := input * 2
	defer func() {
		if output > 100 {
			output = 100
		}
	}()
	return output
}
