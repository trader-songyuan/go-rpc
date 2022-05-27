package connection

import "go-rpc/com/github/sheledon/entity"

type RpcInboundHandler interface {
	Read(*ConnectContext)
}
type RpcOutboundHandler interface {
	Write(*ConnectContext,*entity.RpcMessage)
}