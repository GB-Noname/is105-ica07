package main
import (
	"fmt"
	"net"
	"bufio"

)

func main() {
	p :=  make([]byte, 2048)
	conn, err := net.Dial("udp", "127.0.0.1:8001")
	if err != nil {
		fmt.Printf("Some error %v", err)
		return
	}
	fmt.Fprintf(conn, "Motatt")
	_, err = bufio.NewReader(conn).Read(p)
	if err == nil {
		fmt.Printf("%s\n", p)
	} else {
		fmt.Printf("Some error %v\n", err)
	}
	conn.Close()
}
