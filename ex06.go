package main

import "fmt"

type Livro struct {
	Titulo string
	Autor  string
}

func atualizarTituloLivro(livro *Livro) {
	if livro.Autor == "Anônimo" {
		livro.Titulo = "Desconhecido"
	}
}

func main() {
	livro := Livro{
		Titulo: "Livro A",
		Autor:  "Autor A",
	}
	atualizarTituloLivro(&livro)
	fmt.Println("Título do livro:", livro.Titulo) // Saída: Título do livro: Livro A

	livroAnonimo := Livro{
		Titulo: "Livro B",
		Autor:  "Anônimo",
	}
	atualizarTituloLivro(&livroAnonimo)
	fmt.Println("Título do livro anônimo:", livroAnonimo.Titulo) // Saída: Título do livro anônimo: Desconhecido
}