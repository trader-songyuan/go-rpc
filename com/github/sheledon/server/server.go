package server

import (
	"fmt"
	"go-rpc/com/github/sheledon/connection"
	"go-rpc/com/github/sheledon/constant"
	"log"
	"net"
)
const tcp = "tcp"
type RpcServer struct {
	address *net.TCPAddr
	conPool *connection.Pool
}
func NewRpcServer(address string) *RpcServer {
	tcpAddr, err := net.ResolveTCPAddr(constant.NETWORK, address)
	if err != nil {
		panic(err)
	}
	return &RpcServer{
		address: tcpAddr,
		conPool: connection.NewConnectionPool(),
	}
}
func (server *RpcServer) Listener(){
	listen, err := net.ListenTCP(constant.NETWORK,server.address)
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
		go server.conPool.AddConnection(conn)
	}
}