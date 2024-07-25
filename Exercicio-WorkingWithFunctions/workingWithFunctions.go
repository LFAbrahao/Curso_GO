package main

import (
	"fmt"
	//"math"
)

func main(){
	
	var revenue float64
	var expenses float64
	var taxRate float64

	//fmt.Print("Revenue: ")
	//fmt.Scan(&revenue)     Duas linhas substituidas pela funcao UserInput
	revenue = UserInput("Revenue: ")

	//fmt.Print("Expenses: ")
	//fmt.Scan(&expenses)
	expenses = UserInput("Expenses: ")

	//fmt.Print("Tax Rate: ")
	//fmt.Scan(&taxRate)
	taxRate = UserInput("Tax Rate: ")

	//Earnings before taxes
	//ebt := revenue - expenses

	//earnings after taxes = profit
	//profit := ebt * (1 -taxRate/100)

	//Ratio
	//ratio := ebt/profit

	//concatenacao de string com var, e delimitacao de casas decimais
	fmt.Printf("Total de arrecadacao:  %.2f\n", ebt)

	//print semelhante ao syso("string" + var) de java
	fmt.Println("Total de arrecadacao: ", ebt)

	fmt.Printf("Lucro: %.2f\n", profit)
	fmt.Printf("Razao arrecadacao/lucro: %.2f\n", ratio)
}

func UserInput(infoText string) (float64) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}

func calculosFinanceiros() {

}