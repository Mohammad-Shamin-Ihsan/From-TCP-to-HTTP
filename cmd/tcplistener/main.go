package main

import (
	"bytes"
	"fmt"
	"io"
	"net"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	lines := make(chan string)

	go func() {
		defer f.Close()
		defer close(lines)

		buffer := make([]byte, 8)
		var line []byte

		for {
			n, err := f.Read(buffer)

			if err != nil && err != io.EOF {
				panic(err)
			}

			line = append(line, buffer[:n]...)

			for {
				index := bytes.IndexByte(line, '\n')
				if index == -1 {
					break
				}

				lines <- string(line[:index])
				line = line[index+1:]
			}

			if err == io.EOF {
				if len(line) > 0 {
					lines <- string(line)
				}
				break
			}
		}
	}()

	return lines
}

func main() {
	listener, err := net.Listen("tcp", ":42069")
	if err != nil {
		panic(err)
	}
	defer listener.Close()

	fmt.Println("Listening on :42069")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		fmt.Println("Connection accepted")

		lines := getLinesChannel(conn)

		for line := range lines {
			fmt.Println(line)
		}

		fmt.Println("Connection closed")
	}
}
