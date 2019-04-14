package service

type Service interface {
	Connect(url string) error
	Close() error
	Save(data []byte) error
	Get(id string) ([]byte, error)
}
