package main

import (
	"fmt"
	//"math"
)

func main(){
	
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	//Earnings before taxes
	ebt := revenue - expenses

	//earnings after taxes = profit
	profit := ebt * (1 -taxRate/100)

	//Ratio
	ratio := ebt/profit

	//concatenacao de string com var, e delimitacao de casas decimais
	fmt.Printf("Total de arrecadacao:  %.2f\n", ebt)

	//print semelhante ao syso("string" + var) de java
	fmt.Println("Total de arrecadacao: ", ebt)

	fmt.Printf("Lucro: %.2f\n", profit)
	fmt.Printf("Razao arrecadacao/lucro: %.2f\n", ratio)
}