package test

import (
	"fmt"
	"github.com/gitdebugdemo/standard-go/src/model/demo"
	"github.com/stretchr/testify/assert"
	"testing"
)

var demovar *demo.Demo

func TestAttr(t *testing.T) {
	demovar = &demo.Demo{Name: "test1", Num: 12}
	fmt.Printf("demovar.Name:\t%+v\n", demovar.Name)
	fmt.Printf("demovar.GetName():\t%+v\n", demovar.GetName())
	// 断言:直接打印属性和调用上级方法取属性,最终结果会是不一致的!!
	assert.NotEqual(t, demovar.Name, demovar.GetName())
}
