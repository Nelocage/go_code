package main //想在同一个项目中编译不同的包内文件，必须是main包，不能用自定义的包名

import "fmt"
import game_test "go_code/ebiten2dgame"

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

	//go的else位置很容易错
	//声明数组

	array := [5]string{"a", "b", "c", "d", "e"}
	fmt.Println(array[2])

	//数组循环
	for index, element := range array {
		fmt.Println("数组索引：%d,对应值：%s\n", index, element)
	}
	for _, element := range array {
		fmt.Println("元素为：%s\n", element)
	}

	m := map[string]int{"one": 1, "two": 2, "three": 3}
	m1 := map[string]int{}
	m1["one"] = 1
	m2 := make(map[string]int, 10 /*Initial Capacity*/)
	println(m, m2)
	//为什么不初始化 len ？（len都会赋值为0，但是map是没有办法做0值的， 只支持len访问长度，不支持cap访问容量 ）
	//但是make什么时候有用呢，比如切片是可以自增，每次在自增的时候都会分配新的内存空间，然后同时将数据进行拷贝，这样就会有相当的消耗。如果在初始化的时候可以将容量初始化到我们需要的大小，就可以避免这些，可以提高性能

	//map元素的访问，与其他语言的区别：key不存在的时候，仍会返回零值，不能通过返回nil来判断元素是否存在

	//在创建 map 的同时添加键值对，如果不想添加键值对，使用空大括号 {} 即可，要注意的是，大括号一定不能省略。
	nameAgeMap := map[string]int{"飞雪无情": 20}
	fmt.Println(nameAgeMap)

	game_test.GameMain()

}
