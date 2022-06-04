package service

import (
	"fmt"
	"testing"
	"time"
)

func TestRegister(t *testing.T)  {
	register := ZkRegister{
		hosts: []string{"127.0.0.1:2181"},
	}
	register.RegisterService("ptest1","helloService","127.0.0.1:8080")
	fmt.Println(register.DiscoveryService("helloService"))
	time.Sleep(10000*time.Second)
}
