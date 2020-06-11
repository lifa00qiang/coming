package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine())
	fmt.Println(runtime.NumCgoCall())
}

type TaskTest struct {
	say string
}

func (tt TaskTest) Run() ([]byte, error) {
	return []byte("hello task"), nil
}

//
//func Sum(opts ...interface{}) (r []byte,err error){
//	param := opts[0].([]interface{})
//	sum := param[0].(int) + param[1].(int)
//	bytes ,err:= json.Marshal(sum)
//
//	return bytes,err
//}

//go get github.com/go-task/task
//第一步;
//// 实现这个接口，或者传入一个func（） ，我们能狗通过这个接口的实现，来进行协程调度
//// 定义这个任务调用器的功能，
////第二步： 我们主题结构体，比如;任务，定时器，执行器
//// 抽象化任务调度结构体的属性
////第三步：  改造这个结构，实现协程通信，要求能狗知道协程是进行中还是完成了，
////实现接口
////第四步：  我们就添加中断机制，实现协程中断，，，（最好暂停）
////优化通信
