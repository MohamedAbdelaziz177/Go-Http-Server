package main

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type Request struct {
	Proto    string
	Method   string
	Endpoint string
	Headers  map[string]string
	Params   map[string]string
	Body     []byte
}

func NewRequest(proto string, method string, endpoint string, headers map[string]string, params map[string]string, body []byte) *Request {
	return &Request{
		Proto:    proto,
		Method:   method,
		Endpoint: endpoint,
		Headers:  headers,
		Params:   params,
		Body:     body,
	}
}

func ParseRequest(conn net.Conn) (*Request, error) {

	reader := bufio.NewReader(conn)

	line, err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")

	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid request line: %s", line)
	}

	method := parts[0]
	endpoint := parts[1]
	proto := parts[2]

	query := strings.Split(endpoint, "?")[1]

	params := make(map[string]string)
	if len(query) == 1 {

		queryParamsAsString := strings.Split(query, "&")
	}

	headers := make(map[string]string)

	for {

		line, err := reader.ReadString('\n')

		if err != nil {
			return nil, err
		}

		line = strings.TrimSpace(line)

		if line == "" {
			break
		}

		parts := strings.SplitN(line, ":", 2)

		if len(parts) != 2 {
			return nil, fmt.Errorf("invalid header: %s", line)
		}

		headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1])
	}

	if cl, ok := headers["content-legnth"]; ok {

		legnth, err := strconv.Atoi(cl)

		if err != nil {
			return nil, fmt.Errorf("invalid content-length: %s", cl)
		}

		body := make([]byte, legnth)
		_, err = reader.Read(body)

		if err != nil {
			return nil, fmt.Errorf("invalid body: %s", err)
		}

		return NewRequest(proto, method, endpoint, headers, params, body), nil
	}

	return NewRequest(proto, method, endpoint, headers, params, nil), nil
}
