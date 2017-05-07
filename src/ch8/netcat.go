package main

import (
	"io"
	"log"
	"net"
	"os"
	"sync"
)

func main() {
	servers := []string{"localhost:8001", "localhost:8002", "localhost:8003"}
	var wg sync.WaitGroup
	for _, s := range servers {
		wg.Add(1)
		go func(serv string) {
			conn, err := net.Dial("tcp", serv)
			if err != nil {
				log.Fatal(err)
			}

			if _, err := io.Copy(os.Stdout, conn); err != nil {
				log.Fatal(err)
			}
			conn.Close()
			wg.Done()
		}(s)
	}
	wg.Wait()
}
