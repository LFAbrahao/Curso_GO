package main

import (
	"fmt"
	//"math"
	"os"
	"strconv"
	"errors"
)

const arquivoSaldo = "saldo.txt"

func lerSaldoEmArquivo() (float64, error){
	data, err := os.ReadFile(arquivoSaldo)

	if err != nil{
		return 1000, errors.New("Nao foi possivel ler arquivo de Saldo")
	}

	saldoTexto := string(data)
	saldo, err := strconv.ParseFloat(saldoTexto,64)

	if err != nil {
		return 1000, errors.New("Erro no parse do valor")
	}

	return saldo, nil
}

func gravarSaldoEmArquivo(saldo float64) {
	saldoTexto := fmt.Sprint(saldo) //fmt.Sprint passa o parametro em uma String
	os.WriteFile(arquivoSaldo,[]byte(saldoTexto), 0644) //[]byte converte o parametro em uma collection de bytes 
}


func main () {

	var saldoConta, err = lerSaldoEmArquivo()

	if err != nil{
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("---------")
	}

	fmt.Println("Bem vindo ao sistema bancario!")

	for {
		apresentarMenu()

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