package main

import (
	"encoding/json"
	"fmt"
	"net"
)

type Response struct {
	StatusCode int
	StatusText string
	Headers    map[string]string
	Body       []byte
}

var statusText = map[int]string{
	200: "OK",
	400: "Bad Request",
	404: "Not Found",
	500: "Internal Server Error",
}

func NewResponse(statusCode int, headers map[string]string, body []byte) *Response {
	return &Response{
		StatusCode: statusCode,
		StatusText: statusText[statusCode],
		Headers:    headers,
		Body:       body,
	}
}

func (r *Response) Send(conn net.Conn) error {
	_, err := fmt.Fprintf(conn, "HTTP/1.1 %d %s\r\n", r.StatusCode, r.StatusText)
	if err != nil {
		return err
	}

	for k, v := range r.Headers {
		_, err = fmt.Fprintf(conn, "%s: %s\r\n", k, v)
		if err != nil {
			return err
		}
	}

	fmt.Fprintf(conn, "\r\n")

	if _, err := conn.Write(r.Body); err != nil {
		return err
	}

	return nil
}

func (r *Response) jsonify(data []byte) error {
	jsonBody, err := json.Marshal(data)
	if err != nil {
		return err
	}
	r.Body = jsonBody
	
	if r.Headers == nil {
		r.Headers = make(map[string]string)
	}
	r.Headers["Content-Type"] = "application/json"
	
	return nil
}
