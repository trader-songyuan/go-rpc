package proxy

import "reflect"
var InvocationProxy = invocationProxy{}

type invocationProxy struct {}

type InvocationHandler func(obj interface{}, method InvocationMethod, args []reflect.Value) []reflect.Value

func (ip invocationProxy) NewProxyInstance(target interface{}, handler InvocationHandler) interface{} {
	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr {
		panic("Need a pointer of interface struct")
	}
	if targetValue.Elem().Kind() != reflect.Struct {
		panic("Need a pointer of interface struct")
	}
	targetEle := targetValue.Elem()
	targetType := reflect.TypeOf(target).Elem()
	for i:=0; i< targetEle.NumField(); i++ {
		field := targetEle.Field(i)
		tfield := targetType.Field(i)
		if field.Kind() != reflect.Func {
			continue
		}
		makeFunc := reflect.MakeFunc(field.Type(), func(args []reflect.Value) (results []reflect.Value) {
			method := InvocationMethod{tfield.Name, tfield.Type}
			return handler(target,method,args)
		})
		field.Set(makeFunc)
	}
	return target
}