package main

import "fmt"

type ContaCorrente struct {
    titular       string
    numeroAgencia int
    numeroConta   int
    saldo         float64
}

func main() {
    conta := ContaCorrente{"Luiz", 10, 10, 500000.0}
    fmt.Println(conta)

    var conta2 *ContaCorrente
    conta2 = new(ContaCorrente)
    conta2.titular = "Luiz"
    fmt.Println(*conta2)

    fmt.Println("Memory address conta2: ", &conta2)
}