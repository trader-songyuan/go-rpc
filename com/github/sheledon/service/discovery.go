package service

type Discovery interface {
	DiscoveryService(serviceName string) []string
}


