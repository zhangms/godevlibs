package test

import (
	"fmt"
	"testing"
)

type student struct {
	Name string
	Age  int
}

// for range 分配的对象是一次性分配，后续值拷贝，所以地址相同
func TestRangeStudent(t *testing.T) {
	m := make(map[string]*student)
	students := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for i, stu := range students {
		fmt.Printf("%p, %p\n", &stu, &students[i])
		m[stu.Name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
