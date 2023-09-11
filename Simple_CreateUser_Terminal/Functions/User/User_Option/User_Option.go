package useroption;

import (
	"fmt"
	Clear "main/Functions"
	userinformation "main/Functions/User/User_Information"
	structs "main/Structs"
)

func User_Option ( inputValue string , person *structs.Person )  ( string , error ) {

	switch inputValue {
		
		case "yes":
		{
			Clear.ClearTerminal();
			
			fmt.Print("Seu Nome: ");
			fmt.Scan(&person.Name);

			fmt.Print("Seu Email: ");
			fmt.Scan(&person.Email);

			fmt.Print("Sua Senha: ");
			fmt.Scan(&person.Password);

			fmt.Print("Sua Idade: ");
			fmt.Scan(&person.Idade);

			fmt.Print("Seu Nascimento: ");
			fmt.Scan(&person.Nascimento);

			fmt.Print("Faz/Fez Faculdade? ( yes / no ): ");
			fmt.Scan(&person.Faculdade);

			if person.Faculdade != "yes" && person.Faculdade != "no" {
				
				for {
					fmt.Print("Alternativa errada , escolha novamente: \n");
					fmt.Print("Faz/Fez Faculdade? ( yes / no ): ");
					fmt.Scan(&person.Faculdade);

					if person.Faculdade == "yes" || person.Faculdade == "no" {
						break;
					}
				}
			}

			fmt.Print("Faz/Fez Curso? ( yes / no ): ");
			fmt.Scan(&person.Curso);

			if person.Curso != "yes" && person.Curso != "no" {
				
				for {
					fmt.Print("Alternativa errada , escolha novamente: \n")
					fmt.Print("Faz/Fez Curso? ( yes / no ): ");
					fmt.Scan(&person.Curso);

					if person.Curso == "yes" || person.Curso == "no" {
						break;
					}
				}
			}
			
			Clear.ClearTerminal();
			userinformation.UserInformation(person);
			break;
		}

		default:
		{
			panic("Programa encerrado");
		}
	}

	return inputValue , nil;
}