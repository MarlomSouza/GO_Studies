package models

type Item struct {
	Nome       string
	Valor      float64
	Quantidade int
}

func NewItem(nome string, valor float64, quantidade int) *Item {
	item := new(Item)
	item.Nome = nome
	item.Valor = valor
	item.Quantidade = quantidade
	return item
}