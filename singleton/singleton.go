package singleton

import (
	"fmt"
)

type log struct{}

func (l *log) Info(msg string) {
	fmt.Println(msg)
}

func newLog() *log {
	return &log{}
}

var l *log

// InitLogger инициализирует стандартный логер.
func InitLogger() *log {
	if l == nil {
		l = newLog()
	}
	return l
}

func Info(msg string) {
	InitLogger().Info(msg)
}
