package test

import (
	"github.com/gitdebugdemo/standard-go/src/model/demo"
	"github.com/stretchr/testify/assert"
	"testing"
)

var demovar *demo.Demo

func TestAttr(t *testing.T) {
	demovar = &demo.Demo{Name: "test1", Num: 12}
	// 走了基类的 A流程1
	demovar.A我来定流程()
	// 走了业务类的 A流程1
	demovar.A流程1()
	// 断言:流程方法不能复用,导致一切复用作废
	assert.NotEqual(t, demovar.A我来定流程(), demovar.A流程1())
}
