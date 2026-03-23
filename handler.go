package main

type Handler func(req *Request, res *Response) bool

func HomeHandler(req *Request, res *Response) {
	res.Body = []byte("Hello World")
	res.StatusCode = 200
	res.Headers = map[string]string{
		"Content-Type": "text/plain",
	}
}

func OkHandler(req *Request, res *Response) {
	res.Body = []byte("Your request was successful!")
	res.StatusCode = 200
	res.Headers = map[string]string{
		"Content-Type": "text/plain",
		"Status":       "200 OK",
	}
}

func NotFoundHandler(req *Request, res *Response) {
	res.Body = []byte(nil)
	res.StatusCode = 404
	res.Headers = map[string]string{
		"Content-Type": "text/plain",
		"Status":       "404 Not Found",
	}
}

func InternalServerErrorHandler(req *Request, res *Response) {
	res.Body = []byte(nil)
	res.StatusCode = 500
	res.Headers = map[string]string{
		"Content-Type": "text/plain",
		"Status":       "500 Internal Server Error",
	}
}

// U can add more handlers and register them in the router
