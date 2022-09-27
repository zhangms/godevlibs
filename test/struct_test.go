package test

import (
	"fmt"
	"reflect"
	"testing"
)

type People interface {
	Show()
}

type Student struct{}

func (stu *Student) Show() {
	println("------stu.Show------>", stu)
}

func live() People {
	var stu *Student
	println("1------>", stu)
	return stu
}

//虽然stu是nil，但是返回时将其赋值给People类型时，其实是保存了类似(stu, *Student)这样的数据--对应接口实际的值和类型，故而不为nil
func TestPeople(t *testing.T) {
	ll := live()
	println(ll)
	fmt.Println(reflect.ValueOf(ll).IsNil())
	if ll == nil {
		fmt.Println("AAAAAAA")
	} else {
		fmt.Println("BBBBBBB")
		fmt.Println(reflect.TypeOf(ll), ll)
		var stu *Student
		fmt.Println(reflect.TypeOf(stu), stu)
	}
}

type Parent struct {
	Age int
}

func (p *Parent) SetAge(age int) {
	p.Age = age
}

type Child struct {
	Parent
}

type Child2 struct {
	*Parent
}

func TestParentAndChild(t *testing.T) {
	parent := Parent{10}
	child := Child{parent}
	child2 := &Child2{&parent}
	parent.SetAge(20)
	child.SetAge(30)
	println("parent age:", parent.Age, " child age:", child.Age)
	println("parent age:", parent.Age, " child2 age:", child2.Age)
	fmt.Printf("%#v", child)
}

type Animal interface {
	Speak(str string) string
}

type Dog struct {
}

func (d Dog) Speak(str string) string {
	ret := "dog speak : " + str
	fmt.Println(ret)
	return ret
}

type Cat struct {
}

func (c *Cat) Speak(str string) (result string) {
	result = "cat speak : " + str
	fmt.Println(result)
	return
}

func TestAnimal(t *testing.T) {
	//var dog Animal = Dog{} //可编译
	var dog Animal = &Dog{} //依然可编译
	dog.Speak("hello world")

	//var cat Animal = Cat{} //编译不过
	var cat Animal = &Cat{}
	cat.Speak("hello world")
}
