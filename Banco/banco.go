package main

import (
	"fmt"
	//"math"
	"os"
	"strconv"
)

const arquivoSaldo = "saldo.txt"

func lerSaldoEmArquivo() float64{
	data, _ := os.ReadFile(arquivoSaldo)
	saldoTexto := string(data)
	saldo, _ := strconv.ParseFloat(saldoTexto,64)
	return saldo
}

func gravarSaldoEmArquivo(saldo float64) {
	saldoTexto := fmt.Sprint(saldo) //fmt.Sprint passa o parametro em uma String
	os.WriteFile(arquivoSaldo,[]byte(saldoTexto), 0644) //[]byte converte o parametro em uma collection de bytes 
}


func main () {

	var saldoConta float64 = lerSaldoEmArquivo()

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
			fmt.Println("Seu saldo é R$ ", saldoConta)

		} else if opcao == 2 {
			fmt.Println("Digie o valor a ser depositado: ")
			var deposito float64
			fmt.Scan(&deposito)

			if deposito<=0 {
				fmt.Println("Deposito deve ser maior que 0")
				continue
			}

			saldoConta = deposito + saldoConta
			gravarSaldoEmArquivo(saldoConta)
			fmt.Println("Valor depositado")
			fmt.Println("Seu saldo é ", saldoConta)
		} else if opcao == 3 {
			fmt.Println("Digite o valor do saque")
			var saque float64
			fmt.Scan(&saque)

			if saque <= 0 {
				fmt.Println("Saque deve ser maior que 0")
				continue
			}

			if saque > saldoConta {
				fmt.Println("Valor maior que o saldo da conta")
				continue
			}
			
			saldoConta = saldoConta - saque
			gravarSaldoEmArquivo(saldoConta)
			fmt.Println("Valor sacado")
			fmt.Println("Seu saldo é ", saldoConta)
		} else {
			fmt.Println("Sistema encerrado")
			break
		}

	}
	fmt.Println("obrigado por usar o BANCO")
}