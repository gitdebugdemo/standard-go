package test

import (
	"encoding/json"
	"fmt"
	"github.com/gitdebugdemo/standard-go/src/model/demo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJson_字符串还原到类属性上(t *testing.T) {
	s := `{"Name":"我的名字叫:name","Num":19}`
	var demo_obj *demo.Demo
	demo_obj = new(demo.Demo)
	// json字符串编码成对象的核心函数:json.Unmarshal
	json.Unmarshal([]byte(s), demo_obj)
	fmt.Printf("%+v\n", demo_obj)
	fmt.Printf("%+v\n", demo_obj.GetName())
	fmt.Printf("%+v\n", demo_obj.GetNum())
	assert.Equal(t, demo_obj.GetName(), "我的名字叫:name")
	assert.Equal(t, demo_obj.GetNum(), int64(19))
}
