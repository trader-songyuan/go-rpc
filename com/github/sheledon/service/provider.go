package service

import "fmt"

// 服务提供类，服务端调用
var ServiceProvier = newProvider()
type Provider struct {
	servicesContainer map[string]interface{}
}
func newProvider() *Provider {
	return &Provider{
		servicesContainer: make(map[string]interface{}),
	}
}
func (p *Provider) RegisterService(name string,v interface{}) error{
	if _,ok := p.servicesContainer[name];ok{
		return fmt.Errorf("service %s already exists",name)
	}
	p.servicesContainer[name] = v
	return nil
}
func (p *Provider) GetService(name string) (interface{},error){
	if _,ok := p.servicesContainer[name];!ok{
		return nil,fmt.Errorf("not found service : %s",name)
	}
	return p.servicesContainer[name],nil
}
