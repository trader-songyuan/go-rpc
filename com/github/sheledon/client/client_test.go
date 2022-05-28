package client

import (
	"fmt"
	server2 "go-rpc/com/github/sheledon/server"
	"go-rpc/com/github/sheledon/service"
	"testing"
)

type student struct {
	name string
	age int
}
func (s *student) SayHello()  {
	fmt.Println("hello")
}
type SayHello func()
type StudentC struct {
	name string
	age int
	SayHello
}
func TestTemp(t *testing.T) {
	t.Run("server", func(t *testing.T) {
		server := server2.NewRpcServer("127.0.0.1:8080")
		provier := service.ServiceProvier
		provier.RegisterService("student",&student{"hello",19})
		server.Listener()
	})
	t.Run("client", func(t *testing.T) {
		sc := &StudentC{name: "a",age: 1}
		RegisterRpcProxy("student",sc)
		sc.SayHello()
	})
}
