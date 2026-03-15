package api

type Server interface {
	Start(addr string) error
}
