package test

import (
	"fmt"
	"sync"
	"testing"
)

//[0 0 0 0 0 1 2 3]
func TestArray(t *testing.T) {
	s := make([]int, 5)
	s = append(s, 1, 2, 3)
	fmt.Println(s)

	mut := sync.RWMutex{}
	mut.RLock()
}
