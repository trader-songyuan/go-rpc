package client

import (
	"fmt"
	uuid "github.com/satori/go.uuid"
	"go-rpc/com/github/sheledon/connection"
	"go-rpc/com/github/sheledon/entity"
	"go-rpc/com/github/sheledon/entity/protoc"
	"go-rpc/com/github/sheledon/loadbalance"
	"go-rpc/com/github/sheledon/property/constant"
	"go-rpc/com/github/sheledon/proxy"
	"go-rpc/com/github/sheledon/service"
	"go-rpc/com/github/sheledon/unprocess"
	"go-rpc/com/github/sheledon/utils"
	"net"
	"reflect"
	"strings"
	"sync"
)

type RpcClient struct {
	lock       sync.Mutex
	pool       *connection.Pool
	uprFactory unprocess.UnProcessRequestFactory
	discovery service.Discovery
	lb loadbalance.LoadBalance
}
func NewRpcClient() *RpcClient {
	return &RpcClient{
		pool: connection.NewConnectionPool(),
		lb: loadbalance.Random{},
	}
}
func (cl *RpcClient) SetDiscovery(discovery service.Discovery){
	cl.discovery = discovery
}
func (cl *RpcClient) sendRpcRequest(request *protoc.RpcRequest) unprocess.Promise {
	promise := unprocess.NewPromise()
	unprocess.UprFactory.Set(request.Id, promise)
	rpcMessage := entity.NewRpcMessage(constant.RPC_REQUEST)
	rpcMessage.Body = request
	providerList := cl.discovery.DiscoveryService(request.ServiceName)
	addr := cl.lb.Select(request.ServiceName,providerList)
	conn := cl.GetConnection(addr)
	go conn.ProcessWrite(rpcMessage)
	return promise
}
func (cl *RpcClient) GetConnection(addr string) *connection.RpcConnection {
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
func (cl *RpcClient) connect(addr string) (net.Conn,error){
	tcpAddr, err := net.ResolveTCPAddr(constant.NETWORK, addr)
	if err != nil {
		return nil,err
	}
	return net.DialTCP(constant.NETWORK,nil,tcpAddr)
}
func (cl *RpcClient) GenerateRpcProxy(serviceName string,target interface{}) {
	cl.wrapperFuncProxy(serviceName,target)
}
func (cl *RpcClient) wrapperFuncProxy(serviceName string,target interface{}) interface{}{
	return proxy.InvocationProxy.NewProxyInstance(target, func(obj interface{}, method proxy.InvocationMethod, args []reflect.Value) []reflect.Value {
		uid := uuid.NewV4()
		id := strings.ReplaceAll(uid.String(), "-", "")
		var rpcRequest protoc.RpcRequest
		if len(args) == 0{
			rpcRequest = protoc.RpcRequest{Id: id,ServiceName: serviceName,MethodName: method.Name}
		} else {
			params := make([]*protoc.RpcAny, len(args))
			for i,a:=range args{
				params[i] = utils.ValueTransferToRpcAny(a)
			}
			rpcRequest =protoc.RpcRequest{Id: id,ServiceName: serviceName,MethodName: method.Name,Params: params}
		}
		promise := cl.sendRpcRequest(&rpcRequest)
		pb, _ := promise.Get()
		response := pb.(*protoc.RpcResponse)
		res := make([]reflect.Value,len(response.Body))
		for i,b := range response.Body{
			res[i] = utils.RpcAnyToReflectValue(b)
		}
		return res
	});
}
