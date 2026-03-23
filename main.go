package main

import (
	"fmt"
	"strings"
)

func main() {
	server := NewServer(":5555")

	server.router.Post("/createUser", logRequest, validateRequest, func(req *Request, res *Response) bool {

		var requestObject RequestObject

		err := req.Json(&requestObject)
		if err != nil {
			res.StatusCode = 400
			res.Body = []byte("Invalid JSON")
			return false
		}

		res.Json(ReponseObject{
			UserId: 1,
			Name:   requestObject.Name,
			Email:  requestObject.Email,
		})
		res.StatusCode = 200

		return true
	})

	server.router.Get("/getUser", logRequest, func(req *Request, res *Response) bool {

		responseObject := ReponseObject{
			UserId: 1,
			Name:   "John Doe",
			Email:  ":(((((((((((((@gmail.com",
		}

		res.Json(responseObject)
		res.StatusCode = 200

		return true
	})

	server.router.Get("/sayHello", logRequest, func(req *Request, res *Response) bool {
		res.Json(ReponseObject{
			UserId: 1,
			Name:   "John Doe",
			Email:  "mohamecabdelaziz66@gmail.com",
		})
		res.StatusCode = 200

		return true
	})

	server.ListenAndStart()
}

/*  Helper utilities to test  */

func logRequest(req *Request, res *Response) bool {
	fmt.Printf("Request: %s %s\n", req.Method, req.Path)
	fmt.Printf("Body: %s\n", req.Body)

	return true
}

func validateRequest(req *Request, res *Response) bool {

	if req.Body == nil {
		res.StatusCode = 400
		res.Body = []byte("Invalid JSON")
		return false
	}

	if req.Method != "POST" {
		res.StatusCode = 400
		return false
	}

	var requestObject RequestObject
	err := req.Json(&requestObject)
	if err != nil {
		res.StatusCode = 400
		res.StatusText = statusText[400]
		res.Body = []byte("Invalid JSON")
		return false
	}

	if requestObject.Age < 18 {
		res.StatusCode = 400
		res.StatusText = statusText[400]
		res.Body = []byte("Invalid Age")
		return false
	}

	if requestObject.Email == "" || !strings.Contains(requestObject.Email, "@") {
		res.StatusCode = 400
		res.StatusText = statusText[400]
		res.Body = []byte("Invalid Email")
		return false
	}

	if requestObject.Name == "" || len(requestObject.Name) < 4 {
		res.StatusCode = 400
		res.StatusText = statusText[400]
		res.Body = []byte("Invalid Name")
		return false
	}

	return true

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
