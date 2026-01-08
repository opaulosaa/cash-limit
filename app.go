package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Gasto representa a estrutura de um item de despesa, como se fosse uma classe em java
type Gasto struct {
	Descricao string
	Valor     float64
}

func main() {
	listaGastos := []Gasto{}              //slice de gastos
	scanner := bufio.NewScanner(os.Stdin) //consegue ler os inputs mesmo com espaço, por exemplo "Compras do Supermercado", "lanche unifor" etc.

	for {
		fmt.Println("\n--- CONTROLE DE GASTOS ---")
		fmt.Println("1. Adicionar gasto")
		fmt.Println("2. Conferir gastos")
		fmt.Println("3. Deletar gastos")
		fmt.Println("4. Sair")
		fmt.Print("Escolha uma opção: ")

		scanner.Scan()
		opcao := scanner.Text()

		switch opcao {
		case "1":
			fmt.Print("O que foi gasto? ")
			scanner.Scan()
			desc := scanner.Text()

			fmt.Print("Quanto custou? (Ex: 50.25): ")
			scanner.Scan()
			valorStr := scanner.Text()
			valor, err := strconv.ParseFloat(valorStr, 64) //tudo que é posto no terminal é texto, essa conversão evita que o fluxo quebre
			if err != nil {
				fmt.Println("Valor inválido! Use números e ponto para centavos.")
				continue
			}

			novoGasto := Gasto{Descricao: desc, Valor: valor}
			listaGastos = append(listaGastos, novoGasto)
			fmt.Println("Gasto adicionado com sucesso!")

		case "2":
			fmt.Println("\n--- LISTA DE GASTOS ---")
			if len(listaGastos) == 0 {
				fmt.Println("Nenhum gasto registrado.")
			} else {
				var total float64
				for i, g := range listaGastos { //g é a cópia do objeto Gasto, lá em cima, em struct
					fmt.Printf("%d. %s: R$ %.2f\n", i+1, g.Descricao, g.Valor) //sempre com 2 casas decimais para o valor
					total += g.Valor
				}
				fmt.Printf("-----------------------\nTotal: R$ %.2f\n", total)
			}

		case "3":
			if len(listaGastos) == 0 {
				fmt.Println("Não há nada para deletar.")
				continue
			}
			fmt.Print("Digite o número do gasto que deseja remover: ")
			scanner.Scan()
			idxStr := scanner.Text()         //input do usuário, como texto
			idx, err := strconv.Atoi(idxStr) //função que converte o string para inteiro
			if err != nil || idx < 1 || idx > len(listaGastos) {
				fmt.Println("Índice inválido!")
			} else {
				// Remove o item do slice
				listaGastos = append(listaGastos[:idx-1], listaGastos[idx:]...)
				fmt.Println("Gasto removido!")
			}

		case "4":
			fmt.Println("Saindo... Até logo!")
			return

		default:
			fmt.Println("Opção inválida, tente novamente.")
		}
	}
}
