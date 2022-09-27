package test

import (
	"sync"
	"testing"
)

func ptr(str string) *string {
	return &str
}

func TestPool(t *testing.T) {

	pool := sync.Pool{
		New: func() any {
			return ptr("hello world")
		},
	}
	str := pool.Get().(*string)
	println(str)
	pool.Put(str)
	str2 := pool.Get().(*string)
	println(str2)
}
