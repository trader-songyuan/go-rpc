package connection

type Strategy interface {
	Handle(ctx *ConnectContext)
}
