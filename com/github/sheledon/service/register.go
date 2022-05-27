package service

type Register interface {
	RegisterService(string,interface{}) error
}
type Discovery interface {
	DiscoveryServices(string) []string
}
