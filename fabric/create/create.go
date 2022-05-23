package create

import (
	"context"
)

type Create struct {
	ch chan interface{}
}

func NewCreate() *Create {
	return &Create{
		ch: make(chan interface{}),
	}
}

func (c *Create) Start(ctx context.Context) {
	//
}

func (c *Create) Stop() {
	close(c.ch)
}

func (c *Create) GetCh() chan<- interface{} {
	return c.ch
}
