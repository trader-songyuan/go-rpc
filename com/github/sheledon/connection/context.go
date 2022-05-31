package connection

import (
	"bufio"
	"go-rpc/com/github/sheledon/buffer"
	"go-rpc/com/github/sheledon/entity"
)
type ConnectContext struct {
	ReadBuffer *buffer.ReadBuffer
	writer     *bufio.Writer
	conn       *RpcConnection
	Attr       map[string]interface{}
}
func NewConnectContext(conn *RpcConnection) *ConnectContext {
	return &ConnectContext{
		conn:       conn,
		ReadBuffer: buffer.NewReadBuffer(conn.Conn),
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
