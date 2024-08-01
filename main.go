package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

// Função para lidar com conexões de cliente
func handleConnection(conn net.Conn, port int) {
	defer conn.Close()
	fmt.Printf("Cliente conectado na porta %d\n", port)

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			if err != io.EOF {
				log.Printf("Erro na porta %d: %s\n", port, err)
			}
			break
		}
		message := buffer[:n]
		fmt.Printf("Recebido na porta %d: %s\n", port, message)
		_, err = conn.Write([]byte(fmt.Sprintf("Resposta %d: %s", port, message)))
		if err != nil {
			log.Printf("Erro ao enviar resposta na porta %d: %s\n", port, err)
			break
		}
	}
	fmt.Printf("Cliente desconectado da porta %d\n", port)
}

// Função para criar e iniciar um servidor TCP em uma porta específica
func startServer(port int, wg *sync.WaitGroup) {
	defer wg.Done()
	addr := fmt.Sprintf(":%d", port)
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Erro ao iniciar o servidor na porta %d: %s\n", port, err)
	}
	defer ln.Close()
	fmt.Printf("Servidor TCP escutando na porta %d\n", port)
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Erro ao aceitar conexão na porta %d: %s\n", port, err)
			continue
		}
		go handleConnection(conn, port)
	}
}

func main() {
	var wg sync.WaitGroup
	for port := 10000; port <= 11000; port++ {
		wg.Add(1)
		go startServer(port, &wg)
	}
	wg.Wait()
}
