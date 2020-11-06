package main

import "fmt"

type produto struct {
	nome     string
	preco    float64
	desconto float64
}

// função com receiver do tipo definido
func (p produto) precoComDesconto() float64 {
	return p.preco * (1 - p.desconto)
}

func (p produto) precoTotal(qtde int) float64 {
	return (p.preco * (1 - p.desconto) * float64(qtde))
}

// func precoComDesconto(p produto) float64 {
// 	return p.preco * (1 - p.desconto)
// }

func main() {
	// declarando produto 1 - forma 1
	var produto1 produto

	produto1 = produto{
		nome:     "lapis",
		preco:    2.50,
		desconto: 0.1,
	}

	// declarando produto 2 - forma 2
	produto2 := produto{
		nome:     "borracha",
		preco:    4.00,
		desconto: 0.20,
	}

	fmt.Println("Estrutura de dados")
	fmt.Println(produto1)
	fmt.Println(produto2)

	fmt.Println("Acionamento dos métodos")
	fmt.Println(produto1.precoComDesconto())
	fmt.Println(produto2.precoComDesconto())
	// fmt.Println(precoComDesconto(produto1))

	fmt.Println("Acionamento dos métodos com parâmetros")
	fmt.Println(produto1.precoTotal(2))
	fmt.Println(produto2.precoTotal(2))
}
