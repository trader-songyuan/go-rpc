package connection

import (
	"fmt"
	"go-rpc/com/github/sheledon/constant"
	"go-rpc/com/github/sheledon/entity"
	"go-rpc/com/github/sheledon/entity/protoc"
	"google.golang.org/protobuf/proto"
)
/**
	入站解码器 : 大端序列字节流 ---> rpcMessage
 */
type DecodeHandler struct {
}
func NewDecodeHandler() *DecodeHandler {
	return &DecodeHandler{}
}
func (h DecodeHandler) Read(context *ConnectContext) {
	magicNumber := context.ReadByte()
	version := context.ReadByte()
	id := context.ReadInt64()
	msgType := context.ReadByte()
	contentLength:= context.ReadInt64()
	headLength:= context.ReadInt64()
	bodyLen := contentLength - headLength
	rpcMessage := CreateRpcMessage(id, contentLength, headLength, msgType)
	setMsgBodyByType(rpcMessage)
	checkMagicNumber(magicNumber)
	checkVersion(version)
	if bodyLen > 0 {
		bodyBytes := context.Read(int(bodyLen))
		if err := proto.Unmarshal(bodyBytes, rpcMessage.Body);err!=nil{
			panic(err)
		}
	}
	context.AddAttr(constant.RPC_MESSAGE,rpcMessage)
}
func checkMagicNumber(mn byte) error{
	if constant.MAGIC_NUMBER != mn{
		return fmt.Errorf("error magic number")
	}
	return nil
}
func checkVersion(version byte) error{
	if constant.VERSION != version {
		return fmt.Errorf("required version %d, receive version %d",constant.VERSION,version)
	}
	return nil
}
func CreateRpcMessage(id,contentLength,headLength int64,msgType byte) *entity.RpcMessage{
	message := entity.NewRpcMessage(msgType)
	message.ContentLength = contentLength
	message.HeadLength = headLength
	message.Id = id
	message.MessageType = msgType
	return message
}
func setMsgBodyByType(rpcMessage *entity.RpcMessage){
	switch rpcMessage.MessageType {
	case constant.RPC_REQUEST:
		rpcMessage.Body = &protoc.RpcRequest{}
	case constant.RPC_RESPONSE:
		rpcMessage.Body = &protoc.RpcResponse{}
	}
}
