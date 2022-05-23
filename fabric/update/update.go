package update

import (
	"context"
)

type Update struct {
	ch chan interface{}
}

func NewUpdate() *Update {
	return &Update{
		ch: make(chan interface{}),
	}
}

func (u *Update) Start(ctx context.Context) {
	//
}

func (u *Update) Stop() {
	close(u.ch)
}

func (u *Update) GetCh() chan<- interface{} {
	return u.ch
}
