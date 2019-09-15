package main

import (
	"fmt"
	"reflect"
)

// ARRAYS
//collection of variables of the same type -hashable (key:value)
// apos declarados, nao podem aumentar ou diminuir
// tipados por valor(copias), porem podem ser manipulados por referencia via ponteiros

//func main() {
//	var balance [5] int
//	balance = [5]int{1: 10, 3: 30}
//	fmt.Printf("%v, %T \n", balance,balance)
//	balance = [5]int{50,40,30,20,10}
//	fmt.Printf("%v, %T \n", balance,balance)
//	// declaração + inicialização + inferencia
//	testeArray := [4]string{"alice","bob", "foo", "bar"}
//	fmt.Println(testeArray[3])
//	fmt.Println(testeArray[:2])  // [inicio: fim] exclusivo
//	fmt.Println(testeArray[1:2])  // [inicio: fim] exclusivo
//
//	// implicit - ellipses
//	x := [...]int{10, 20, 30}
//	fmt.Printf("%v, %T \n", x,x)

//  multidimensoes
// a = [3][4]int{
//   {0, 1, 2, 3} ,   /*  initializers for row indexed by 0 */
//   {4, 5, 6, 7} ,   /*  initializers for row indexed by 1 */
//   {8, 9, 10, 11}   /*  initializers for row indexed by 2 */
//}

//}

// SLICES (FATIAS)
// abstração do array
// estruturas flexível e extensivel para manipulação de coleçao de dados
// composto de elementos do mesmo tipo
// flexible and extensible data structure to implement and manage collections of data. Slices are made up of multiple elements, all of the same type. A slice is a segment of dynamic arrays that can grow and shrink as you see fit. Like arrays, slices are index-able and have a length. Slices have a capacity and length property.
//  utilizam os parametros length(tamanho atual) e capacity(capacidade total)
// zero value of a slice is nil
// imagens: https://blog.golang.org/go-slices-usage-and-internals

func main() {
	//var intSlice []int
	//fmt.Printf("%v, %T \n", intSlice, intSlice)
	//fmt.Println(reflect.ValueOf(intSlice).Kind())

	// inicializacao literal
	//var strSlice = []string{"India", "Canada", "Japan"}
	//fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))

	// inicializacao via make()
	//var strSlice = make([]string, 10, 20) // when length and capacity is different
	//fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
	//fmt.Println(reflect.ValueOf(strSlice).Kind())

	// inicializacao new()
	var intSlice = new([50]int)[0:10]
	fmt.Println(reflect.ValueOf(intSlice).Kind())
	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(intSlice)

	// diferença new e make
	// make only makes slices, maps, and channels - no pointers - return T - good for Create a slice with space preallocated or with len != cap
	// it returns an initialized (not zeroed) value of type T (not *T), and slices must be initialized before execution

	//	new only returns pointers to initialised memory. - returns *T
	// result = new(T) same as
	// var temp T
	// result = &temp

	// usando append para adicionar novos intens ao slice
	//a := make([]int, 2, 5)
	//b :=  make([]int, 2, 5)
	//a = append(a, 30, 40, 50)
	//fmt.Println("Slice A after appending data:", a)
	//fmt.Printf("Length is %d Capacity is %d\n", len(a), cap(a))
	//
	//c := append(a, b...)  //usage of triple-dot ... ellipsis used to append a slice
	//fmt.Printf("Length is %d Capacity is %d\n", len(c), cap(c))
	//for index, values := range c{
	//	fmt.Println(index, "-", values, "/n")
	//}

	// add copy and remove
}
