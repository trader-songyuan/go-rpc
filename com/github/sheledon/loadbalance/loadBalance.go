package loadbalance

type LoadBalance interface {
	Select(serviceName string,addresses []string) string
}
