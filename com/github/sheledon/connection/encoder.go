package connection

import (
	"bytes"
	"encoding/binary"
	"go-rpc/com/github/sheledon/entity"
	"go-rpc/com/github/sheledon/utils/serializer"
	"sync"
)

type atomic struct {
	mutex sync.Mutex
	num int64
}

func (a *atomic) getAndAdd() int64{
	a.mutex.Lock()
	defer a.mutex.Unlock()
	a.num++
	return a.num
}
type EncodeHandler struct {
	encoder serializer.Serializer
	idUtils *atomic
}
func NewEncodeHandler() *EncodeHandler {
	return &EncodeHandler{
		encoder: serializer.NewDefaultSerializer(),
		idUtils: new(atomic),
	}
}
func (h *EncodeHandler) Write(ctx *ConnectContext, msg *entity.RpcMessage)  {
	//按照协议格式向通道写入数据
	ctx.Write(msg.MagicNumber)
	ctx.Write(msg.Version)
	ctx.WriteBytes(TransferInt64ToBytes(h.idUtils.getAndAdd()))
	ctx.Write(msg.MessageType)
	headLen := msg.HeadLength
	bodyBytes,_:= h.encoder.Serialize(msg.Body)
	contLen := headLen + int64(len(bodyBytes))
	ctx.WriteBytes(TransferInt64ToBytes(contLen))
	ctx.WriteBytes(TransferInt64ToBytes(headLen))
	ctx.WriteBytes(bodyBytes)
	ctx.Flush()
}
// 默认都是大端序列传递
func TransferInt64ToBytes(i int64) []byte{
	buf := new(bytes.Buffer)
	binary.Write(buf,binary.BigEndian,&i)
	return buf.Bytes()
}


