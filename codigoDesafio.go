package main

import "fmt"

func main() {
	ListaItens := ListaItens{
		[]Item{
			{"Biscoito", 1, 2321},
			{"Refrigerante", 1, 3000},
			{"pipoca", 1, 3000},
		},
	}

	ListaEmails := []string{"daniel@golang.com", "joão@golang.com", "junior@golang.com"}

	retorno := contaTotal(ListaItens, ListaEmails...)
	fmt.Println(retorno)
}

type Item struct {
	nomeItem   string
	quantidade int
	preço      int
}

type ListaItens struct {
	Itens []Item
}

func contaTotal(l ListaItens, ListaEmails ...string) map[string]int {

	var total int

	mapaTotal := make(map[string]int)

	if len(ListaEmails) == 0 || len(l.Itens) == 0 {
		return mapaTotal
	}

	for _, v := range l.Itens {
		total += v.preço * v.quantidade
	}

	divisao := total / len(ListaEmails)
	resto := total % len(ListaEmails)

	for _, v := range ListaEmails {
		if resto >= 1 {
			mapaTotal[v] = divisao + 1
			resto--
		} else {
			mapaTotal[v] = divisao
		}
	}

	return mapaTotal

}
