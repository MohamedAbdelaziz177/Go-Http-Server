package main

type Router struct {
	routes map[string][]Handler
}

func NewRouter() *Router {
	return &Router{
		routes: make(map[string][]Handler),
	}
}

func (r *Router) registerRoute(path string, handlers ...Handler) {
	r.routes[path] = append(r.routes[path], handlers...)
}

func (r *Router) Get(path string, handlers ...Handler) {
	r.registerRoute("GET "+path, handlers...)
}

func (r *Router) Post(path string, handlers ...Handler) {
	r.registerRoute("POST "+path, handlers...)
}

func (r *Router) Put(path string, handlers ...Handler) {
	r.registerRoute("PUT "+path, handlers...)
}

func (r *Router) Delete(path string, handlers ...Handler) {
	r.registerRoute("DELETE "+path, handlers...)
}

func (r *Router) ServeRequest(req *Request, res *Response) error {

	key := req.Method + " " + req.Path

	if handlers, ok := r.routes[key]; ok {

		for _, handler := range handlers {
			if !handler(req, res) {
				break
			}
		}

	} else {
		NotFoundHandler(req, res)
	}

	return nil
}
