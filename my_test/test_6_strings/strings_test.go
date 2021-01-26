package test_6_strings

import (
	"fmt"
	"strings"
	"testing"
)

const url  = "http://localhost:9050/v1/testDemo"

func TestRemoveHttp(t *testing.T){
	fmt.Println(url)
	a := strings.TrimLeft(url,"http://") // 去掉前缀 ：http://
	fmt.Println(a)
	urlList := strings.Split(a,"/") // 字符串分割获取路由分组
	fmt.Println(urlList)
	return
}

func TestTrimLeft(t *testing.T){
	fmt.Println(url)
	a := strings.TrimLeft(url,"https://")
	fmt.Println("=====>",a) //这怎么就删了？  A:原来是只要这边出现了就删了   不能随便用啊。。。
	b := strings.TrimLeft(url,"httpASD://")
	fmt.Println("=====>",b)
	c := strings.TrimLeft(url,"httpASD://AAAAAA")
	fmt.Println("=====>",c)
	d := strings.TrimLeft(url,"cal50oHlaCLlHhTtTtPp://")
	fmt.Println("=====>",d)
	return
}