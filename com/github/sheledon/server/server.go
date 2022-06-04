package server

import (
	"fmt"
	"go-rpc/com/github/sheledon/connection"
	"go-rpc/com/github/sheledon/property/constant"
	"go-rpc/com/github/sheledon/service"
	"log"
	"net"
)
type RpcServer struct {
	address *net.TCPAddr
	conPool *connection.Pool
	provider *service.Provider
}
func NewRpcServer(address string) *RpcServer {
	tcpAddr, err := net.ResolveTCPAddr(constant.NETWORK, address)
	if err != nil {
		panic(err)
	}
	return &RpcServer{
		address: tcpAddr,
		conPool: connection.NewConnectionPool(),
		provider: service.GetProvider(address),
	}
}
func (server *RpcServer) SetRegister(r *service.ZkRegister) {
	server.provider.SetRegister(r)
}
func (server *RpcServer) RegisterService(providerName,serviceName string,target interface{})  {
	server.provider.RegisterService(providerName,serviceName,target)
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