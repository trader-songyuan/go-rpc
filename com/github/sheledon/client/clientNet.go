package client

import (
	"fmt"
	"go-rpc/com/github/sheledon/connection"
	"go-rpc/com/github/sheledon/constant"
	"go-rpc/com/github/sheledon/entity"
	"net"
	"sync"
)

var unProcessRequestFactory = newUnProcessRequestFactory()

type Client struct {
	lock sync.Mutex
	pool *connection.Pool
}

func NewClient() *Client{
	return &Client{
		pool: connection.NewConnectionPool(),
	}
}
func (cl *Client) sendRpcRequest(request entity.RpcRequest) Promise {
	promise := NewPromise()
	unProcessRequestFactory.Set(request.Id, promise)
	rpcMessage := entity.NewRpcMessage(constant.RPC_REQUEST)
	rpcMessage.Body = request
	addr := "127.0.0.1:8080"
	conn := cl.GetConnection(addr)
	conn.SendMsg(rpcMessage)
	return promise
}
func (cl *Client) GetConnection(addr string) *connection.RpcConnection{
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
	return net.Dial("tcp", addr)
}
