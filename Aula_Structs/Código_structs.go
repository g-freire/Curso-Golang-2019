package main

import (
	"fmt"
	"reflect"
)

// letra maiuscula nos campos e na struct para exportar
type Pessoa struct {
	Nome string
	Idade byte
	Parentes []string
}

type Desenvolvedor struct{
	// embarcaou/ embebbed
	Pessoa
	Linguagens []string
	Especialidade string
	BebeCafé bool
}

type Animal struct{
	Nome string `required max: "100"`
}


func main() {
	// ** o que é uma struct?
	// tipo que aceita outros tipos, maior versatilidade

	// ** criação anonima inline com inicialização
	//developer := struct{Nome string; Idade byte;}{Nome:"alice", Idade:22}
	//fmt.Println(developer)

	// ** struct é tipado por Valor - ou seja, cria cópia
	//outroDeveloper := developer
	//outroDeveloper.Nome = "bob"
	//fmt.Println(developer)
	//fmt.Println(outroDeveloper)

	// **  para alterar sem criar cópia deve-se usar ponteiros - alterar diretamente o endereço de memória
	// outroDeveloper := &developer
	// outroDeveloper.Nome = "bob"
	// fmt.Println(developer)
	// fmt.Println(outroDeveloper)

	// boa prática declarar os campos(fields) - clean, manutenção
	//developer := Pessoa{
	//	Nome: "Gustavo",
	//	Idade: 28,
	//	//Parentes: ["alice", "bob"]
	//	Parentes: []string {
	//		"alice",
	//		"bob",
	//	},
	//}
	//fmt.Println(developer)
	// **manipulação da struct
	//fmt.Println(developer.Nome)
	//fmt.Println(developer.Parentes[0])

	// COMPOSIÇÂO - tipo tem outro tipo
	//dev := Desenvolvedor{}
	//dev.Linguagens = []string{"C", "SQL", "C++"}
	//dev.BebeCafé = true
	//dev.Especialidade = "DBA"
	//dev.Nome = "bob"  // composição da struct Pessoa
	//fmt.Println(dev)

	// outra forma de declarar - sintaxe literal
	// observar que tenho que tipar Pessoa dessa vez
	//dev := Desenvolvedor{
	//	Linguagens:[]string{"C","SQL", "C++"},
	//	BebeCafé:true,
	//	Especialidade:"DBA",
	//	Pessoa: Pessoa{ Nome:"bob"},
	//}
	//fmt.Println(dev)

	// outra forma de declarar - sintaxe literal
	// observar que tenho que tipar Pessoa dessa vez
	//dev := Desenvolvedor{
	//	Linguagens:[]string{"C","SQL", "C++"},
	//	BebeCafé:true,
	//	Especialidade:"DBA",
	//	Pessoa: Pessoa{ Nome:"bob"},
	//}
	//fmt.Println(dev)

	// tag/ limitadores - adicionar reflect package
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Nome")
	fmt.Println(field.Tag)

	cachorro := Animal{Nome:"boris"}
	fmt.Println(cachorro)


}
