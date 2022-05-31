package connection

import (
	"fmt"
	"go-rpc/com/github/sheledon/constant"
	"go-rpc/com/github/sheledon/entity"
	"go-rpc/com/github/sheledon/entity/protoc"
	"go-rpc/com/github/sheledon/service"
	utils "go-rpc/com/github/sheledon/utils"
	"reflect"
)

type RpcRequestStrategy struct {
	provider *service.Provider
}
func NewRpcRequestStrategy() RpcRequestStrategy {
	return RpcRequestStrategy{
		provider: service.ServiceProvier,
	}
}
func (s RpcRequestStrategy) Handle(ctx *ConnectContext) {
	message := ctx.GetAttr(constant.RPC_MESSAGE).(*entity.RpcMessage)
	request := message.Body.(*protoc.RpcRequest)
	svice, err := s.provider.GetService(request.ServiceName)
	if err != nil {
		//响应错误
		return
	}
	instance := reflect.ValueOf(svice)
	method := instance.MethodByName(request.MethodName)
	vparams := make([]reflect.Value,len(request.Params))
	for i,p := range request.Params{
		//转换为value类型
		vparams[i] = utils.RpcAnyToReflectValue(p)
	}
	res,err := s.invokeMethod(method, vparams)
	//响应请求
	var rpcResponse protoc.RpcResponse
	if err != nil {
		panic(err)
		rpcResponse = createRpcResponse(request,nil,err.Error(),constant.INVOKE_ERROR)
	} else {
		ares := make([]*protoc.RpcAny,len(res))
		for i,r := range res{
			ares[i] = utils.ValueTransferToRpcAny(r)
		}
		rpcResponse = createRpcResponse(request,ares,"nil",constant.SUCCESS)
	}
	message = entity.NewRpcMessage(constant.RPC_RESPONSE)
	message.Body = &rpcResponse
	ctx.SendRequest(message)
}
func (s *RpcRequestStrategy) invokeMethod(method reflect.Value,params []reflect.Value) (res []reflect.Value,rerr error){
	defer func() {
		if err := recover();err != nil {
			rerr = fmt.Errorf("invoke fail : %v",err)
		}
	}()
	res = method.Call(params)
	return
}
func createRpcResponse(request *protoc.RpcRequest,body []*protoc.RpcAny,err string,status int32)  protoc.RpcResponse{
	return protoc.RpcResponse{
		Id: request.Id,
		ServiceName: request.ServiceName,
		Code: status,
		Body: body,
		Err: err,
	}
}
