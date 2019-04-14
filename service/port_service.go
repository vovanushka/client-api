package service

import (
	"log"

	"github.com/vovanushka/port-service/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type PortService struct {
	conn   *grpc.ClientConn
	client api.PortClient
}

func (s *PortService) Connect(url string) error {
	var err error
	s.conn, err = grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}

	s.client = api.NewPortClient(s.conn)

	return err
}

func (s *PortService) Save(data []byte) error {
	_, err := s.client.SavePort(context.Background(), &api.PortMessage{Data: data})
	if err != nil {
		log.Fatalf("Error when calling SavePort: %s", err)
	}

	return err
}

func (s *PortService) Get(id string) ([]byte, error) {
	port, err := s.client.GetPort(context.Background(), &api.PortIDMessage{Id: id})
	if err != nil {
		return nil, err
	}
	return port.Data, err
}

func (s *PortService) Close() error {
	return s.conn.Close()
}
