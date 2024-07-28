package main

import (
	"fmt"
	//"math"
)


func main () {

	var saldo float64

	fmt.Println("Bem vindo ao sistema bancario!")

	for {
		fmt.Println("Escolha uma opcao: ")
		fmt.Println("1 - Saldo")
		fmt.Println("2 - Depositar")
		fmt.Println("3 - Sacar")
		fmt.Println("4 - Sair")

		var opcao int
		fmt.Println("Digite a opcao desejada: ")
		fmt.Scan(&opcao)

		if opcao == 1 {
			fmt.Println("Seu saldo é R$ ", saldo)

		} else if opcao == 2 {
			fmt.Println("Digie o valor a ser depositado: ")
			var deposito float64
			fmt.Scan(&deposito)

			if deposito<=0 {
				fmt.Println("Deposito deve ser maior que 0")
				continue
			}

			saldo = deposito + saldo
			fmt.Println("Valor depositado")
			fmt.Println("Seu saldo é ", saldo)
		} else if opcao == 3 {
			fmt.Println("Digite o valor do saque")
			var saque float64
			fmt.Scan(&saque)

			if saque <= 0 {
				fmt.Println("Saque deve ser maior que 0")
				continue
			}

			if saque > saldo {
				fmt.Println("Valor maior que o saldo da conta")
				continue
			}
			
			saldo = saldo - saque
			fmt.Println("Valor sacado")
			fmt.Println("Seu saldo é ", saldo)
		} else {
			fmt.Println("Sistema encerrado")
			break
		}

	}
	fmt.Println("obrigado por usar o BANCO")
}