package fabric

import (
	"context"
	"fmt"

	"example/fabric/create"
	deleteUser "example/fabric/delete"
	"example/fabric/update"
)

type Auth3 interface {
	Start(ctx context.Context)
	Stop()
	GetCh() chan<- interface{}
}

const (
	CreateUser = iota + 1
	Update
	Delete
)

// Provider структура, позволяющая объявлять различных обработчиков по их id
type Provider struct{}

func NewProvider() *Provider {
	return &Provider{}
}

type Provider2 struct{}

func NewProvider2() *Provider2 {
	return &Provider2{}
}

func (p *Provider2) GetManager(id int) Auth3 {
	// TODO implement me
	panic("implement me")
}

func (p *Provider) GetManager(id int) Auth3 {
	switch id {
	case CreateUser:
		return create.NewCreate()
	case Delete:
		return deleteUser.NewDelete()
	case Update:
		return update.NewUpdate()
	default:
		panic(fmt.Errorf("не найденно %v", id))
	}
}
