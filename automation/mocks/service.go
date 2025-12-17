package main

//go:generate mockgen -source=service.go -destination=service_mock.go -package=main

type DataStore interface {
	GetData(id string) (string, error)
	SaveData(id, data string) error
}

type MyService struct {
	store DataStore
}

// go install github.com/golang/mock/mockgen@latest
// go generate ./...