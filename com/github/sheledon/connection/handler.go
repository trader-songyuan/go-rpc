package handler

import (
	"go-rpc/com/github/sheledon/connection"
	"go-rpc/com/github/sheledon/entity"
)

type RpcInboundHandler interface {
	Read(*connection.ConnectContext)
}
type RpcOutboundHandler interface {
	Write(*connection.ConnectContext,*entity.RpcMessage)
}