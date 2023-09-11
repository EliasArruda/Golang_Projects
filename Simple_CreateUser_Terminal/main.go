package main;

import (
	"fmt"
	useroption "main/Functions/User/User_Option"
	structs "main/Structs"
);

func main () {

	var input string;
	var UserInformation structs.Person;

	fmt.Print("Registrar? ( yes / no ): ");
	fmt.Scan(&input);

	useroption.User_Option(input , &UserInformation);
}