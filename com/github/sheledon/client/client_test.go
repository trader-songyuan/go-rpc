package client

import (
	"fmt"
	"go-rpc/com/github/sheledon/entity/protoc"
	server2 "go-rpc/com/github/sheledon/server"
	"go-rpc/com/github/sheledon/service"
	"testing"
	"time"
)

// 服务接口传递参数，对于struct必须传递指针，返回值也是如此


// 接口
type GetByName func(stu *protoc.Student,i8 int8,i16 int16,i32 int32,name string) *protoc.Student
type GetByName2 func(ui8 uint8,ui16 uint16,f32 float32,f64 float64,b bool)
type GetByName3 func(isl []int32,stu *protoc.Student) []int32
// 服务提供者，定义结构体，实现接口
type StuVo struct {}
func (sv StuVo) GetByName(stu *protoc.Student,i8 int8,i16 int16,i32 int32,name string) *protoc.Student{
	fmt.Println(stu.Name,"  ",stu.Age,"  ",stu.Money)
	fmt.Printf("i8: %d, i16: %d, i32: %d\n",i8,i16,i32)
	return &protoc.Student{
		Name: name,
		Age: 100,
		Money: 100000,
	}
}
func (sv StuVo) GetByName2(ui8 uint8,ui16 uint16,f32 float32,f64 float64,b bool){
	fmt.Printf("ui8: %d, ui16: %d\n",ui8,ui16)
	fmt.Printf("f32: %0.5f, f64: %0.5f\n",f32,f64)
	fmt.Printf("bool: %v\n",b)
}
func (sv StuVo) GetByName3(isl []int32,stu *protoc.Student) []int32{
	fmt.Println(isl)
	fmt.Println(stu.Name,"  ",stu.Age,"  ",stu.Money)
	return []int32{1,2,3,4,5}
}
// 服务消费者，定义struct，将接口定义为 struct 字段
type StuClient struct {
	GetByName
	GetByName2
	GetByName3
}

func TestClient(t *testing.T)  {
	t.Run("server", func(t *testing.T) {
		server := server2.NewRpcServer("127.0.0.1:8080")
		provier := service.ServiceProvier
		provier.RegisterService("student",StuVo{})
		server.Listener()
	})
	t.Run("client", func(t *testing.T) {
		sc := StuClient{}
		RegisterRpcProxy("student",&sc)
		student := sc.GetByName(&protoc.Student{Name: "123", Age: 10, Money: 1000}, -1,10,29,"zhangsan")
		fmt.Println(student.Name)
		sc.GetByName2(10,19,10.9,18.9,false)
		time.Sleep(1000*time.Second)
	})
}
