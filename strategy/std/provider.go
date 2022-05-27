package std

import (
	"io"
)

type Provider struct {
}

func (p *Provider) Open(path string) ([]byte, error) {
	// TODO implement me
	panic("implement me")
}

func (p *Provider) Create(path, name string, reader io.Reader) (int64, error) {
	// TODO implement me
	panic("implement me")
}

func (p *Provider) Remove(path string) error {
	// TODO implement me
	panic("implement me")
}

func (p *Provider) Directory() string {
	// TODO implement me
	panic("implement me")
}

func (p *Provider) Name() string {
	// TODO implement me
	panic("implement me")
}
