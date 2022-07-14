package main //想在同一个项目中编译不同的包内文件，必须是main包，不能用自定义的包名

import "fmt"

func main() {
	//方法一：声明并初始化
	slice1 := []int{1, 2, 3}

	//方法二：声明并不初始化，还没有给slice分配空间  此时运行 slice2[0]=1 会报错
	var slice2 []int
	slice2 = make([]int, 6)
	slice2[0] = 1 //执行完make后才可以进行赋值

	//方法三：声明切片并同时开辟新的空间
	var slice3 []int = make([]int, 3)

	//方法四：方法三的简写方式,通过推导
	slice4 := make([]int, 3)
	slice4 = append(slice4, 6)

	fmt.Println(slice1, slice2, slice3, slice4)
	//v格式的话 可以打印任何变量的详细信息
	fmt.Printf("len=%d,slice=%v\n", len(slice1), slice1)
}
