package utils

import (
	"fmt"
	"go-rpc/com/github/sheledon/entity/protoc"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"reflect"
	"strconv"
)
func ValueTransferToRpcAny(value reflect.Value) *protoc.RpcAny{
	kind := value.Kind()
	if kind == reflect.Ptr{
		value = value.Elem()
		kind = value.Kind()
	}
	var anyValue []*anypb.Any
	var avalue *anypb.Any
	var anyType protoc.AnyOriginalType
	var eleType protoc.AnyOriginalType
	switch kind {
	case reflect.String:
		avalue,_ = anypb.New(wrapperspb.String(value.String()))
		anyValue = append(anyValue, avalue)
		anyType = protoc.AnyOriginalType_STRING
	case reflect.Int8:
		avalue,_ =anypb.New(wrapperspb.Int32(int32(value.Int())))
		anyValue = append(anyValue, avalue)
		anyType = protoc.AnyOriginalType_INT_8
	case reflect.Int16:
		avalue,_ =anypb.New(wrapperspb.Int32(int32(value.Int())))
		anyValue = append(anyValue, avalue)
		anyType = protoc.AnyOriginalType_INT_16
	case reflect.Int32:
		avalue,_ =anypb.New(wrapperspb.Int32(int32(value.Int())))
		anyValue = append(anyValue, avalue)
		anyType = protoc.AnyOriginalType_INT_32
	case reflect.Int64:
		avalue,_ =anypb.New(wrapperspb.Int64(value.Int()))
		anyValue = append(anyValue, avalue)
		anyType = protoc.AnyOriginalType_INT_64
	case reflect.Uint8:
		avalue,_ = anypb.New(wrapperspb.UInt32(uint32(value.Uint())))
		anyValue = append(anyValue, avalue)
		anyType = protoc.AnyOriginalType_UINT_8
	case reflect.Uint16:
		avalue,_ = anypb.New(wrapperspb.UInt32(uint32(value.Uint())))
		anyValue = append(anyValue, avalue)
		anyType = protoc.AnyOriginalType_UINT_16
	case reflect.Uint32:
		avalue,_ = anypb.New(wrapperspb.UInt32(uint32(value.Uint())))
		anyValue = append(anyValue, avalue)
		anyType = protoc.AnyOriginalType_UINT_32
	case reflect.Uint64:
		avalue,_ =anypb.New(wrapperspb.UInt64(value.Uint()))
		anyValue = append(anyValue, avalue)
		anyType = protoc.AnyOriginalType_UINT_64
	case reflect.Float32:
		avalue,_ =anypb.New(wrapperspb.Double(value.Float()))
		anyValue = append(anyValue, avalue)
		anyType = protoc.AnyOriginalType_FLOAT_32
	case reflect.Float64:
		avalue,_=anypb.New(wrapperspb.Double(value.Float()))
		anyValue = append(anyValue, avalue)
		anyType = protoc.AnyOriginalType_FLOAT_64
	case reflect.Bool:
		avalue,_=anypb.New(wrapperspb.String(strconv.FormatBool(value.Bool())))
		anyValue = append(anyValue, avalue)
		anyType = protoc.AnyOriginalType_BOOL
	//case reflect.Array:
	//	anyType = protoc.AnyOriginalType_ARRAY
	//	fallthrough
	//case reflect.Slice:
	//	anyType = protoc.AnyOriginalType_SLICE
	//	for i:=0;i<value.Len();i++ {
	//		rpcAny := ValueTransferToRpcAny(value.Index(i))
	//		if i == 0 {
	//			eleType = rpcAny.Type
	//		}
	//		anyValue = append(anyValue, rpcAny.Value...)
	//	}
	case reflect.Struct:
		avalue,_=anypb.New(value.Addr().Interface().(proto.Message))
		anyValue = append(anyValue, avalue)
		anyType = protoc.AnyOriginalType_STRUCT
	default:
		panic(fmt.Sprintf("not support this kind: %v",kind))
	}
	return &protoc.RpcAny{
		Type:  anyType,
		Value: anyValue,
		EleType: eleType,
	}
}
func RpcAnyToReflectValue(rpcAny *protoc.RpcAny)  reflect.Value{
	var rv interface{}
	switch rpcAny.Type {
	case protoc.AnyOriginalType_STRING:
		rv = string(rpcAny.Value[0].GetValue())
	case protoc.AnyOriginalType_INT_8:
		w := wrapperspb.Int32(0)
		err := rpcAny.Value[0].UnmarshalTo(w)
		if err != nil {
			panic(err)
		}
		rv = int8(w.Value)
	case protoc.AnyOriginalType_INT_16:
		w := wrapperspb.Int32(0)
		err := rpcAny.Value[0].UnmarshalTo(w)
		if err != nil {
			panic(err)
		}
		rv = int16(w.Value)
	case protoc.AnyOriginalType_INT_32:
		w := wrapperspb.Int32(0)
		err := rpcAny.Value[0].UnmarshalTo(w)
		if err != nil {
			panic(err)
		}
		rv = w.Value
	case protoc.AnyOriginalType_INT_64:
		w := wrapperspb.Int64(0)
		err := rpcAny.Value[0].UnmarshalTo(w)
		if err != nil {
			panic(err)
		}
		rv = w.Value
	case protoc.AnyOriginalType_UINT_8:
		w := wrapperspb.UInt32(0)
		err := rpcAny.Value[0].UnmarshalTo(w)
		if err != nil {
			panic(err)
		}
		rv = uint8(w.Value)
	case protoc.AnyOriginalType_UINT_16:
		w := wrapperspb.UInt32(0)
		err := rpcAny.Value[0].UnmarshalTo(w)
		if err != nil {
			panic(err)
		}
		rv = uint16(w.Value)
	case protoc.AnyOriginalType_UINT_32:
		w := wrapperspb.UInt32(0)
		err := rpcAny.Value[0].UnmarshalTo(w)
		if err != nil {
			panic(err)
		}
		rv = w.Value
	case protoc.AnyOriginalType_UINT_64:
		w := wrapperspb.UInt64(0)
		err := rpcAny.Value[0].UnmarshalTo(w)
		if err != nil {
			panic(err)
		}
		rv = w.Value
	case protoc.AnyOriginalType_FLOAT_32:
		w := wrapperspb.Double(0)
		err := rpcAny.Value[0].UnmarshalTo(w)
		if err != nil {
			panic(err)
		}
		rv = float32(w.Value)
	case protoc.AnyOriginalType_FLOAT_64:
		w := wrapperspb.Double(0)
		err := rpcAny.Value[0].UnmarshalTo(w)
		if err != nil {
			panic(err)
		}
		rv = w.Value
	case protoc.AnyOriginalType_BOOL:
		rv , _ = strconv.ParseBool(string(rpcAny.Value[0].GetValue()))
	case protoc.AnyOriginalType_STRUCT:
		rv ,_ = rpcAny.Value[0].UnmarshalNew()
	//	支持这俩种类型又陷入了 如何将 interface{} ---> 具体类型的值 的问题
	//case protoc.AnyOriginalType_ARRAY:
	//case protoc.AnyOriginalType_SLICE:
	//	var rva = make([]reflect.Value,len(rpcAny.Value))
	//	for i,rc := range rpcAny.Value{
	//		rva[i] = RpcAnyToReflectValue(&protoc.RpcAny{Type: rpcAny.EleType,Value: []*anypb.Any{rc}})
	//	}
	//	rv = rva
	default:
		panic("not support this kind")
	}
	return reflect.ValueOf(rv)
}