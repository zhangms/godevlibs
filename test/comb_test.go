package test

import (
	"fmt"
	"testing"
)

type people struct {
}

func (p *people) showA() {
	fmt.Println("showA")
	p.showB()
}

func (p *people) showB() {
	fmt.Println("showB")
}

type teacher struct {
	people
}

func (t *teacher) showB() {
	fmt.Println("teacher showB")
}

func TestComb(t *testing.T) {
	teach := &teacher{}
	teach.showA()
}
