package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"net"
	"strconv"
	"time"
)

var ccS = color.New(color.FgGreen, color.Bold) //green
var ccF = color.New(color.FgRed, color.Bold)   //red
func main() {
	port := flag.Int("port", 8080, "set server port , default 8080")
	flag.Parse()
	listener, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(*port))
	fmt.Printf("The server[%s] is starting...\n", strconv.Itoa(*port))
	if err != nil {
		ccF.Println("The server failed to start:", err)
		return
	}

	// todo go server.TransmitMsg() client function calls
	for {
		time.Sleep(20 * time.Millisecond)
		conn, err := listener.Accept()
		if err != nil {
			ccF.Println("The receiving client failed to connect ", err)
			return
		}
		ccS.Printf("[%v]clint connects successfully\n", conn.RemoteAddr())
		//todo Concurrency, one goroutine per client
		//go server.HandleConnect(conn)
	}
}
