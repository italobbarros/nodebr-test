package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"
)

const (
	startPort         = 10000
	endPort           = 11000
	numClients        = 1000
	messagesPerClient = 500
	message           = "Teste de carga"
	path_result       = "result/node-optimized"
)

// Função para criar um cliente e enviar mensagens
func createClient(port int, wg *sync.WaitGroup) {
	defer wg.Done()
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		log.Printf("Erro na conexão com a porta %d: %s\n", port, err)
		return
	}
	defer conn.Close()

	var totalDuration float64
	var messageCount int

	for i := 0; i < messagesPerClient; i++ {
		startTime := time.Now()

		_, err := conn.Write([]byte(message))
		if err != nil {
			log.Printf("Erro ao enviar mensagem para a porta %d: %s\n", port, err)
			return
		}

		buff := make([]byte, 1024)
		_, err = conn.Read(buff)
		if err != nil {
			log.Printf("Erro ao ler resposta da porta %d: %s\n", port, err)
			return
		}

		elapsed := time.Since(startTime).Seconds()
		totalDuration += elapsed
		messageCount++
	}

	averageTime := totalDuration / float64(messageCount)
	result := map[string]float64{
		"average_response_time": averageTime,
	}

	// Criar diretório se não existir
	dir := fmt.Sprintf(path_result)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		log.Fatalf("Erro ao criar diretório %s: %s\n", dir, err)
	}

	// Salvar os resultados em um arquivo JSON
	filePath := fmt.Sprintf("%s/cliente%d.json", dir, port)
	file, err := os.Create(filePath)
	if err != nil {
		log.Fatalf("Erro ao criar o arquivo %s: %s\n", filePath, err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(result); err != nil {
		log.Fatalf("Erro ao escrever JSON no arquivo %s: %s\n", filePath, err)
	}
}

func main() {
	var wg sync.WaitGroup

	startTime := time.Now()

	for port := startPort; port <= endPort && port < startPort+numClients; port++ {
		wg.Add(1)
		go createClient(port, &wg)
	}

	wg.Wait()
	totalTime := time.Since(startTime).Seconds()

	fmt.Printf("Teste concluído. Tempo total: %.2f segundos\n", totalTime)
}
