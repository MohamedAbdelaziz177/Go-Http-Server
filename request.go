package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strconv"
	"strings"
)

type Request struct {
	Proto   string
	Method  string
	Path    string
	Headers map[string]string
	Params  map[string]string
	Body    []byte
}

func NewRequest(proto string, method string, path string, headers map[string]string, params map[string]string, body []byte) *Request {
	return &Request{
		Proto:   proto,
		Method:  method,
		Path:    path,
		Headers: headers,
		Params:  params,
		Body:    body,
	}
}

func ParseRequest(conn net.Conn) (*Request, error) {

	var reader *bufio.Reader = bufio.NewReader(conn)

	line, err := reader.ReadString('\n')

	if err != nil {
		return nil, err
	}

	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")

	if len(parts) != 3 {
		return nil, fmt.Errorf("invalid request line: %s", line)
	}

	method := strings.ToUpper(parts[0])
	endpoint := parts[1]
	proto := strings.ToUpper(parts[2])

	path, params := parsePathAndQueryParamsFrom(endpoint)
	headers, err := parseHeaders(reader)

	if err != nil {
		return nil, err
	}

	if cl, ok := headers["content-length"]; ok {

		length, err := strconv.Atoi(cl)

		if err != nil {
			return nil, fmt.Errorf("invalid content-length: %s", cl)
		}

		body, err := parseBody(reader, length)

		if err != nil {
			return nil, err
		}

		return NewRequest(proto, method, path, headers, params, body), nil
	}

	return NewRequest(proto, method, path, headers, params, nil), nil
}

func parsePathAndQueryParamsFrom(url string) (string, map[string]string) {

	var path, query string

	if epParts := strings.Split(url, "?"); len(epParts) == 2 {
		path = epParts[0]
		query = epParts[1]
	} else {
		path = url
	}

	params := make(map[string]string)

	if query != "" {

		queryParamsAsString := strings.Split(query, "&")

		for _, qp := range queryParamsAsString {

			keyValuePair := strings.SplitN(qp, "=", 2)

			if len(keyValuePair) == 2 {
				params[strings.TrimSpace(keyValuePair[0])] = strings.TrimSpace(keyValuePair[1])

			} else if len(keyValuePair) == 1 {
				params[strings.TrimSpace(keyValuePair[0])] = ""
			}
		}
	}

	return path, params
}

func parseHeaders(reader *bufio.Reader) (map[string]string, error) {

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

		headers[strings.ToLower(strings.TrimSpace(parts[0]))] = strings.TrimSpace(parts[1])
	}

	return headers, nil
}

func parseBody(reader *bufio.Reader, length int) ([]byte, error) {

	body := make([]byte, length)
	_, err := io.ReadFull(reader, body)

	if err != nil {
		return nil, fmt.Errorf("invalid body: %s", err)
	}

	return body, nil
}
