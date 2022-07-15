package main

import (
	cool "awesomeProject/testpackage" //包起别名
	"fmt"
	"github.com/gin-gonic/gin"
	quiz "go_code/go_exercise/quizPart1"
	"net/http"
	"sync"
)

func main() {
	// StudyGin()
	//fmt.Println("helllo")
	//basic.Struct_define()

	//zeroMQ.Mytest()
	//basic.Fmain()

	quiz.Part01()

}

type Dog struct {
	name string
}

//idea 在结构体名字右击->生成->可以直接实现对应的接口
func (d Dog) Run() {
	//TODO implement me
	panic("implement me")
}

func (d Dog) Say() {
	//TODO implement me
	panic("implement me")
}

type Cat struct {
	name string
	sex  bool
}

func (c Cat) Run() {
	fmt.Println(c.name, "正在左右横跳")
}

//必须实现接口中的所有方法，无一遗漏，才算实现接口
func (c Cat) Say() {
	fmt.Println(c.name, "正在大吵大闹")
}

//接口的另一个作用，是用来解耦合
var L cool.Animal //先在这里进行定义

//定义一个与请求参数绑定的结构体，这种结构体必须要加注释
type PostParms struct {
	Name string `json:"name" binding:"required"` //这个是结构体注释,类似于验证器 validate，也可以自定义规则，通过自定义的标签
	Age  int    `json:"age"`
	Sex  bool   `json:"sex"`
}

func BaseTest() {
	c := Cat{
		name: "mery",
		sex:  false,
	}
	fmt.Println(cool.A)
	interfacetest(c)
	//cool.Main1()

	var wg sync.WaitGroup
	wg.Add(1)                //等待这个数字减到0 才继续执行
	go cool.Routine(&wg)     //引用，拿到原来的那个，这块只能传递引用，不能拿到值的话
	wg.Wait()                //等待协程执行完毕
	for i := 0; i < 3; i++ { //go中的循环
		fmt.Println(i)
	}

	c1 := make(chan int, 5) //chan后面的类型是存的东西的类型，数字的缓冲的个数，默认为0
	c2 := make(chan int)    //无缓冲区的使用
	c3 := make(chan int, 5)
	var readc <-chan int = c1  //一个c1 channel的只读变量
	var writec chan<- int = c1 //一个c1 channel的只写变量

	writec <- 2
	<-readc

	c1 <- 10086 //向channel中存入1
	//fmt.Println(<-c1)

	//通过打断点，体会执行步骤
	go func() {
		for i := 0; i < 5; i++ {
			c1 <- i
		}
	}()

	for i := 0; i < 3; i++ {
		fmt.Println(<-c1)
	}

	close(c1) //channel 是可以关闭的，关闭后仍然是可以从里面取，但是不可以再往里面放，只要不需要再往里面放，就应该关闭
	//已经关掉了，就不能再往里面放
	//go里面可以再定义其他函数,但是格式有区别
	go func() {
		c2 <- 1
	}()
	fmt.Println(<-c2)

	c3 <- 1
	c3 <- 2
	c3 <- 3
	c3 <- 4
	c3 <- 5
	//close(c3) //go使用for 一次性拿到channel中的值，在这之前必须close掉，好像也可以不用？？
	for v := range c1 {
		fmt.Println(v)
	}

	//如果有多个case满足，会随机执行或全部执行
	//select { //select可以保证channel 不会panic
	//case <-c: //select的每个case,只要case后面的内容可执行，就会执行
	//}

	cool.TestChannel()

}

func StudyGin() {
	router := gin.Default()                    //默认中间件  包含 Logger、Recovery 中间件
	router.GET("/ping", func(c *gin.Context) { //context就是gin的上下文
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	router.GET("/user13/:name", func(c *gin.Context) {
		//拿到传入的变量值
		name := c.Param("name") //这个是uri 中的地址
		//返回的信息
		c.String(http.StatusOK, "hello %s", name)
	})

	//这种是get带参数，而且参数是通过浏览器的那种方式传递过去的
	router.POST("/user/:name", func(c *gin.Context) {
		name := c.Param("name")                        //这个是从url中拿到的
		user := c.Query("user")                        //这个是传入的参数
		school := c.DefaultQuery("school", "zhongguo") //如果不带参数，可以指定默认值
		id := c.Query("id")
		pwd := c.Query("pwd")
		c.JSON(200, gin.H{
			"id":     id,
			"user":   user,
			"pwd":    pwd,
			"name":   name,
			"school": school,
		})
	})

	//获取get参数
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		//下面的简写想当于是 c.Request.URL.Query().Get("lastname") 的简写
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)

	})
	//通过form表单进行提交，现在已经很少用。基本都是用json进行提交
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous") // 此方法可以设置默认值

		//这个就是返回的信息
		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})

	//使用绑定，可以直接将请求的参数的映射到结构体里。不需要使用must band  而是should band
	router.POST("/test-bind", func(c *gin.Context) {
		var p PostParms //将结构体实例化

		err := c.ShouldBindJSON(&p) //需要拿一个error去接
		if err != nil {
			c.JSON(200, gin.H{
				"msg":  "报错",
				"data": gin.H{},
			})
		} else {
			c.JSON(200, gin.H{
				"msg":  "ok",
				"data": p,
				"test": p.Sex,
			}) //
		}

	})

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)

	}

	// Simple group: v2
	v2 := router.Group("/v2")
	{
		v2.POST("/login", loginEndpoint)

	}

	router.Run(":6666") // 可以指定端口号

}

func loginEndpoint(c *gin.Context) {
	fmt.Println("hello")

}

func interfacetest(a cool.Animal) {

	a.Run()
	//time.AfterFunc(10*time.Second, a.Say) //类似于js 中的settimeout
	//time.Sleep(3 * time.Second)  系统沉睡三秒中
	//time.After( ) this version will return a channel that will send a value after the given amount of time. This can be useful in combination with the select statement if you want a timeout while waiting on one or more channels.
	a.Say()

	//放到这里进行实例化并进行挂载
	L = a
}
