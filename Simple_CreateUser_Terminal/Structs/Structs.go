package structs;

type Information struct
{
	Name , Email string;
	Idade , Nascimento int;
	Faculdade , Curso string;
	Password string;
}

type Profissional struct 
{
	Empresa string;
	Money int;
	Gastos float64;
	Investimentos int;
}

type Tipos_Investimento struct
{
	RoupasCalçados int;
	Maquiagem int;
	Automóveis int;
	Comida int;
	Agricultura int;
}

type Person struct 
{
	Information;
	Profissional;
	Tipos_Investimento;
}

