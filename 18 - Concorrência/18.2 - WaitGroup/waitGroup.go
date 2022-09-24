package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(2) // contagem de quantas vezes a goroutine vai executar as funçoes com WaitGroup
	// As funções vão ser executadas de forma concorrente
	go func() {
		escrever("Ola mundo!")
		waitGroup.Done() // -1
	}()

	go func() {
		escrever("Programando em Go")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // Avisa a função para aguardar a contagem da goroutine

	escrever("Olá Mundo!") // goroutine
	escrever("Programando em GO!")
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
