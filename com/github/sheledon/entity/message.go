package entity

import (
	"go-rpc/com/github/sheledon/constant"
)

type RpcMessage struct {
	MagicNumber byte
	Version byte
	Id int64
	MessageType byte
	ContentLength int64
	HeadLength int64
	Body interface{}
}

func NewDefaultRpcMessage() *RpcMessage {
	return &RpcMessage{
		MagicNumber: constant.MAGIC_NUMBER,
		Version:     constant.VERSION,
		HeadLength:  constant.HEAD_LENGTH,
	}
}
type RpcRequest struct {
	Id string
	ServiceName string
	MethodName string
	Params []interface{}
}
func NewRpcRequest(id,serviceName,methodName string,params ...interface{}) RpcRequest{
	return RpcRequest{
		id,serviceName,methodName,params,
	}
}
type RpcResponse struct {
	Id string
	ServiceName string
	Body interface{}
}

func NewRpcMessage(msgType byte) *RpcMessage {
	return &RpcMessage{
		MagicNumber: constant.MAGIC_NUMBER,
		Version: constant.VERSION,
		HeadLength: constant.HEAD_LENGTH,
		MessageType: msgType,
	}
}
