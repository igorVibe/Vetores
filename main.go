package main

import "fmt"

func main() {

	//Slice
	numeros := []int{1, 2, 3, 4, 5}
	fmt.Println(numeros)
	numeros = append(numeros, 6, 7, 8)
	fmt.Println(numeros, len(numeros), cap(numeros))
	

	// Slice
	var nomes = []string{"Igor", "Silva", "Dusk", "Robert"}
	fmt.Println(nomes)
	
	fmt.Println(nomes, len(nomes), cap(nomes))
	rangeOne := nomes[1:3]
	fmt.Println(rangeOne)
	nomes = append(nomes, "Ronald, Landhon")
	fmt.Println(nomes, len(nomes), cap(nomes))

	

}