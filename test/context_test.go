package test

import (
	"context"
	"fmt"
	"testing"
)

func TestContext(t *testing.T) {

	ctx := context.Background()
	ctx2 := context.WithValue(ctx, "hello", "world")
	fmt.Println(ctx2.Value("hello"))

	ctx3, cancel := context.WithCancel(ctx2)
	fmt.Println(ctx3.Value("hello"))
	cancel()

	for {
		select {
		case done := <-ctx2.Done():
			fmt.Println("done", done)
			return
		default:
			running()
		}
	}
}

func running() {
}
