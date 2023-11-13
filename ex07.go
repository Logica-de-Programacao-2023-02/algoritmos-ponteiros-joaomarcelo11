package main

import "fmt"

type Conta struct {
	Saldo   float64
	Titular string
}

func adicionarValorConta(conta *Conta, valor float64) {
	conta.Saldo += valor
}

func main() {
	conta := Conta{
		Saldo:   100.0,
		Titular: "João",
	}
	adicionarValorConta(&conta, 50.0)
	fmt.Println("Saldo da conta:", conta.Saldo) // Saída: Saldo da conta: 150.0
}