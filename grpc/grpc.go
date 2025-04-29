package grpc

type Grpc interface {
	Ignite() error
	Stop() error
	// Add more methods here
}

type grpc struct {}

func Newgrpc() (Grpc, error) {
  panic("implement me")
}

func (m grpc) Ignite() error {
  panic("implement me")
}

func (m grpc) Stop() error {
  panic("implement me")
}
