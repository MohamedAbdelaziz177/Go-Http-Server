package main

import "fmt"

func main() {
	server := NewServer(":5555")

	server.router.Get("/hello", func(req *Request, res *Response) {

		if name, ok := req.Params["name"]; ok {
			res.Body = []byte("Hello " + name)
		} else {
			res.Body = []byte("Hello World")
		}

		res.StatusCode = 200
		res.Headers = map[string]string{
			"Content-Type": "text/plain",
		}
	})

	server.router.Post("/hello", func(req *Request, res *Response) {

		fmt.Printf("Request: %s %s\n", req.Method, req.Path)
		fmt.Printf("Body: %s\n", req.Body)

		res.Body = []byte("Hello World")
		res.StatusCode = 200
		res.Headers = map[string]string{
			"Content-Type": "text/plain",
		}
	})

	server.ListenAndStart()
}
