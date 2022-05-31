package server

import (
	"fmt"
	"go-rpc/com/github/sheledon/connection"
	"log"
	"net"
)
const tcp = "tcp"
type RpcServer struct {
	address string
	conPool *connection.Pool
}
func NewRpcServer(address string) *RpcServer {
	return &RpcServer{
		address: address,
		conPool: connection.NewConnectionPool(),
	}
}
func (server *RpcServer) Listener(){
	listen, err := net.Listen(tcp, server.address)
	if err != nil {
		panic(fmt.Sprintf("fail start server, err: %v",err))
	} else {
		log.Printf("success start rpc server , listen : %s",listen.Addr())
	}
	for {
		conn, err2 := listen.Accept()
		if err2 != nil {
			log.Println("accept failed,err:", err)
			continue
		} else {
			log.Printf("accept connection,addr : %v \n",conn.RemoteAddr())
		}
		go server.process(conn)
	}
}
func (server *RpcServer) process(conn net.Conn){
	server.conPool.AddConnection(conn).ProcessRead()
}
