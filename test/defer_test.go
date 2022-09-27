package test

import "testing"

func TestDefer(t *testing.T) {
	deferCall()
}

func deferCall() {
	defer println("print1")
	defer println("print2")
	defer println("print3")
	panic("exception")
}
