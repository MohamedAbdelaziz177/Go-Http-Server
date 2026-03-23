package main

import "fmt"

type Middlware struct {
	handler Handler
	next    Handler
}

func NewMiddlware(handler Handler, next Handler) *Middlware {
	return &Middlware{
		handler: handler,
		next:    next,
	}
}

func (m *Middlware) setHandler(handler Handler) {
	m.handler = handler
}

func (m *Middlware) setNext(next Handler) {
	m.next = next
}

func (m *Middlware) execute(req *Request, res *Response) {
	m.handler(req, res)
	m.callNext(req, res)
}

func (m *Middlware) callNext(req *Request, res *Response) {
	if m.next != nil {
		m.next(req, res)
	}
}

func registerMiddlewares(middlewares []Handler) (*Middlware, error) {

	if len(middlewares) == 0 {
		return nil, fmt.Errorf("No middlewares provided")
	}

	if len(middlewares) == 1 {
		return NewMiddlware(middlewares[0], nil), nil
	}

	var firstMW *Middlware = NewMiddlware(middlewares[0], nil)
	var currentMW *Middlware = firstMW

	for _, middleware := range middlewares[1:] {
		currentMW.setNext(middleware)
		currentMW = NewMiddlware(currentMW.next, nil)
	}

	return firstMW, nil
}
