package main

import (
	"fmt"
)

func main() {

	// ** o que é um mapa? estrutura que mapeia chaves a valores
	//populaçãoCidadesRJ := map[string]int64{
	//	"Rio de Janeiro": 6715942,
	//	"São Gonçalo": 6715942,
	//	"Duque de Caxias": 6715942,
	//	"Nova Iguaçu": 6715942,
	//	"Niterói": 511786,
	//}
	//fmt.Println(populaçãoCidadesRJ)

	//keySlice := map[[]int]string{} // slice nao pode ser chave
	//fmt.Println(keySlice)
	//keyArray := map[[3]int]string{}
	//fmt.Println(keyArray)

	// ** criando com make()
	makePopulaçãoCidadesRJ := make(map[string]int)
	makePopulaçãoCidadesRJ = map[string]int{
		"Rio de Janeiro": 6715942,
		"São Gonçalo": 6715942,
		"Duque de Caxias": 6715942,
		"Nova Iguaçu": 6715942,
	}
	fmt.Println(makePopulaçãoCidadesRJ)

	// **  acessando chave
	//log.Println("Após adicionar novo key/value")
	//fmt.Println(makePopulaçãoCidadesRJ["Rio de Janeiro"])
	// ** caso a chave não exista, retornara 0
	//fmt.Println(makePopulaçãoCidadesRJ["Ripo de Janeiro"])
	// ** teste de valores
	//_, ok := makePopulaçãoCidadesRJ["Ripo de Janeiro"]
	//fmt.Println(ok)


	// ** adicionar dados - reparar que ordem não é mantida
	//makePopulaçãoCidadesRJ["Niterói"] = 511786
	//log.Println("Após adicionar novo key/value")
	//fmt.Println(makePopulaçãoCidadesRJ)

	// ** deletar uma chave/valor - método delete()
	// ** mapas usam REFERENCIAS - cuidado com  efeitos colaterais

	//delete(makePopulaçãoCidadesRJ,"Nova Iguaçu" )
	//log.Println("Após deletar")
	//fmt.Println(makePopulaçãoCidadesRJ)

	// ** retorna tamanho do mapa - método len()
	//fmt.Println(len(makePopulaçãoCidadesRJ))


}
