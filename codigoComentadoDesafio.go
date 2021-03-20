package main

import "fmt"

func main() {

	//aqui preencho os campos da estrutura ListaItens preenchendo os dados do tipo []Item
	ListaItens := ListaItens{
		[]Item{
			{"Biscoito", 1, 2321},
			{"Refrigerante", 1, 3000},
			{"pipoca", 1, 3000},
		},
	}

	//aqui declaro a ListaEmails como uma slice de strings contendo os emails das pessoas
	ListaEmails := []string{"daniel@golang.com", "joão@golang.com", "junior@golang.com"}

	//Aqui atribuo o retorno da função contaTotal a variavel chamada "retorno"
	retorno := contaTotal(ListaItens, ListaEmails...)

	//Imprimo a variavel retorno
	fmt.Println(retorno)
}

//Aqui declaro a estrutura Item com três tipos de dados referentes ao Item (nomeItem, quantidade, preço)
type Item struct {
	nomeItem   string
	quantidade int
	preço      int
}

//Aqui declaro a estrutura ListaItens com um tipo de dado do tipo slice ([]Item)
type ListaItens struct {
	Itens []Item
}

//A função contaTotal recebe dois parâmetros, um sendo do tipo ListaItens (l ListaItens) e outro do tipo slice of strings (ListaEmails)
//a funçaõ retorna um amapa onde as chaves são os emails (string) e o valor a pagar em centavos (int)
func contaTotal(l ListaItens, ListaEmails ...string) map[string]int {

	//aqui declaro a variável total
	var total int

	//aqui declaro o mapa para ser retornado pela função
	mapaTotal := make(map[string]int)

	//Caso a função receba listas vazias ela retornará um mapa vazio
	if len(ListaEmails) == 0 || len(l.Itens) == 0 {
		return mapaTotal
	}

	//esse loop calcula o total a ser pago da lista de compras, ou seja, o preço de cada item
	//(v.preço) multiplicado pela quantidade de cada item (v.quantidade)
	for _, v := range l.Itens {
		total += v.preço * v.quantidade
	}

	divisao := total / len(ListaEmails)
	resto := total % len(ListaEmails)

	//esse looping atribui ao mapa (mapaTotal) quanto cada "email" (pessoa) deverá pagar
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
