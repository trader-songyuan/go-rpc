package service

import (
	"fmt"
	"sync"
)

// 服务提供类，服务端调用
var pmap = make(map[string]*Provider)
var lock sync.Mutex
type Provider struct {
	servicesContainer map[string]interface{}
	register *ZkRegister
	serverAddress string
}
func GetProvider(serverAddress string) *Provider {
	if _,ok := pmap[serverAddress];!ok {
		lock.Lock()
		defer lock.Unlock()
		if _,ok = pmap[serverAddress];!ok {
			pmap[serverAddress] = &Provider{
				servicesContainer: make(map[string]interface{}),
				serverAddress: serverAddress,
			}
		}
	}
	return pmap[serverAddress]
}
func (p *Provider) SetRegister(register *ZkRegister){
	p.register = register
}
func (p *Provider) RegisterService(providerName string,serviceName string,v interface{}) error{
	if _,ok := p.servicesContainer[serviceName];ok{
		return fmt.Errorf("service %s already exists", serviceName)
	}
	if p.register != nil {
		p.register.RegisterService(providerName,serviceName,p.serverAddress)
	}
	p.servicesContainer[serviceName] = v
	return nil
}
func (p *Provider) GetService(name string) (interface{},error){
	if _,ok := p.servicesContainer[name];!ok{
		return nil,fmt.Errorf("not found service : %s",name)
	}
	return p.servicesContainer[name],nil
}
