package sheledon

import (
	"fmt"
	"go-rpc/com/github/sheledon/connection"
	"go-rpc/com/github/sheledon/constant"
	"go-rpc/com/github/sheledon/entity"
	server2 "go-rpc/com/github/sheledon/server"
	"go-rpc/com/github/sheledon/service"
	"net"
	"testing"
	"time"
)

type User struct {
	name string
	age int
}
func (u User) Hello() {
	fmt.Printf("name : %s, age: %d",u.name,u.age)
}
func TestRpc(t *testing.T) {
	t.Run("server", func(t *testing.T) {
		server := server2.NewRpcServer("127.0.0.1:8080")
		provier := service.ServiceProvier
		provier.RegisterService("user",User{"hello",19})
		server.Listener()
	})
	t.Run("client", func(t *testing.T) {
		conn, err := net.Dial("tcp", "127.0.0.1:8080")
		if err != nil{
			t.Fatal("fail connect")
		}
		defer conn.Close()
		request := entity.NewRpcRequest("123","user","Hello")
		encoder := connection.NewEncodeHandler()
		context := connection.NewConnectContext(connection.NewRpcConnection(conn))
		rpcMessage := entity.NewRpcMessage(constant.RPC_REQUEST)
		rpcMessage.Body = request
		encoder.Write(context,rpcMessage)
		time.Sleep(10000*time.Second)
	})
}

