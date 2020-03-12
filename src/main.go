package main

import (
	"fmt"
	appdemo "github.com/gitdebugdemo/standard-go/src/app/demo"
	demo "github.com/gitdebugdemo/standard-go/src/model/demo"
)

//声明2个相同名字的类的
var (
	demo_obj    *demo.Demo
	appdemo_obj *appdemo.Demo
)

/**
main
	1.展示了如果存在2个相同package名字,但是在不同目录下,也是可以不冲突,采用别名调用起来(php就没有这么蛋疼的问题)
	2.代码复用下.一个完整的go类型的类,由3个文件组成.和php编写方式一致
*/
func main() {
	fmt.Printf("\n=================\n")
	//第一个package.类的操作
	{
		demo_obj = new(demo.Demo)
		appdemo_obj = new(appdemo.Demo)
		demo_obj.Name = "你好啊,i am demo_obj"
		demo_obj.Num = 20
		demo_obj.Invoke()
		fmt.Printf("demo_obj:%T,%+v\nSay_world:%v\n", demo_obj, *demo_obj, demo_obj.Say_world())
	}
	fmt.Printf("\n=================\n")
	//第二个package.类的操作
	{
		appdemo_obj.Name = "i am appdemo_obj"
		fmt.Printf("appdemo_obj:相同包名字冲突的调用方式:%T,%+v\n", appdemo_obj, appdemo_obj)
	}
	fmt.Printf("\n=================\n")
	//实现例子1
	{
		参数是基类_传入是子类对象(demo_obj)
	}
	fmt.Printf("\n=================\n")
	//实现例子2 - 错误 - 因为传入了一份拷贝,而不是引用
	{
		fmt.Printf("[❌错误例子_因为传递是值]打印出修改前的名字:%+v\n", demo_obj.GetName())
		错误例子_因为传递是值(*demo_obj)
		fmt.Printf("[❌错误例子_因为传递是值]打印出修改后的名字:%+v\n", demo_obj.GetName())
	}
	//实现例子2 - 正确 - 因为传入引用
	{
		fmt.Printf("[✅正确例子_因为传递的是地址]打印出修改前的名字:%+v\n", demo_obj.GetName())
		正确例子_因为传递的是地址(demo_obj)
		fmt.Printf("[✅正确例子_因为传递的是地址]打印出修改后的名字:%+v\n", demo_obj.GetName())
	}
}

/**
实现例子:
1:函数接受的参数是基类,传递进具体子类对象
*/
func 参数是基类_传入是子类对象(obj demo.Demo_interface) {
	obj.Invoke()
	fmt.Printf("[接受的参数是基类,传递进具体子类对象]打印出参数:%+v\n", obj)
}

/**
实现例子:
2:函数接受的参数是类,可以改变传递的类对象的值
*/
func 错误例子_因为传递是值(obj demo.Demo) {
	obj.SetName("修改了你的名字,demo_obj")
}

/**
实现例子:
2:函数接受的参数是类,可以改变传递的类对象的值
*/
func 正确例子_因为传递的是地址(obj *demo.Demo) {
	obj.SetName("修改了你的名字,demo_obj")
}
