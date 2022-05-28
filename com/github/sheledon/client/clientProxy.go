package client

import (
	"github.com/satori/go.uuid"
	"go-rpc/com/github/sheledon/entity"
	"go-rpc/com/github/sheledon/proxy"
	"reflect"
	"strings"
)
var client = NewClient()

func RegisterRpcProxy(serviceName string,target interface{}) interface{}{
	return proxy.InvocationProxy.NewProxyInstance(target, func(obj interface{}, method proxy.InvocationMethod, args []reflect.Value) []reflect.Value {
		uid := uuid.NewV4()
		id := strings.ReplaceAll(uid.String(), "-", "")
		var rpcRequest entity.RpcRequest
		if len(args) == 0{
			rpcRequest = entity.NewRpcRequest(id,serviceName,method.Name)
		} else {
			rpcRequest = entity.NewRpcRequest(id, serviceName, method.Name, args)
		}
		promise := client.sendRpcRequest(rpcRequest)
		pb, _ := promise.Get()
		response := pb.(entity.RpcResponse)
		var res = make([]reflect.Value, len(response.Body))
		for i,ele := range response.Body{
			res[i] = reflect.ValueOf(ele)
		}
		return res
	});
}
