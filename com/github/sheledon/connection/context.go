package connection

import (
	"bufio"
	"go-rpc/com/github/sheledon/buffer"
)
type ConnectContext struct {
	readBuffer *buffer.ReadBuffer
	writer     *bufio.Writer
	conn       *RpcConnection
	Obj        interface{}
}
func NewConnectContext(conn *RpcConnection) *ConnectContext{
	return &ConnectContext{
		conn:       conn,
		readBuffer: buffer.NewReadBuffer(conn.conn),
		writer:     bufio.NewWriter(conn.conn),
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
func (ctx *ConnectContext) SetObj(obj interface{})  {
	ctx.Obj = obj
}
func (ctx *ConnectContext) GetObj() interface{}{
	return ctx.Obj
}
