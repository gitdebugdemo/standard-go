package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestA1(t *testing.T) {
	s := `{"Name":"我的名字叫:name","Num":19}`
	var demo_obj map[string]interface{}
	// 注意这里需要用make，不能用new
	demo_obj = make(map[string]interface{})
	// json字符串编码成对象的核心函数:json.Unmarshal 。 传递指针进去，才能修改变量
	json.Unmarshal([]byte(s), &demo_obj)
	fmt.Printf("%+v\n", demo_obj)
	fmt.Printf("%t,%+v\n", demo_obj["Name"], demo_obj["Name"])
	// 连带类型也能解析出来
	fmt.Printf("%t,%+v\n", demo_obj["Num"], demo_obj["Num"])
}
