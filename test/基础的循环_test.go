package test

import (
	"fmt"
	"testing"
)

/**
for - range 循环
*/
func TestA0(t *testing.T) {
	var a_obj []string
	a_obj = []string{"a", "b", "c"}
	// 最基础的循环 - 带索引
	for index, value := range a_obj {
		fmt.Printf("%+v,%+v\n", index, value)
	}
}

/**
for - range 循环
*/
func TestA1(t *testing.T) {
	var a_obj []string
	a_obj = []string{"a", "b", "c"}
	// 最基础的循环 - 不需要索引
	for _, value := range a_obj {
		fmt.Printf("%+v\n", value)
	}
}

/**
for 基础循环 - 循环索引来取值，比较啰嗦，影响编码心情
*/
func TestA2(t *testing.T) {
	var a_obj []string
	a_obj = []string{"a", "b", "c"}
	for i := 0; i < len(a_obj); i++ {
		fmt.Printf("%+v\n", a_obj[i])
	}
}
