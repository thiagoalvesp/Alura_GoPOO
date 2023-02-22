package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "Saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", c.saldo
	} else {
		return "Valor do deposito menor que zero", c.saldo
	}
}

func main() {
	contaDoGuilherme := ContaCorrente{titular: "Guilherme", numeroAgencia: 589, numeroConta: 123456, saldo: 125.5}
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
