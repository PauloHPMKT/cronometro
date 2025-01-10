package main

import (
	"fmt"

	countdown "github.com/PauloHPMKT/cronometro/application/usecase"
)

type Timer interface {
	Start()
}

func main() {
	// Entrada do usuário
	minutes := getUserInput()

	// Criar uma instancia do cronometro
	cronometer := countdown.NewCountdownTimer(minutes)

	// Iniciar o cronometro
	cronometer.Start()
}

func getUserInput() int {
	defaultTime := 25

	fmt.Println("Digite a quantidade de minutos: ")
	fmt.Scan("%d", &defaultTime)
	return defaultTime
}
