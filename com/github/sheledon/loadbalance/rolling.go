package loadbalance

// 循环负载均衡

type Rolling struct {
	
}
func (r Rolling) Select(list []interface{}) interface{}{
	return 0
}
