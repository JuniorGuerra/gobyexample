package main

import "fmt"

// recursividad...
// un ejemplo seria recorrer todos los elementos de unas carpetas, y si adentro hay mas supcartetas
// hay que revisar mas subcarpetas...
//Torres de hanoi

func hanoi(n, source, dest, aux int) {

	if n > 0 {
		hanoi(n-1, source, aux, dest)
		fmt.Println("Disco from postiton", n, " to ", dest)
		hanoi(n-1, source, dest, aux)
	}

}

func main() {
	// example(1)
	// countdown(5)
	i := 0
	forRecursivo(i, 10, 1)
	//fmt.Println(factorial(5))
	// hanoi(4, 1, 3, 2)
}

func forRecursivo(declaration int, limit int, increment int) {
	fmt.Println(declaration, limit, increment)
	declaration += increment
	if declaration < limit {
		forRecursivo(declaration, limit, increment)
	}
}

// 5 -> X = 5 * fact(4)
// 5 -> X = 4 * fact(3)

func factorial(number int) int {

	if number > 1 {
		fmt.Println("n -> X = ", number, " * fact -> n - 1 = ", number-1, " = ", factorial(number-1))
		number = number * factorial(number-1)
	}

	return number
}

// number = 5
// number = 4
// number = 3
// number = ...

func countdown(number int) {
	fmt.Println(number)

	if number != 0 {
		countdown(number - 1)
	} else {
		fmt.Println("BOOM!!!!!!!!!!")
	}
}

// LAS LLAMADAS A LA FUNCION DE RECURSIVIDAD SE LLAMAN DESPUES DE UNA CONDICION!!!
// siempre se van a ejecutar completamente la funcion una vez termine su condicion...
func example(i int) {
	i++
	if i == 2 {
		example(i)
	}
	if i == 3 {
		example(i)
	}
	fmt.Println(i)
}
