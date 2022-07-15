package main

import (
	"fmt"
	"reflect"
)

/*
断言：
把一个接口类型指定为他的原始类型,从而获得结构体的相应的字段以及原来结构体上定义的方法
因为接口只能调用接口定义的方法，是无法获得结构体的字段和方法了


反射：
官方说法：在编译时不知道类型的情况下，可更新变量 运行时查看值 调用方法以及直接对他们的布局进行操作的机制，成为反射
通俗一点：可以知道本数据的原始数据类型和数据内容 方法等，并且可以进行一定的操作

为什么要用反射：
通过接口或者其他的方法接收到了类型不固定的数据的时候，需要写太多的switch case 断言代码
此时代码不灵活且通用性差，反射这时候就可以无视类型，改变原数据结构中的数据

反射包 reflect
reflect.Valueof() 获取输入参数接口中的数据的值
reflect.TypeOf() 动态获取输入参数接口中的值的类型
reflect.TypeOf().Kind() 用来判断类型
reflect.Valueof().Fidle(int) 用来获取值
*/

type User struct {
	Name string
	Age  int
	Sex  bool
}

type Student struct {
	User
}

//给结构体定义方法
func (u User) SayName(name string) {
	fmt.Println("名字是", name)
	//a := [...]int{1, 2, 3, 4}

}

func check(v interface{}) {
	//switch和断言，联合使用
	switch v.(type) { //这个是固定的写法
	case User:
		a := v.(User).Name
		fmt.Println("传入的是一个user,名字是", a)
		break
	case Student:
		fmt.Println("传入的是一个student")
		break

	}

}

func AssertMain() {
	u := User{
		Name: "ii",
		Sex:  false,
		Age:  10,
	}

	TestReflect(u)

}

func TestReflect(inner interface{}) {
	t := reflect.TypeOf(inner)
	v := reflect.ValueOf(inner)
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(v.Field(i)) //field 里面需要传递个int

	}

}
