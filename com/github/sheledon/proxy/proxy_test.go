package proxy

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"reflect"
	"strings"
	"testing"
)

type SayHello func(name string)
type GetNameById func(name string) (int,string)
type Person struct {
	name string
	age int
	SayHello
	GetNameById
}
func TestReflect(t *testing.T)  {
	person := &Person{name: "xiao",age: 10}
	InvocationProxy.NewProxyInstance(person, func(obj interface{}, method InvocationMethod, args []reflect.Value) []reflect.Value {
		fmt.Println("前置处理")
		//fmt.Println(args)
		fmt.Println("后置处理")
		return []reflect.Value{reflect.ValueOf(10),reflect.ValueOf("a")}
	})
	//proxy.SayHello("fdsa")
	//id, s := proxy.GetNameById("fdsa")
	//fmt.Println(id)
	//fmt.Println(s)
	person.GetNameById("trewtrew")
}
func TestUUID(t *testing.T)  {
	uid := uuid.NewV4()
	uids := strings.ReplaceAll(uid.String(), "-", "")
	fmt.Println(uids)
}
