package ch11

import (
	"testing"
)

func TestFakeFunc(t *testing.T) {
	var data string
	saved := realFunc //保存原来的realfunc
	defer func() { realFunc = saved }()
	realFunc = func(d string) {
		data = d + "demo"
	}
	CheckInfo("demo")
	if data != "demo" {
		t.Errorf(`data=%s want "demo"`, data)
	}
}
