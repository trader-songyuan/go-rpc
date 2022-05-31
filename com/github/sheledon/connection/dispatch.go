package connection

import (
	"go-rpc/com/github/sheledon/constant"
	"go-rpc/com/github/sheledon/entity"
)
// 根据 message 类型进行分发
type DispatchHandler struct {}
func NewDisPatchHandler() DispatchHandler {
	return DispatchHandler{}
}
func (dh DispatchHandler) Read(context *ConnectContext) {
	message := context.GetAttr(constant.RPC_MESSAGE).(*entity.RpcMessage)
	strate := dh.getStrategy(message.MessageType)
	strate.Handle(context)
}
func (dh DispatchHandler) getStrategy(msgType byte) Strategy {
	if s,ok := strategyMap[msgType];ok {
		return s
	}
	return nil
}
var strategyMap = map[byte]Strategy{
	constant.RPC_REQUEST:  NewRpcRequestStrategy(),
	constant.RPC_RESPONSE: NewRpcResponseStrategy(),
}