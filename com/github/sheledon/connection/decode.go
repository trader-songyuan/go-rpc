package connection

import (
	"fmt"
	"go-rpc/com/github/sheledon/constant"
	"go-rpc/com/github/sheledon/entity"
	"go-rpc/com/github/sheledon/utils/serializer"
)

type DecodeHandler struct {
	decoder serializer.Serializer
}
func NewDecodeHandler() *DecodeHandler {
	return &DecodeHandler{
		decoder: serializer.NewDefaultSerializer(),
	}
}
func (h DecodeHandler) Read(context *ConnectContext) {
	magicNumber, _ := context.readBuffer.ReadByte()
	version, _ := context.readBuffer.ReadByte()
	id, _ := context.readBuffer.ReadInt64()
	msgType, _ := context.readBuffer.ReadByte()
	contentLength, _ := context.readBuffer.ReadInt64()
	headLength, _ := context.readBuffer.ReadInt64()
	bodyLen := contentLength - headLength
	bodyEntity := getMsgBodyByType(msgType)
	if bodyLen > 0 {
		bodyBytes := context.readBuffer.Read(int(bodyLen))
		h.decoder.Deserialize(bodyBytes, bodyEntity)
	}
	checkMagicNumber(magicNumber)
	checkVersion(version)
	rpcMessage := createRpcMessage(id, contentLength, headLength, bodyEntity)
	context.SetObj(rpcMessage)
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
func createRpcMessage(id,contentLength,headLength int64,body interface{}) *entity.RpcMessage{
	message := entity.NewDefaultRpcMessage()
	message.ContentLength = contentLength
	message.HeadLength = headLength
	message.Body = body
	message.Id = id
	return message
}
func getMsgBodyByType(btype byte) interface{}{
	switch btype {
	case constant.RPC_REQUEST:
		return new(entity.RpcRequest)
	case constant.RPC_RESPONSE:
		return new(entity.RpcResponse)
	default:
		return nil
	}
}
