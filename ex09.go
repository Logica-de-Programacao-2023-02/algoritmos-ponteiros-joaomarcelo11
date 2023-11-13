package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
	Preco  float64
}

func aplicarDesconto(livro *Livro) {
	desconto := livro.Preco * 0.1
	livro.Preco -= desconto
}

func main() {
	livro := Livro{
		Titulo: "Aventuras na Floresta",
		Autor:  "João Silva",
		Preco:  50.0,
	}

	aplicarDesconto(&livro)

	fmt.Println("Livro:", livro.Titulo)
	fmt.Println("Autor:", livro.Autor)
	fmt.Println("Preço com desconto:", livro.Preco) // Saída: Preço com desconto: 45.0
}