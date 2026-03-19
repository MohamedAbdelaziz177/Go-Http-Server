package main

import "fmt"

func main() {
	server := NewServer(":5555")

	server.router.Post("/createUser", func(req *Request, res *Response) {
		var requestObject RequestObject

		err := req.Json(&requestObject)
		if err != nil {
			res.StatusCode = 400
			res.Body = []byte("Invalid JSON")
			return
		}

		fmt.Printf("Request: %s %s\n", req.Method, req.Path)
		fmt.Printf("Body: %s\n", req.Body)

		res.Json(ReponseObject{
			UserId: 1,
			Name:   requestObject.Name,
			Email:  requestObject.Email,
		})
		res.StatusCode = 200
	})

	server.router.Get("/getUser", func(req *Request, res *Response) {

		fmt.Printf("Request: %s %s\n", req.Method, req.Path)
		fmt.Printf("Body: %s\n", req.Body)

		responseObject := ReponseObject{
			UserId: 1,
			Name:   "John Doe",
			Email:  "[EMAIL_ADDRESS]",
		}

		res.Json(responseObject)
		res.StatusCode = 200

	})

	server.ListenAndStart()
}

type ReponseObject struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
}

type RequestObject struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}
