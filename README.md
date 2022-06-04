# GO-RPC

使用go语言实现的简易RPC框架，采用protobuf作为序列化方式，zookeeper为注册中心；手动实现类似于java的动态代理机制简化rpc框架的使用。同时底层部分设计参考了Nettey的一些思路。

## 一. examples
#### 定义接口
```go
type GetByName func(stu *protoc.Student,i8 int8,i16 int16,i32 int32,name string) *protoc.Student
type GetByName2 func(ui8 uint8,ui16 uint16,f32 float32,f64 float64,b bool)
type GetByName3 func(isl []int32,stu *protoc.Student) []int32
```
#### 创建服务提供者
```go
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
```
#### 创建消费者
```go
// 服务消费者，定义struct，将接口定义为 struct 字段
type StuClient struct {
	GetByName
	GetByName2
	GetByName3
}
```
#### RPC
```go
func TestClient(t *testing.T)  {
	t.Run("server", func(t *testing.T) {
		address := "127.0.0.1:8080"
		server := server2.NewRpcServer(address) // 创建rpc server
		server.SetRegister(service.NewZkRegister([]string{"127.0.0.1:2181"})) // 定义zookeeper注册中心
		server.RegisterService("provider1","student",StuVo{}) // 手动注册服务
		server.Listener() // 开启rpc服务
	})
	t.Run("client", func(t *testing.T) {
		sc := StuClient{}
		client := NewRpcClient() // 创建rpc消费者客户端
		client.SetDiscovery(service.NewZkRegister([]string{"127.0.0.1:2181"})) // 定义zookeeper注册中心
		client.GenerateRpcProxy("student",&sc) // 动态代理，为结构体方法生成代理
		student := sc.GetByName(&protoc.Student{Name: "123", Age: 10, Money: 1000}, -1,10,29,"zhangsan") // 调用方法，进行rpc
		fmt.Println(student.Name)
		sc.GetByName2(10,19,10.9,18.9,false)
		time.Sleep(1000*time.Second)
	})
}
```
