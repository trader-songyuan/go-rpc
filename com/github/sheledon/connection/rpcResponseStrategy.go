package connection

import (
	"go-rpc/com/github/sheledon/entity"
	"go-rpc/com/github/sheledon/entity/protoc"
	"go-rpc/com/github/sheledon/property/constant"
	"go-rpc/com/github/sheledon/unprocess"
)

type RpcResponseStrategy struct {}
func NewRpcResponseStrategy() RpcResponseStrategy {
	return RpcResponseStrategy{}
}
func (s RpcResponseStrategy) Handle(ctx *ConnectContext)  {
	message := ctx.GetAttr(constant.RPC_MESSAGE).(*entity.RpcMessage)
	response := message.Body.(*protoc.RpcResponse)
	promise:= unprocess.UprFactory.Get(response.Id)
	if promise != nil {
		if response.Code == constant.INVOKE_ERROR {
			promise.CompleteFailure(response.Err)
		} else if response.Code == constant.SUCCESS {
			promise.Complete(response)
		}
	}
}
