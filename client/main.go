package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:8080")
	if err != nil {
		log.Fatal("could not resolve tcp address:", err)
	}
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatal("could not dial:", err)
	}
	defer conn.Close()

	// WRITE TO CONNECTION
	reqBody := `{"f": "v"}`
	httpReq := "GET / HTTP/1.1\r\n"
	httpReq += fmt.Sprintf("Content-Length: %d\r\n", len(reqBody))
	httpReq += "Host: localhost:8080\r\n"
	//httpReq += "Connection: close\r\n"
	httpReq += "\r\n"
	httpReq += reqBody
	_, err = conn.Write([]byte(httpReq))
	if err != nil {
		log.Fatal("could not write to connection:", err)
	}
	// Server does not respond unless the connection
	// is either closed by the client or by the server
	// by providing the `Connection: close` header
	conn.CloseWrite()

	//READ FROM CONNECTION
	bs, err := ioutil.ReadAll(conn)
	if err != nil {
		log.Fatal("could not read from connection:", err)
	}
	fmt.Printf("got this response:\n%s", bs)
}
