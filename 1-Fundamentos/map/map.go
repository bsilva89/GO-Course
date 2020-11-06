package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[123] = "teste1"
	aprovados[456] = "teste2"

	fmt.Println(aprovados)

	for codigo, nome := range aprovados {
		fmt.Println(codigo, ": ", nome)
		fmt.Println(codigo, ":", nome)
		fmt.Printf("%s (codigo: %d).\n", nome, codigo)
	}

	fmt.Println(aprovados[456])

	delete(aprovados, 456)

	fmt.Println(aprovados[456])

	//--------TESTE 2--------//

	nomeSalario := map[string]float64{
		"teste1": 1234.56,
		"teste2": 5678.90,
	}

	nomeSalario["teste3"] = 345.6
	println(nomeSalario)

	for nome, salario := range nomeSalario {
		fmt.Println(nome, ": ", salario)
	}
	//notar que não mantém a ordem
}
