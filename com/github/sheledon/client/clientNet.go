package client

import (
	"fmt"
	"go-rpc/com/github/sheledon/connection"
	"go-rpc/com/github/sheledon/constant"
	"go-rpc/com/github/sheledon/entity"
	"go-rpc/com/github/sheledon/entity/protoc"
	"go-rpc/com/github/sheledon/unprocess"
	"net"
	"sync"
)

type Client struct {
	lock       sync.Mutex
	pool       *connection.Pool
	uprFactory unprocess.UnProcessRequestFactory
}
func NewClient() *Client{
	return &Client{
		pool: connection.NewConnectionPool(),
	}
}
func (cl *Client) sendRpcRequest(request *protoc.RpcRequest) unprocess.Promise {
	promise := unprocess.NewPromise()
	unprocess.UprFactory.Set(request.Id, promise)
	rpcMessage := entity.NewRpcMessage(constant.RPC_REQUEST)
	rpcMessage.Body = request
	addr := "127.0.0.1:8080"
	conn := cl.GetConnection(addr)
	go conn.ProcessWrite(rpcMessage)
	return promise
}
func (cl *Client) GetConnection(addr string) *connection.RpcConnection {
	conn, err := cl.pool.GetConnection(addr)
	if err!=nil || conn==nil {
		cl.lock.Lock()
		defer cl.lock.Unlock()
		if conn,err = cl.pool.GetConnection(addr); err!=nil || conn==nil{
			con,err := cl.connect(addr) // todo 重试逻辑
			fmt.Println(err)
			cl.pool.AddConnection(con)
		}
	}
	rc, _ := cl.pool.GetConnection(addr)
	return rc
}
func (cl *Client) connect(addr string) (net.Conn,error){
	tcpAddr, err := net.ResolveTCPAddr(constant.NETWORK, addr)
	if err != nil {
		return nil,err
	}
	return net.DialTCP(constant.NETWORK,nil,tcpAddr)
}
