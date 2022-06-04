package service

import (
	"fmt"
	"github.com/samuel/go-zookeeper/zk"
	"strings"
	"sync"
	"time"
)
const PATH_PREFIX = "/GO_RPC"
const SPLIT = "/"
var ACLS = zk.WorldACL(zk.PermAll)
type ZkRegister struct {
	lock sync.Mutex
	hosts []string
	con *zk.Conn
	discoveryMap map[string][]string
}
func NewZkRegister(hosts []string) *ZkRegister{
	return &ZkRegister{
		hosts: hosts,
		discoveryMap: make(map[string][]string),
	}
}
type WatchChildrenCallback func(*zk.Stat,<-chan zk.Event)

func (register *ZkRegister) RegisterService(provider,serviceName,address string) {
	if !register.isConnectionAlive() {
		register.connect()
	}
	path := register.getPath(serviceName)
	if exists, _, err := register.con.Exists(path); !exists || err != nil {
		register.recursionCreatePath(path,nil,0)
	}
	path+=SPLIT+address
	create, err := register.con.Create(path, []byte(provider),zk.FlagEphemeral, ACLS)
	if err != nil {
		fmt.Printf("fail register service: %s ; err : %v\n",path,err)
		return
	}
	fmt.Printf("register service: %s\n",create)
}
func (register *ZkRegister) connect() {
	if !register.isConnectionAlive() {
		eventCallbackOption := zk.WithEventCallback(register.sessionCallback)
		connect, evtChan, err := zk.Connect(register.hosts,10*time.Minute,eventCallbackOption)
		if err != nil {
			panic(err)
		}
		var maxTry = 10
		for i:=0;i<maxTry;i++{
			select {
			case evt:=<-evtChan:
				if evt.State == zk.StateHasSession{
					register.con = connect
					return
				}
			case timeout := <-time.After(10*time.Second):
				panic(fmt.Sprintf("fail connect zk: %v", timeout))
			}
		}
	}
}
func (register *ZkRegister) getPath(serviceName string) string{
	return PATH_PREFIX + SPLIT + serviceName
}
func (register *ZkRegister) DiscoveryService(serviceName string) []string{
	if ch,ok := register.discoveryMap[serviceName]; ok {
		return ch
	}
	register.lock.Lock()
	defer register.lock.Unlock()
	if ch,ok := register.discoveryMap[serviceName]; ok {
		return ch
	}
	register.pullServiceProviderList(serviceName)
	return register.discoveryMap[serviceName]
}
func (register *ZkRegister) recursionCreatePath(path string ,data []byte, flags int32){
	ps := strings.Split(path,SPLIT)
	curPath := ""
	for _,p := range ps{
		if len(p) == 0{
			continue
		}
		curPath += SPLIT + p
		exists, _, err := register.con.Exists(curPath)
		if err != nil || !exists {
			_, err = register.con.Create(curPath, data, flags, ACLS)
			if err != nil {
				panic(err)
			}
		}
	}
}
func (register *ZkRegister) isConnectionAlive() bool{
	return  register.con != nil && register.con.State()==zk.StateHasSession
}
func (register *ZkRegister) sessionCallback(event zk.Event)  {
	fmt.Printf("zk connection(%s) state: { %s }\n",event.Server,event.State.String())
}
func (register *ZkRegister) pullServiceProviderList(serviceName string) {
	if !register.isConnectionAlive() {
		register.connect()
	}
	path := register.getPath(serviceName)
	children, _ , event ,err := register.con.ChildrenW(path)
	if err != nil {
		panic(err)
	}
	go func() {
		<- event
		register.pullServiceProviderList(serviceName)
	}()
	register.discoveryMap[serviceName] = children
}