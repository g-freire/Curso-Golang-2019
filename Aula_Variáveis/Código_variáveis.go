package main

// declaração a nível de pacote - F12 para ver os tipos


var (
		NomePacote  = "alice"
		IdadePacote byte = 2
)

func main() {
	// existem 3 formas de declarar variaveis - escopo da funcao

	//// 1 - declaracao as variáveis e seus tipos explicitamente
	//var idade int
	//var nome string
	//nome = "gustavo"
	//idade = 28
	//fmt.Println(idade,nome)

	// 2 -  declaração com inicializacao inline
	//var idade byte  = 28
	//var nome string =  "gustavo"
	//fmt.Println(idade,nome)

	// 3 -  declaração com inferencia de tipo
	//nome  :=  "gustavo"
	//idade := 28
	//fmt.Printf("%v, %T \n", nome,nome)
	//fmt.Printf("%v, %T \n", idade,idade)

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

}
