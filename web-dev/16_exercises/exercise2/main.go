package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"html/template"
	"bytes"
	"strings"
)

var t *template.Template
var r map[string]string // request status

func init () {
	t = template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	l, er := net.Listen("tcp", ":8080")
	if er != nil {
		log.Fatalln(er)
	}
	defer l.Close()

	for {
		c, er := l.Accept()
		if er != nil {
			log.Fatalln(er)
		}

		go serve(c)
	}
}

func serve(c net.Conn) {

	s := bufio.NewScanner(c)
	i := 0
	r = make(map[string]string)

	for s.Scan() {
		ln := s.Text()
		if i == 0 {
			sl := strings.Fields(ln)
			r["method"] = sl[0]
			r["url"] = sl[1]
			route(c);
		}
		if ln == "" {
			fmt.Println("end of scan")
			break;
		}
		i++
	}

	defer c.Close()
}

func route (c net.Conn) {
	switch {
	case r["method"] == "GET" && r["url"] == "/":
		r["title"] = "Home Page"
		header(c)
	case r["method"] == "GET" && r["url"] == "/about":
		r["title"] = "About Page"
		header(c)
	}
	 
}

func header(c net.Conn) {
	var b bytes.Buffer
	t.ExecuteTemplate(&b, "index.gohtml", r)
	body := b.String();

	io.WriteString(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(c, "Content-Type: text/html\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}
