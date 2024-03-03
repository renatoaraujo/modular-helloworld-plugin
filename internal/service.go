package internal

import "fmt"

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) SayHello(name string) string {
	return fmt.Sprintf("Hello, %s", name)
}
