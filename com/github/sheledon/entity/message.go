package entity

import (
	"go-rpc/com/github/sheledon/property/constant"
	"google.golang.org/protobuf/proto"
)

type RpcMessage struct {
	MagicNumber byte
	Version byte
	Id int64
	MessageType byte
	ContentLength int64
	HeadLength int64
	Body proto.Message
}
func NewRpcMessage(msgType byte) *RpcMessage {
	return &RpcMessage{
		MagicNumber: constant.MAGIC_NUMBER,
		Version:     constant.VERSION,
		HeadLength:  constant.HEAD_LENGTH,
		MessageType: msgType,
	}
}
