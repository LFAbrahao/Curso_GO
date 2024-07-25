package main

import "fmt"

func calculadoraLucro () {

	var vendas float64
	var gastos float64
	var imposto float64

	fmt.Print("Digite o total de Vendas: ")
	fmt.Scan(&vendas)

	fmt.Print("Digite o total de gastos: ")
	fmt.Scan(&gastos)

	fmt.Print("Digite o total de Impostos: ")
	fmt.Scan(&imposto)

	entradas:=vendas-gastos
	lucro:=vendas-((vendas*imposto)/100)
	

}