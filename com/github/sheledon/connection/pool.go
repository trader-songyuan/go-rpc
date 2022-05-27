package connection

import (
	"log"
	"net"
)
type Pool struct {
	// key: remoteAddr value : connection
	pool map[string]*rpcConnection
}
type rpcConnection struct {
	conn net.Conn
	connContext *ConnectContext
	pipeline *Pipeline
}
func NewConnectionPool() *Pool {
	return &Pool{
		pool: make(map[string]*rpcConnection),
	}
}
func NewRpcConnection(conn net.Conn) *rpcConnection {
	log.Printf("create new rpcConnection : %s",conn.RemoteAddr())
	rc := &rpcConnection{
		conn:     conn,
		pipeline: NewDefaultPipeline(),
	}
	rc.connContext = NewConnectContext(rc)
	return rc
}
func (r *rpcConnection) close()  {
	r.conn.Close()
}
func (cp *Pool) AddConnection(conn net.Conn){
	key := conn.RemoteAddr().String()
	if oc,ok := cp.pool[key];ok{
		if oc.conn != conn {
			defer oc.close()
		}
	}else{
		cp.pool[key] = NewRpcConnection(conn)
	}
	cp.pool[key].process()
}
func (r *rpcConnection) process(){
	r.pipeline.processRequest(r.connContext)
}
