package client

import (
	"github.com/satori/go.uuid"
	"go-rpc/com/github/sheledon/entity/protoc"
	"go-rpc/com/github/sheledon/proxy"
	"go-rpc/com/github/sheledon/utils"
	"reflect"
	"strings"
)
var client = NewClient()

func RegisterRpcProxy(serviceName string,target interface{}) interface{}{
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
		promise := client.sendRpcRequest(&rpcRequest)
		pb, _ := promise.Get()
		response := pb.(*protoc.RpcResponse)
		res := make([]reflect.Value,len(response.Body))
		for i,b := range response.Body{
			res[i] = utils.RpcAnyToReflectValue(b)
		}
		return res
	});
}

