package main

import (
	"fmt"
	"net"
	"protocol"
	"log"
	"strconv"
	"time"
)

func sender(conn net.Conn,ch chan int) {
	for i := 0; i < 10; i++ {
		words := `{"state":`+ strconv.Itoa(i) + `,"msg": "下单成功","data": {"order_sn": "201711091746352997","order_amount": "200","shipping_price": "12"}}`
		data := protocol.Pack([]byte(words))
		conn.Write(data)
	}
	ch <-1
}

func main() {

	ch := make(chan int)
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	fmt.Println("connect success")

	go sender(conn,ch)

	<-ch

	time.Sleep(8*time.Second)
}