package client

import (
	"errors"
	"go-rpc/com/github/sheledon/entity"
	"reflect"
	"time"
)

/**
client 网络相关
*/
type UnProcessRequestFactory struct {
	// key : id
	promiseMap map[string]Promise
}
type Promise struct {
	successChannel chan interface{}
	failChanel chan error
	Types reflect.Type
}
func (promise Promise) Get() (interface{},error){
	var response interface{}
	var err error
	defer promise.close()
	select {
	case response = <-promise.successChannel:
		return response,nil
	case err = <- promise.failChanel:
		return nil,err
	}
}
func (promise Promise) GetTimeOut(timeout time.Duration)(interface{},error){
	var response interface{}
	var err error
	defer promise.close()
	select {
	case response = <-promise.successChannel:
		return response,nil
	case err = <- promise.failChanel:
		return nil,err
	case <-time.After(timeout):
		return nil,errors.New("get time out")
	}
}
func NewPromise() Promise {
	return Promise{successChannel: make(chan interface{})}
}
func (p Promise) close()  {
	close(p.successChannel)
	close(p.failChanel)
}
func newUnProcessRequestFactory() UnProcessRequestFactory{
	return UnProcessRequestFactory{
		promiseMap: make(map[string]Promise),
	}
}
func (p Promise) Complete(response entity.RpcResponse) {
	p.successChannel <-response
}
func (p Promise) CompleteFailure(err error) {
	p.failChanel <- err
}

func (f UnProcessRequestFactory) Set(id string,promise Promise)  {
	f.promiseMap[id] = promise
}
func (f UnProcessRequestFactory) Get(id string) Promise{
	p,_ := f.promiseMap[id]
	return p
}
