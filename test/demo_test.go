package test

import (
	"fmt"
	"github.com/gitdebugdemo/standard-go/src/model/demo"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
测试一个类对象的 get/set 操作
*/
func TestDemo_get_set(t *testing.T) {
	// 初始化类对象
	b := *new(demo.Demo)
	// 执行赋值
	b.
		SetNum(19).
		SetName("name")
	fmt.Printf("\n%T\n", b)
	fmt.Printf("\n%+v\n", b)
	assert.Equal(t, b.GetName(), "name")
	assert.Equal(t, b.GetNum(), int64(19))
}

/**
初始化一组对象
*/
func TestDemo_new_array(t *testing.T) {
	// 有一个属性是来自继承的:赋值的时候,需要一层一层声明下去
	b := []demo.Demo{
		demo.Demo{Demo_trait: demo.Demo_trait{Name: "ok"}, Num: 1},
		demo.Demo{Num: 2},
	}
	// 有一个属性是来自继承的:但是取值的时候,可以有继承的属性
	fmt.Printf("\n%+v\n", b[0].Name)
	fmt.Printf("\n%T\n", b)
	fmt.Printf("\n%+v\n", b)
	assert.Equal(t, b[0].Name, `ok`)
	assert.Equal(t, b[1].Num, int64(2))
}
