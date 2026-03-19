package main

type Router struct {
	routes map[string]Handler
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string]Handler),
	}
}

func (r *Router) registerRoute(path string, handler Handler) {
	r.routes[path] = handler
}

func (r *Router) Get(path string, handler Handler) {
	r.registerRoute("GET "+path, handler)
}

func (r *Router) Post(path string, handler Handler) {
	r.registerRoute("POST "+path, handler)
}

func (r *Router) Put(path string, handler Handler) {
	r.registerRoute("PUT "+path, handler)
}

func (r *Router) Delete(path string, handler Handler) {
	r.registerRoute("DELETE "+path, handler)
}

func (r *Router) ServeRequest(req *Request, res *Response) {
	if handler, ok := r.routes[req.Method+" "+req.Path]; ok {
		handler(req, res)
	} else {
		NotFoundHandler(req, res)
	}
}
