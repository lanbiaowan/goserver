package main

import (
	"fmt"
	"net"
	"os"
	"protocol"
	"io"
	"time"
)

func main() {
	netListen, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}

	defer netListen.Close()

	fmt.Println("Waiting for clients")
	for {
		conn, err := netListen.Accept()
		if err != nil {
			continue
		}

		fmt.Println(conn.RemoteAddr().String(), " tcp connect success")
		readchan := make(chan []byte)
		heartchan := make(chan []byte,2)
		go handleConnection(conn,readchan)

		go worker(readchan,heartchan)

		go heartBeat(heartchan,conn)
	}
}

func handleConnection(conn net.Conn,readchan chan []byte) {
	buffer := make([]byte, 1024)
	tmp := make([]byte, 1024)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("client %s is close!\n",conn.RemoteAddr().String())
			}
			return
		}
		buffer = append(buffer,tmp...)
		tmp = protocol.Unpack(buffer[:n],readchan)
	}
}


func worker(message chan []byte,heartchan chan []byte){

	for{
		data := <- message
		heartchan <-data
		fmt.Println(string(data),"\n\n")

	}


}


func  heartBeat(c chan []byte,conn net.Conn){
	for{
		select {
		case <-c:
			fmt.Println("心跳检查")

		case <-time.After(5*time.Second):

			fmt.Println("客户端断开了连接")
			conn.Close()
			return

		}

	}




}