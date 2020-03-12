package test

import (
	"encoding/json"
	"fmt"
	"github.com/gitdebugdemo/standard-go/src/model/demo"
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
对象转成json
*/
func TestJSON_对象转json字符串_test(t *testing.T) {
	// 初始化类对象
	demo_obj := *new(demo.Demo)
	// 执行赋值
	demo_obj.
		SetNum(19).
		SetName("我的名字叫:name")
	//对象编码成json的核心函数 - json.Marshal
	json_string, err := json.Marshal(demo_obj)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
	fmt.Printf("%+v\n", string(json_string))
	assert.Equal(t, json_string, []byte(`{"Name":"我的名字叫:name","Num":19}`))
}
