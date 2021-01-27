package test_6_strings

import (
	"fmt"
	"reflect"
	"testing"
)

func TestRangeIt(t *testing.T) {
	var a = "hello, world!"
	{
		for _, v := range a {
			fmt.Print(v, " ")
			fmt.Println(reflect.TypeOf(v).String())
		}
	}
	fmt.Println("=========================>")
	{
		for i := 0; i < len(a); i++ {
			fmt.Print(a[i], " ")
			fmt.Println(reflect.TypeOf(a[i]))
		}
	}
	fmt.Println("=========================>")
	{
		b := []rune(a)
		for _, v := range b {
			fmt.Println(v)
		}
	}
	fmt.Println("=========================>")
	{
		v4 := `
		床前明月光,
			疑似地上霜.
				举着望明月,
			低头思故乡.
		`
		v6 := []rune(v4)
		v7 := "故"
		for k, v := range v6 {
			if string(v) == v7 {
				fmt.Printf("找到字符---\"%s\",\n其索引为%d\n", v7, k)
				fmt.Printf("%d--%c--%T\n", k, v, v)
			}
		}
	}
	fmt.Println("=========================>")
	{
		c := "这是什么鬼东西？?????"
		for i, v := range c { // 这是编码过的
			fmt.Printf("No.%02d   =====>   %c ===>", i, v)
			fmt.Println(v)
		}
	}
	fmt.Println("=========================>")
	{
		c := "这是什么鬼东西？?????"
		for i := 0; i < len(c); i++ {
			fmt.Printf("No.%02d   =====>   %c ===>", i, c[i])
			fmt.Println(c[i])
		}
	}
	return
}
