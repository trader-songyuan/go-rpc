package connection

import (
	"bytes"
	"encoding/binary"
	"go-rpc/com/github/sheledon/entity"
	"google.golang.org/protobuf/proto"
	"sync"
)
/**
	出站编码器：rpcMessage ---> 大端序列字节流
 */
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
	idUtils *atomic
}
func NewEncodeHandler() *EncodeHandler {
	return &EncodeHandler{
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
	bodyBytes,_:= proto.Marshal(msg.Body)
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


