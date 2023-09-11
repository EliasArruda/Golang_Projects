package userinformation

import (
	"fmt"
	Clear "main/Functions"
	investimento "main/Investimento"
	structs "main/Structs"
	"time"
)

func Options ( Options *structs.Person , Select *int8 ) {
	
	fmt.Println("Welcome: " , Options.Name);	
	fmt.Print("O que você deseja?:\n\n");
	fmt.Print("[0] Ver meu Saldo\n[1] Cadastrar Minha Empresa\n[2] Gastos do Mês\n\n");
	fmt.Print("Escolha: ");
	fmt.Scan(Select);
}

func UserInformation (person *structs.Person) {
	
	var input int8;
	
	for {
		var ReturnOptions = func(){
			Options (person , &input);
		};

		
		if person == nil {
			person = &structs.Person{
				Profissional: structs.Profissional{
					Empresa: "",
					Money: 0,
					Gastos: 0,
					Investimentos: 0,
				},
			};
		};
			
		ReturnOptions ();
			
		switch input {
			
			case 0:
			{
				Clear.ClearTerminal ();
				
				fmt.Println("O Saldo da sua conta é: R$" , person.Money);
				fmt.Print("Como Podemos te ajudar?: \n\n");
				fmt.Print("[0]Pedir Empréstimo\n[1]Investir\n[2]Sair\n\n");
				fmt.Print("Escolha: ");
				fmt.Scan(&input);

				if input == 0 {
					
					Clear.ClearTerminal();
					var setValue int;

					fmt.Print("Qual seria o valor?: ");
					fmt.Scan(&setValue);
					person.Money = person.Money + setValue;

					fmt.Print("Empréstimo feito com sucesso....");
					time.Sleep(3*time.Second);
					Clear.ClearTerminal();
					continue;

				} else if input == 1 {
					
					Clear.ClearTerminal();
					
					var setInvestimento int;

					fmt.Print("Na onde irá investir?:\n\n");
					fmt.Print("[0]Roupas / Calçados\n[1]Maquiagem\n[2]Automóveis\n[3]Comida\n[4]Agricultura\n\n");
					fmt.Print("Escolha: ")
					fmt.Scan(&setInvestimento);

					if setInvestimento == 0 {

						investimento.Investimentos (

							setInvestimento, 
							&person.Tipos_Investimento.RoupasCalçados, 
							&person.Tipos_Investimento.RoupasCalçados,
							&person.Money,
							"Seu saldo de investimento em [ Roupas / Calçados ]: R$",
						);
						
						continue;

					} else if setInvestimento == 1 {
						
						investimento.Investimentos (

							setInvestimento, 
							&person.Tipos_Investimento.Maquiagem, 
							&person.Tipos_Investimento.Maquiagem,
							&person.Money,
							"Seu saldo de investimento em [ Maquiagem ]: R$",
						);
						
						continue;

					} else if setInvestimento == 2 {
						
						investimento.Investimentos (

							setInvestimento, 
							&person.Tipos_Investimento.Automóveis, 
							&person.Tipos_Investimento.Automóveis,
							&person.Money,
							"Seu saldo de investimento em [ Automóveis ]: R$",
						);
						
						continue;

					} else if setInvestimento == 3 {
						
						investimento.Investimentos (

							setInvestimento, 
							&person.Tipos_Investimento.Comida, 
							&person.Tipos_Investimento.Comida,
							&person.Money,
							"Seu saldo de investimento em [ Comida ]: R$",
						);

						continue;

					} else if setInvestimento == 4 {
						
						investimento.Investimentos (

							setInvestimento, 
							&person.Tipos_Investimento.Agricultura, 
							&person.Tipos_Investimento.Agricultura,
							&person.Money,
							"Seu saldo de investimento em [ Agricultura ]: R$",
						);

						continue;

					}
					continue;
				}
				
				Clear.ClearTerminal();
				continue;
			};

			case 1:
			{
				if len(person.Empresa) >= 1 {

					Clear.ClearTerminal();
					fmt.Printf("Empresa já existente: %v.... " , person.Empresa);
					time.Sleep(3*time.Second);
					Clear.ClearTerminal();
					continue;
				}

				Clear.ClearTerminal();
				
				var setEmpresaValue string;
				fmt.Print("Digite o nome da sua Empresa: ")
				fmt.Scan(&setEmpresaValue);
				person.Empresa = setEmpresaValue;
				
				Clear.ClearTerminal();

				fmt.Printf("Empresa: %v , criado com sucesso....\n\n" , person.Empresa );
				time.Sleep(3*time.Second);
				Clear.ClearTerminal();
				continue;
			};

			case 2:
			{
				Clear.ClearTerminal();

				var CalculaGasto int = 
				person.RoupasCalçados + person.Maquiagem + person.Automóveis + 
				person.Comida + person.Agricultura;

				fmt.Printf("Nome: %v\n" , person.Name);
				fmt.Printf("Empresa: %v\n\n" , person.Empresa);
				fmt.Printf("Saldo: R$%v\n" , person.Money);
				fmt.Printf("Gasto: R$%v\n", CalculaGasto);
				fmt.Printf("Seus Investimentos: \n{\n  Roupas / Calçados: R$ %v\n  Maquiagem: R$ %v\n  Automóveis: R$ %v\n  Comida: R$ %v\n  Agricultura: R$ %v\n}", 
					person.RoupasCalçados,
					person.Maquiagem,
					person.Automóveis,
					person.Comida,
					person.Agricultura,
				);

				fmt.Print("\n\n[0]Sair: ")
				fmt.Scan(&input);
				Clear.ClearTerminal();
				continue;
			};

			default:
			{
				Clear.ClearTerminal();
				continue;
			};
		};
	}
}