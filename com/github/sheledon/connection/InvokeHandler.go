package connection

import (
	"fmt"
	"go-rpc/com/github/sheledon/entity"
	"go-rpc/com/github/sheledon/service"
	"reflect"
)

type InvokeHandler struct {
	provider *service.Provider
}

func NewInvokeHandler() *InvokeHandler{
	return &InvokeHandler{
		provider: service.ServiceProvier,
	}
}
func (h *InvokeHandler) Read(context *ConnectContext)  {
	message := context.Obj.(*entity.RpcMessage)
	request := message.Body.(*entity.RpcRequest)
	svice, err := h.provider.GetService(request.ServiceName)
	if err != nil {
		//响应错误
		return
	}
	instance := reflect.ValueOf(svice)
	method := instance.MethodByName(request.MethodName)
	res := h.invokeMethod(method, request.Params)
	for _,v := range res {
		fmt.Println(v.String())
	}
}
func (h *InvokeHandler) invokeMethod(method reflect.Value,params []interface{}) []reflect.Value{
	vparams := make([]reflect.Value, len(params))
	for i,p := range params{
		vparams[i] = reflect.ValueOf(p)
	}
	result := method.Call(vparams)
	return result
}

