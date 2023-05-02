package main

import (
	"fmt"

	"Alura_GoPOO/contas"
)

func main() {
	titularConta := 
	contaDoGuilherme := contas.ContaCorrente{Titular: { Nome:"Guilherme"} , NumeroAgencia: 589, NumeroConta: 123456, Saldo: 125.5}
	//contaDaBruna := contas.ContaCorrente{"Bruna", 222, 111222, 200}
	fmt.Println(contaDoGuilherme)
	//fmt.Println(contaDaBruna)

	var contaDaCris *contas.ContaCorrente
	contaDaCris = new(contas.ContaCorrente)
	contaDaCris.Titular = "Cris"
	contaDaCris.NumeroAgencia = 500

	var contaDaCris2 *contas.ContaCorrente
	contaDaCris2 = new(contas.ContaCorrente)
	contaDaCris2.Titular = "Cris"
	contaDaCris2.NumeroAgencia = 500

	fmt.Println(*contaDaCris)
	fmt.Println(*contaDaCris2)
	fmt.Println(*contaDaCris == *contaDaCris2)

	contaDaCris.Saldo = 11
	fmt.Println(contaDaCris.Sacar(10))
	status, valor := contaDaCris.Depositar(500)
	fmt.Println(contaDaCris.Saldo)

	fmt.Println(status, valor)

}
