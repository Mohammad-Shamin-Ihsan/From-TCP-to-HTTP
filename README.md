
From TCP to HTTP

A Go project built while following the Boot.dev "From TCP to HTTP" course. The purpose of this project is to understand 
how HTTP works by building its components from the ground up instead of relying on Go's net/http package.
At this stage, the project implements a basic TCP server that accepts incoming connections, reads data as a byte stream,
reconstructs complete lines, and prints the raw HTTP requests received from clients.

---

## Learning Objectives

This project focuses on understanding:

* Go modules and project structure
* File I/O using `os` and `io`
* Byte slices and buffers
* Reading data in fixed-size chunks
* Building complete lines from streamed bytes
* Goroutines
* Channels
* Interfaces (`io.ReadCloser`)
* TCP networking with Go's `net` package
* The structure of raw HTTP requests

---

## Features Implemented

* Read data 8 bytes at a time
* Reconstruct complete lines from streamed data
* Encapsulate parsing logic in a reusable `getLinesChannel()` function
* Process data concurrently using goroutines
* Communicate parsed lines using Go channels
* Create a TCP listener on port **42069**
* Accept multiple TCP connections
* Print incoming HTTP requests exactly as received

---

## Project Structure

```text
.
├── cmd
│   └── tcplistener
│       └── main.go
├── go.mod
└── README.md
```

---

## Running the Project

Clone the repository:

```bash
git clone https://github.com/Mohammad-Shamin-Ihsan/From-TCP-to-HTTP.git
cd From-TCP-to-HTTP
```

Run the TCP listener:

```bash
go run ./cmd/tcplistener
```

Or capture all output to a file:

```bash
go run ./cmd/tcplistener | tee /tmp/request.log
```

---

## Testing with cURL

### GET Request

```bash
curl http://localhost:42069/coffee
```

The server will display the raw HTTP request.

---

### POST Request

```bash
curl -X POST \
-H "Content-Type: application/json" \
-d '{"flavor":"dark mode"}' \
http://localhost:42069/coffee
```

This demonstrates how an HTTP request contains:

* Request line
* Headers
* Blank line
* Request body

---

## Example Output

```http
GET /coffee HTTP/1.1
Host: localhost:42069
User-Agent: curl/8.x.x
Accept: */*
```

Example POST request:

```http
POST /coffee HTTP/1.1
Host: localhost:42069
User-Agent: curl/8.x.x
Content-Type: application/json
Content-Length: 24

{"flavor":"dark mode"}
```

---

## Concepts Learned

* Streams vs. complete messages
* TCP connections
* HTTP request format
* Buffered reading
* Incremental parsing
* Interfaces in Go
* Concurrent programming with goroutines
* Communication using channels

---

## Future Work

The next milestones for this project include:

* Parse HTTP request lines
* Parse HTTP headers
* Parse request bodies using `Content-Length`
* Construct HTTP responses
* Build a minimal HTTP server without using Go's `net/http`
* Support multiple routes and methods

---

## Resources

* Boot.dev — From TCP to HTTP
* Go Programming Language
* RFC 9112 — HTTP/1.1

---

## License

This project is intended for educational purposes while learning Go, networking, and HTTP fundamentals.
