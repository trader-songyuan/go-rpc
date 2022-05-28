package connection

import (
	"errors"
	"go-rpc/com/github/sheledon/entity"
	"log"
	"net"
)
type Pool struct {
	// key: remoteAddr value : connection
	pool map[string]*RpcConnection
}
type RpcConnection struct {
	conn net.Conn
	connContext *ConnectContext
	pipeline *Pipeline
}
func NewConnectionPool() *Pool {
	return &Pool{
		pool: make(map[string]*RpcConnection),
	}
}
func NewRpcConnection(conn net.Conn) *RpcConnection {
	log.Printf("create new rpcConnection : %s",conn.RemoteAddr())
	rc := &RpcConnection{
		conn:     conn,
		pipeline: NewDefaultPipeline(),
	}
	rc.connContext = NewConnectContext(rc)
	return rc
}
func (r *RpcConnection) close()  {
	r.conn.Close()
}
func (cp *Pool) AddConnection(conn net.Conn) *RpcConnection{
	key := conn.RemoteAddr().String()
	if oc,ok := cp.pool[key];ok{
		if oc.conn != conn {
			defer oc.close()
		}
	}else{
		cp.pool[key] = NewRpcConnection(conn)
	}
	return cp.pool[key]
}
func (cp *Pool) GetConnection(addr string)  (rc *RpcConnection,err error) {
	var ok bool
	if rc,ok = cp.pool[addr];!ok {
		return nil,errors.New("not exists")
	}
	return
}
func (r *RpcConnection) ProcessRequest(){
	r.pipeline.processRequest(r.connContext)
}
func (r *RpcConnection) SendMsg(msg *entity.RpcMessage){
	r.pipeline.sendRequest(r.connContext,msg)
}
