package main

import (
	"net"
	"fmt"
	"log"
)

func handleConnection(c net.Conn) {
	//some code...

	//Simple read from connection
	buffer := make([]byte, 1024)
	c.Read(buffer)

	//simple write to connection
	c.Write([]byte("Hello from server"))

	// listen to incoming udp packets
	pc, err := net.ListenPacket("udp", "host:port")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()

	//simple read
	buffer := make([]byte, 1024)
	pc.ReadFrom(buffer)

	//simple write
	pc.WriteTo([]byte("Hello from client"), addr)
