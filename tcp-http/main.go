package main

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"syscall"
)

type httpReq struct {
	path     string
	method   string
	protocol string
	headers  map[string]string
	body     string
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to create listener", err.Error())
	}
	defer l.Close()

	fmt.Println("listening for clients")
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("failed to accept connection", err.Error())
		}

		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	//conn.SetReadDeadline(time.Now().Add(3 * time.Second))
	defer conn.Close()
	body := `{"healthy": true}`
	req(conn, func(r httpReq) {
		fmt.Println("req:", r)
		res(conn, body)
	})
}

func req(conn net.Conn, handler func(httpReq)) {
	scanner := bufio.NewScanner(conn)
	scanner.Split(func(data []byte, atEOF bool) (int, []byte, error) {
		advance, token, err := bufio.ScanLines(data, atEOF)
		if i := bytes.IndexByte(data, '\n'); i < 0 {
			return 0, data, bufio.ErrFinalToken
		}
		return advance, token, err
	})
	r := httpReq{
		headers: map[string]string{},
	}
	var i, newlines, bodyLength, contentLength int
	for scanner.Scan() {
		l := scanner.Text()
		if i == 0 {
			reqLine := strings.Split(l, " ")
			if len(reqLine) < 3 {
				return
			}
			r.method = reqLine[0]
			r.path = reqLine[1]
			r.protocol = reqLine[2]
		}
		if l == "" {
			newlines++
		}
		// PARSE HEADERS
		if i > 0 && newlines == 0 {
			header := strings.SplitN(l, ":", 2)
			name := header[0]
			value := strings.TrimSpace(header[1])
			r.headers[name] = value
			if name == "Content-Length" {
				length, err := strconv.Atoi(value)
				contentLength = length
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		// PARSE REQUEST BODY
		if newlines == 1 && l != "" {
			r.body += l
			bodyLength += len(l) + 1
		}
		if newlines > 1 {
			bodyLength+=1
			r.body+="\n"+l
		}

		if bodyLength != 0 && bodyLength == contentLength {
			break
		}
		i++
	}
	handler(r)
}

func res(conn net.Conn, body string) {
	// CRLF; carriage-return line-feed
	// STATUS_LINE CRLF
	write(conn, "HTTP/1.1 200 OK\r\n")
	// HEADERS CRLF
	write(conn, fmt.Sprintf("Content-Length: %d\r\n", len(body)))
	write(conn, "Content-Type: application/json\r\n")
	// CRLF
	write(conn, "\r\n")
	// RESPONSE_BODY
	write(conn, body)
}

func write(conn net.Conn, s string) {
	_, err := conn.Write([]byte(s))
	if errors.Is(err, syscall.EPIPE) {
		return
	}
	if err != nil {
		log.Fatal(err)
	}
}
