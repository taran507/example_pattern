package abstractfactory

import (
	"fmt"

	"example/fabric"
)

type AuthProvider interface {
	GetManager(id int) fabric.Auth3
}

const (
	vk = iota + 1
	google
)

type AbstractProvider struct{}

func (a *AbstractProvider) GetProvider(providerID int) AuthProvider {
	switch providerID {
	case vk:
		return fabric.NewProvider()
	case google:
		return fabric.NewProvider2()
	default:
		panic(fmt.Errorf("не найденно %v", providerID))
	}
}
