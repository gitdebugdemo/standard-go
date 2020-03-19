package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

type name struct {
	I int
}

var name_var name

// 创建包内的一个全局变量
var one sync.Once

// 单例模式函数
func Newname(i int) name {
	one.Do(func() {
		fmt.Printf("%+v:%+v\n", "被触发执行了", i)
		name_var = name{I: i}
	})
	return name_var
}

func TestA单实例(t *testing.T) {
	for i := 0; i <= 20; i++ {
		// 直接并发执行20个，但是结果是肯定。 返回的a都是一致
		go func() {
			a := Newname(i)
			fmt.Printf("%+v\n", a)
		}()
	}
	// 暂停是为了等打印信息出现再结束程序，否则可能出现程序结束了，打印还没执行出来。
	time.Sleep(time.Second)
}
