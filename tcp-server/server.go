package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"strings"
)

var host string = "127.0.0.1:9999"

func process(conn net.Conn) {
	client_ip := conn.RemoteAddr()
	log.Println(client_ip, "connected")
	defer conn.Close() // 关闭连接
	for {
		// 读取数据
		reader := bufio.NewReader(conn)
		data, err := reader.ReadString('\n')
		switch err {
		case nil:
			data := strings.TrimSpace(data)
			log.Println(client_ip, "data:", data)
			response := "hello " + data + "\n"
			conn.Write([]byte(response)) // 发送数据
		case io.EOF:
			log.Println(client_ip, "client closed the connection by terminating the process.")
			return
		default:
			log.Printf("error: %v\n", err)
			return
		}
	}
}

func main() {
	listen, err := net.Listen("tcp", host)
	if err != nil {
		log.Fatalln("listen failed, err:", err)
	}
	log.Println("listen on:", host)
	for {
		conn, err := listen.Accept() // 建立连接
		if err != nil {
			log.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
