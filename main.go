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

func (c *ContaCorrente) Depositar(valorDoDeposito ...float64)  {

    for _, valor := range valorDoDeposito {
        c.saldo += valor
    }

}

func (c *ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
    if valorDaTransferencia < c.saldo && valorDaTransferencia > 0 {
        c.saldo -= valorDaTransferencia 
        contaDestino.Depositar(valorDaTransferencia)
        return true
    } else {
        return false
    }
}

func main() {
    conta := ContaCorrente{"Luiz", 10, 10, 500000.0}
    fmt.Println(conta)

    fmt.Println(conta.saldo)

	fmt.Println(conta.Sacar(400))
    
	fmt.Println(conta.saldo)
    conta.Depositar(100, 100, 200)
    fmt.Println(conta.saldo)
    //var conta2 *ContaCorrente
    //conta2 = new(ContaCorrente)
    //conta2.titular = "Luiz"
    //fmt.Println(*conta2)

    //fmt.Println("Memory address conta2: ", &conta2)
}