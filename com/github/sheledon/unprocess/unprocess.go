package unprocess

import (
	"errors"
	"time"
)
/**
client 网络相关
*/
var UprFactory = NewUnProcessRequestFactory()
type UnProcessRequestFactory struct {
	// key : id
	promiseMap map[string]*Promise
}
type Promise struct {
	successChannel chan interface{}
	failChanel chan string
}
func (promise Promise) Get() (interface{},error){
	var response interface{}
	defer promise.close()
	select {
	case response = <-promise.successChannel:
		return response,nil
	case err := <- promise.failChanel:
		return nil,errors.New(err)
	}
}
func (promise Promise) GetTimeOut(timeout time.Duration)(interface{},error){
	var response interface{}
	defer promise.close()
	select {
	case response = <-promise.successChannel:
		return response,nil
	case errs := <- promise.failChanel:
		return nil,errors.New(errs)
	case <-time.After(timeout):
		return nil,errors.New("get time out")
	}
}
func NewPromise() Promise {
	return Promise{
		successChannel: make(chan interface{}),
		failChanel: make(chan string),
	}
}
func (p Promise) close()  {
	close(p.successChannel)
	close(p.failChanel)
}
func NewUnProcessRequestFactory() UnProcessRequestFactory {
	return UnProcessRequestFactory{
		promiseMap: make(map[string]*Promise),
	}
}
func (p Promise) Complete(obj interface{}) {
	p.successChannel <- obj
}
func (p Promise) CompleteFailure(err string) {
	p.failChanel <- err
}

func (f UnProcessRequestFactory) Set(id string,promise Promise)  {
	f.promiseMap[id] = &promise
}
func (f UnProcessRequestFactory) Get(id string) *Promise {
	p,_ := f.promiseMap[id]
	return p
}
