package main

// declaração a nível de pacote - F12 para ver os tipos


var (
		NomePacote  = "alice"
		IdadePacote byte = 2
)

func main() {

	// estatica - tipagem das variaveis declaradas sao checadas em tempo de compilacao(declaracao explicita) - C, C#, Java , Go
	// dinamica - tipagem das variaveis declaradas sao checadas em tempo de execucao(declaracao implicita) - Javascript, Python, PHP

	//  tipagem forte - seguem regras estritas em tempo de compilacao 
	//  ex. não compilam, erro em operacao entre strings e inteiros
	//  tipagem fraca - ex.: não seguem regras sintáticas estritas de operações aritméticas entre tipos na fase de compilacao 
	//   ex. compila sem erros uma soma de strings e inteiros

	// existem 3 formas de se declarar variaveis 

	//// 1 - declaracao as variáveis e seus tipos explicitamente definidos
	//var idade int
	//var nome string
	//nome = "gustavo"
	//idade = 28
	//fmt.Println(idade,nome)

	// var  i, j, k int;

	// 2 -  declaração com inicializacao inline
	//var idade byte  = 11
	//var nome string =  "bob"
	//fmt.Println(idade,nome)

	// 3 -  declaração com inferencia de tipo
	//nome  :=  "bob"
	//idade := 11
	//fmt.Printf("%v, %T \n", nome,nome)
	//fmt.Printf("%v, %T \n", idade,idade)

	//var a, b, c = 3, 4, "foo"
	//
	//fmt.Println(a)
	//fmt.Println(b)
	//fmt.Println(c)
	//fmt.Printf("a is of type %T\n", a)
	//fmt.Printf("b is of type %T\n", b)
	//fmt.Printf("c is of type %T\n", c)



	// 4 -  declaração nivel pacote
	//fmt.Println(NomePacote,IdadePacote)
	//
	//// shadowing nivel pacote
	//NomePacote  := "bob"
	//var	IdadePacote byte = 15
	//fmt.Println(NomePacote,IdadePacote)


	// Casting - conversão de tipos
	//var nota float64 = 8.8
	//fmt.Printf("%v, %T \n", nota,nota)
	//fmt.Printf("%v, %T \n", int(nota),int(nota))
	//fmt.Printf("%v, %T \n", float32(nota),float32(nota))

	// cuidado para nao perder informacao - float -> int
	//var novaNota int
	//novaNota = int(nota)
	//fmt.Printf("%v, %T \n", novaNota, novaNota)
	//fmt.Printf("%v, %T \n", float32(novaNota),float32(novaNota))

	// conversao de int -> string (strconv) - ASCII - array of bytes
	//o := 111
	////k := 107
	////println(string(o), string(k))
	//StringConversion := strconv.Itoa(o)
	//fmt.Printf("%v, %T \n", StringConversion, StringConversion)



	// lvalue − Expressions that refer to a memory location is called "lvalue" expression.
	// An lvalue may appear as either the left-hand or right-hand side of an assignment.

	// rvalue − The term rvalue refers to a data value that is stored at some address in memory.
	// An rvalue is an expression that cannot have a value assigned to it which means an rvalue may appear on the right
	// - but not left-hand side of an assignment.
	// x = 20.0  valid
	// 10 = 20.0  invalid
	// 10 = 20.0  invalid
}
