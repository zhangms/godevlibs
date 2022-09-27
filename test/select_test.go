package test

import (
	"fmt"
	"runtime"
	"testing"
)

func TestSelect(t *testing.T) {
	runtime.GOMAXPROCS(1)
	ic := make(chan int, 1)
	sc := make(chan string, 1)
	ic <- 1
	sc <- "hello"
	select {
	case v := <-ic:
		fmt.Println(v)
	case v := <-sc:
		panic(v)
	}

}
