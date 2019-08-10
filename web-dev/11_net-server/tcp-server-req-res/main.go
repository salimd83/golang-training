package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			u := strings.Fields(ln)[1]
			fmt.Println("***URL", u)
			respond(conn, u)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func respond(conn net.Conn, u string) {
	var data string
	switch {
	case u == "/":
		data = "<h1>Hello</h1>"
	case u == "/about":
		data = "<h1>About Us</h1>"
	case u == "/contact":
		data = "<h1>Contact us</h1>"
	default:
		data = "<strong>Page Not Found: <br /> <small>404 error</small></strong>"
	}
	body := fmt.Sprint(`<!DOCTYPE html><html lang="en"><head><title></title></head><body>`, data, `</body></html>`)

	fmt.Fprint(conn, "HTTP/1.1 200 ok\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
