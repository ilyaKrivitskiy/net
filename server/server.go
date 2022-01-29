package main

import (
	"fmt"
	"log"
	"net"
)

var dict = map[string]string{
	"red":    "красный",
	"orange": "оранжевый",
	"yellow": "желтый",
	"green":  "зеленый",
	"blue":   "синий",
	"purple": "фиолетовый",
}

func main() {

	lst, err := net.Listen("tcp", "127.0.0.1:4545")
	if err != nil {
		log.Fatalln(err)
	}
	defer lst.Close()

	fmt.Println("Server is listening...")

	for {
		conn, err := lst.Accept()
		if err != nil {
			conn.Close()
			fmt.Println(err)
			continue
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()
	for {
		input := make([]byte, 1024*4)
		n, err := conn.Read(input)
		if n == 0 || err != nil {
			fmt.Println("Invalid data...", err)
			break
		}
		source := string(input[:n])
		item, ok := dict[source]
		if !ok {
			item = "undefined"
		}
		fmt.Println(source, "-", item)
		conn.Write([]byte(item))
	}
}
