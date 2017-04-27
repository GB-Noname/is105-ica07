package main

import (
	"fmt"
	"net"
)

//Connect udp
conn, err := net.Dial("udp", "host:port")
if err != nil {
return err
}
defer conn.Close()

//simple Read
buffer := make([]byte, 1024)
conn.Read(buffer)

//simple write
conn.Write([]byte("Hello from client"))
