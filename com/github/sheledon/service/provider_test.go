package service

import (
	"fmt"
	"reflect"
	"testing"
)

type student struct {
	name string
}
func (s student) GoSchool()  {
	fmt.Printf("go to school: %s \n",s.name)
}
func (s student) GetHello(message string) string {
	return fmt.Sprintf("stu: %s say hello msg %s \n",s.name,message)
}
type myString string
func TestNewProvider(t *testing.T) {
	providers := NewProvider()
	stu := student{name: "hututu"}
	const sname = "student"
	err := providers.registerService(sname,stu)
	shutdownTest(t,err)
	service, err := providers.getService(sname)
	shutdownTest(t,err)
	value :=reflect.ValueOf(service)
	method := value.MethodByName("GoSchool")
	method.Call([]reflect.Value{})
	typeof := reflect.ValueOf(service)
	method2 := typeof.MethodByName("GetHello")
	msg := "fdasfdsa"
	result := method2.Call([]reflect.Value{reflect.ValueOf(msg)})
	for _,v := range result{
		fmt.Println(v.String())
	}
}
func shutdownTest(t *testing.T,err error)  {
	t.Helper()
	if err!=nil {
		t.Fatalf("fail register, err: %v",err)
	}
}

