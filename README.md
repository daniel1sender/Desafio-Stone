# Desafio-Stone
Desafio Stone em GO

Para testar o resultado do programa basta alterar as lista de itens ([]Item) ou a lista de emails (ListaEmails) acrescentando ou excluindo os dados. Para cada item na lista de itens existem três valores (ex: {"Biscoito", 1, 2321}) o primeiro se refere ao nome do produto, o segundo a quantidade desse produto e o terceiro ao preço em centavos. 

ListaItens := ListaItens{
		[]Item{
			{"Biscoito", 1, 2321},
			{"Refrigerante", 1, 3000},
			{"pipoca", 1, 3000},
		},
	}
  
  ListaEmails := []string{"daniel@golang.com", "joão@golang.com", "junior@golang.com"}
  
  
  Ao rodar o programa, ele retornará uma mapa onde as chaves são os emails das pessoas e os valores são o total a pagar por pessoa em centavos.  
  
  
  
