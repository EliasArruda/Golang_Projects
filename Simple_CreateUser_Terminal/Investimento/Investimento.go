package investimento

import (
	"fmt"
	Clear "main/Functions"
	"time"
)

func Investimentos ( 

	input int, 
	person *int, 
	VariávelCalc *int,
	Saldo_Conta *int,
	Comment string,
){

	Clear.ClearTerminal();
	fmt.Print("Saldo Atual: R$ " , *Saldo_Conta , "\n");
	fmt.Print( Comment , *person , "\n\n");
	fmt.Print("Qual será o valor do investimento: ");
	fmt.Scan(&input);
	

	*VariávelCalc = *person + input; 
	*Saldo_Conta = *Saldo_Conta  - input;

	
	Clear.ClearTerminal();	
	fmt.Println("Investimento efetuado com sucesso....");
	time.Sleep(3*time.Second);
	Clear.ClearTerminal();
}