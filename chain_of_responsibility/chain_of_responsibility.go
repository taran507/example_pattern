package chain_of_responsibility

import "fmt"

type Executor interface {
	SetNext(Executor) Executor
	Execute(login string) string
}

type Email struct {
	next Executor
}

func (e *Email) SetNext(executor Executor) Executor {
	e.next = executor
	return e
}

func (e *Email) Execute(login string) string {
	if e.isEmail(login) {
		return "email"
	}
	if e.next != nil {
		return e.next.Execute(login)
	}
	return ""
}

func NewEmail() *Email {
	return &Email{}
}

type Phone struct {
	next Executor
}

func (p *Phone) SetNext(executor Executor) Executor {
	p.next = executor
	return p
}

func (p *Phone) Execute(login string) string {
	if p.isPhone(login) {
		return "phone"
	}
	if p.next != nil {
		return p.next.Execute(login)
	}
	return ""
}

func NewPhone() *Phone {
	return &Phone{}
}

func main() {
	e := NewEmail()
	e.SetNext(NewPhone())

	loginType := e.Execute("+79089881078")
	fmt.Println(loginType)
}

func (e *Email) isEmail(login string) bool {
	return false
}

func (p *Phone) isPhone(login string) bool {
	return false
}
