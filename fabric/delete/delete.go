package delete

import (
	"context"
)

type Delete struct {
	ch chan interface{}
}

func NewDelete() *Delete {
	return &Delete{
		ch: make(chan interface{}),
	}
}

func (d *Delete) Start(ctx context.Context) {
	//
}

func (d *Delete) Stop() {
	close(d.ch)
}

func (d *Delete) GetCh() chan<- interface{} {
	return d.ch
}
