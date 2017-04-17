package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	defer func() {
		//var r interface{}
		//r = recover()
		if r := recover(); r != nil {
			fmt.Println(r, "recover")
		}
	}()

	s, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	defer s.Close()

	for {
		conn, er := s.Accept()

		if er != nil {
			log.Println(er)
			continue
		}

		go func(c net.Conn) {
			b := make([]byte, 15)

			fmt.Println("iniciando conexão com", c.RemoteAddr().String())

			c.Read(b)

			fmt.Println(c.RemoteAddr().String(), `está dizendo :"`, string(b)+`"`)

			c.Write([]byte("mensagem recebida"))

			fmt.Println("fechando conexão com", c.RemoteAddr().String())

			c.Close()
		}(conn)
	}

}
