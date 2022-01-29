package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:4545")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for {
		var req string
		fmt.Println("Input some colour to translate:")
		_, err := fmt.Scanln(&req)
		if err != nil {
			fmt.Println("Invalid input...\n", err)
			continue
		}
		n, err := conn.Write([]byte(req))
		if n == 0 || err != nil {
			fmt.Println("Error occured.", err)
		}

		buf := make([]byte, 1024*4)
		nbytes, e := conn.Read(buf)
		if nbytes == 0 || e != nil {
			log.Fatalln(e)
		}
		fmt.Print("Translate: ", string(buf[:nbytes]))
		fmt.Println()
	}
}
