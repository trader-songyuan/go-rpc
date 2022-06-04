package loadbalance

import (
	"math/rand"
	"time"
)

// 循环负载均衡
type Random struct {}
func (r Random) Select(sname string,addresses []string) string{
	rand.Seed(time.Now().UnixNano())
	n := rand.Int31n(int32(len(addresses)))
	n++
	pos := int32(len(addresses)) % n
	return addresses[pos]
}
