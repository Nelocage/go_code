package main

import "fmt"

var A string = "hello"

type Animal interface { //接口和里面的字段名，同样要首字母大写，以确保可见性
	Run()
	Say()
}

func Main1() {
	//变量定义 单引号，只能用一个字符，与php不同
	var name1 string
	name1 = "我是声明变量的第一种方式"

	var name2 = "我是声明变量的第二种方式"

	name3 := "我是声明变量的第三种方式"

	fmt.Println(name1, name2, name3)

	//多变量声明的第一种方式，以逗号分隔，声明与赋值分开，若不赋值，存在默认值
	var add1, add2, add3 string
	add1, add2, add3 = "china", "shanghai", "jingan"

	fmt.Println(add1, add2, add3)

	//第二种，直接赋值，下面的变量类型可以是不同的类型
	var phone1, phone2, phone3 = 1, false, "hello"
	if !phone2 {
		fmt.Println(phone1, phone3)
	}

	//第三种，集合类型
	var (
		age1 int
		age2 string
	)
	println(age1, age2)

	//显式类型定义：   常量赋值后不使用，不会报错
	const b1 string = "abc"
	//隐式类型定义：
	const b2 = "abc"

	//常量可以用作枚举
	const (
		Unknown = 0
		Female  = 1
		Male    = 2
	)

	const (
		x uint16 = 16
		y
		s = "abc"
		z
	)
	fmt.Printf("%T,%v\n", y, y) //T为打印其类型，v为打印其值
	fmt.Printf("%T,%v\n", z, z)

}
