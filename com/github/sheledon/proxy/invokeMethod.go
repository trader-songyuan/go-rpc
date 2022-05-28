package proxy

import "reflect"

type InvocationMethod struct {
	Name string
	Type reflect.Type
}

func (im InvocationMethod) Invoke(obj interface{}, args []reflect.Value) []reflect.Value {
	v := reflect.ValueOf(obj).MethodByName(im.Name)

	if !v.IsValid() {
		panic("Can not found method " + im.Name + " in " + reflect.ValueOf(obj).Type().String())
	}

	return v.Call(args)
}

