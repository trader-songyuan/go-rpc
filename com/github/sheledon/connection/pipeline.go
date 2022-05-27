package connection

type Pipeline struct {
	inboundHandlers []RpcInboundHandler
	outboundHandlers []RpcOutboundHandler
}
func NewDefaultPipeline() *Pipeline{
	p := &Pipeline{
		inboundHandlers: make([]RpcInboundHandler,0),
		outboundHandlers: make([]RpcOutboundHandler,0),
	}
	p.addInboundHandler(NewDecodeHandler())
	p.addInboundHandler(NewInvokeHandler())
	p.addOutboundHandler(NewEncodeHandler())
	return p
}
func (p *Pipeline) addInboundHandler(handler RpcInboundHandler) {
	p.inboundHandlers = append(p.inboundHandlers,handler)
}
func (p *Pipeline) addOutboundHandler(handler RpcOutboundHandler){
	p.outboundHandlers = append(p.outboundHandlers,handler)
}
func (p *Pipeline) processRequest(ctx *ConnectContext) {
	for _,h := range p.inboundHandlers{
		h.Read(ctx)
	}
}
