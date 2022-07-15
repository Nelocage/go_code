package main

import (
	"fmt"
	"sync"
)

//goroutine 和channel的示例代码
//类似与其他语言的协程，或者js中的异步

//goroutine 与主程序隔离，自己跑在另外的地方
//channel  各协程与主程序通信的机制

//在调用那个一个方法的前面加上go 就是goroutine 他会让方法异步的执行 相当于协程

func Routine(wg *sync.WaitGroup) {
	fmt.Println("跑步中.....")
	wg.Done() //运行完之后，执行down
}

/*

协程管理器
	var wg.sync.WaitGroup
	wg.Add()
	wg.Done()
	wg.Wait()
*/
//

//channel 是goroutine之间的通讯桥梁
//定义chan 分为五种
//可读可取 c:make(chan int )
//可读 var readChan <-chan int = c
//可取 var setChan chan<- int = c
//有缓冲 c := make(chan int, 5)
//无缓冲 c := make(chan int )

//channel开启后是可以close的，当不再需要并且已经set完成的时候，就需要 close它
//注意：如果用到了range ,则必须在range 前就给他关闭

//如何使用channel来进行goroutine通信
func SetChan(writec chan<- int) {
	for i := 0; i < 10; i++ {
		fmt.Println("我在set函数里")
		writec <- i
	}

}

func GetChan(readc <-chan int) {
	for i := 0; i < 10; i++ {

		//channel 一定要加<-，不可以直接打印变量，否则只会拿到地址
		fmt.Println("我在get函数里收到set的信息是", <-readc) //直接打印会显示地址

	}

}

func TestChannel() {
	c := make(chan int)
	var readc <-chan int = c
	var writec chan<- int = c
	go SetChan(writec)
	GetChan(readc)

}
