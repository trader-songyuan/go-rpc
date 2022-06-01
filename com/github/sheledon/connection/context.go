package connection

import (
	"bufio"
	"fmt"
	"go-rpc/com/github/sheledon/buffer"
	"go-rpc/com/github/sheledon/entity"
)
type ConnectContext struct {
	readBuffer *buffer.ReadBuffer
	writer     *bufio.Writer
	conn       *RpcConnection
	Attr       map[string]interface{}
}
func NewConnectContext(conn *RpcConnection) *ConnectContext {
	return &ConnectContext{
		conn:       conn,
		readBuffer: buffer.NewReadBuffer(conn.Conn),
		writer:     bufio.NewWriter(conn.Conn),
		Attr:       make(map[string]interface{}),
	}
}
func (ctx *ConnectContext) WriteBytesAndFlushed(bytes []byte) {
	ctx.writer.Write(bytes)
	ctx.writer.Flush()
}
func (ctx *ConnectContext) WriteBytes(bytes []byte)  {
	ctx.writer.Write(bytes)
}
func (ctx *ConnectContext) Write(b byte)  {
	ctx.writer.WriteByte(b)
}
func (ctx *ConnectContext) Flush() {
	ctx.writer.Flush()
}
func (ctx *ConnectContext) AddAttr(key string,value interface{})  {
	ctx.Attr[key] = value
}
func (ctx *ConnectContext) GetAttr(key string) interface{}{
	if attr,ok := ctx.Attr[key]; ok{
		return attr
	}
	return nil
}
func (ctx *ConnectContext) SendRequest(msg *entity.RpcMessage) {
	ctx.conn.ProcessWrite(msg)
}
func (ctx *ConnectContext) ResetAttr() {
	ctx.Attr = make(map[string]interface{})
}
func (ctx *ConnectContext) ReadByte() byte{
	b, err := ctx.readBuffer.ReadByte()
	ctx.handleReadError(err)
	return b
}
func (ctx *ConnectContext) ReadInt64() int64{
	i64, err := ctx.readBuffer.ReadInt64()
	ctx.handleReadError(err)
	return i64
}
func (ctx *ConnectContext) Read(lens int) (reb []byte) {
	bs,err := ctx.readBuffer.Read(lens)
	ctx.handleReadError(err)
	return bs
}
func (ctx *ConnectContext) handleReadError(err error)  {
	if err != nil {
		fmt.Println(err)
		ctx.conn.close()
	}
}
