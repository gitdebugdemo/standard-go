package test

import (
	"fmt"
	"testing"
)

func TestA并发输出顺序很奇怪(t *testing.T) {
	for i := 1; i < 20; i++ {
		go fmt.Printf("%+v\n", i)
	}
}
