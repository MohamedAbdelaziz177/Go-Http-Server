package main

import "fmt"

type Middlware struct {
	handler Handler
	next    *Middlware
}

func NewMiddlware(handler Handler, next *Middlware) *Middlware {
	return &Middlware{
		handler: handler,
		next:    next,
	}
}

func (m *Middlware) setHandler(handler Handler) {
	m.handler = handler
}

func (m *Middlware) setNext(next *Middlware) {
	m.next = next
}

func (m *Middlware) execute(req *Request, res *Response) {
	if m.handler(req, res) {
		m.callNext(req, res)
	}

}

func (m *Middlware) callNext(req *Request, res *Response) {
	if m.next != nil {
		m.next.execute(req, res)
	}
}

func registerMiddlewares(handlers []Handler) (*Middlware, error) {

	if len(handlers) == 0 {
		return nil, fmt.Errorf("No middlewares provided")
	}

	if len(handlers) == 1 {
		return NewMiddlware(handlers[0], nil), nil
	}

	var firstMW *Middlware = NewMiddlware(handlers[0], nil)
	var currentMW *Middlware = firstMW

	for _, handler := range handlers[1:] {
		currentMW.next = NewMiddlware(handler, nil)
		currentMW = currentMW.next
	}

	return firstMW, nil
}
