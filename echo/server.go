package main

import (
	"log"
	"net"
)

func main() {

	log.Println("Listening")
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		client, err := listener.Accept()
		if err != nil {
			continue
		}
		handleClient(client)
	}

}

func handleClient(client net.Conn) {
	for {
		buf := make([]byte, 4096)
		// read from client
		numbytes, err := client.Read(buf)
		if numbytes == 0 || err != nil {
			return
		}
		// write to client
		client.Write(buf)
	}
}
