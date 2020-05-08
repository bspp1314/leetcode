package main

import (
	"fmt"
	"reflect"
)


type User struct {
	Name   string `json:"name,omitempty"` //这引号里面的就是tag
	Passwd string `json:"passwd,omitempty"`
}

func main() {
	user := &User{"chronos", "pass"}
	s := reflect.TypeOf(user).Elem() //通过反射获取type定义
	for i := 0; i < s.NumField(); i++ {
		fmt.Println(s.Field(i).Tag.Get("json")) //将tag输出出来
	}

}
