package main

import (
	"fmt"

	"github.com/thiagoalvesp/Alura_GoPOO/contas"
)

func main() {
	contaDoGuilherme := contas.ContaCorrente{titular: "Guilherme", numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}
	contaDaBruna := ContaCorrente{"Bruna", 222, 111222, 200}
	fmt.Println(contaDoGuilherme)
	fmt.Println(contaDaBruna)

	var contaDaCris *ContaCorrente
	contaDaCris = new(ContaCorrente)
	contaDaCris.titular = "Cris"
	contaDaCris.numeroAgencia = 500

	var contaDaCris2 *ContaCorrente
	contaDaCris2 = new(ContaCorrente)
	contaDaCris2.titular = "Cris"
	contaDaCris2.numeroAgencia = 500

	fmt.Println(*contaDaCris)
	fmt.Println(*contaDaCris2)
	fmt.Println(*contaDaCris == *contaDaCris2)

	contaDaCris.saldo = 11
	fmt.Println(contaDaCris.Sacar(10))
	status, valor := contaDaCris.Depositar(500)
	fmt.Println(contaDaCris.saldo)

	fmt.Println(status, valor)

}
