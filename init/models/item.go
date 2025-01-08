package models

import "errors"

type Item struct {
	Nome       string
	Valor      float64
	Quantidade int
}

func NewItem(nome string, valor float64, quantidade int) (*Item, error) {

	if nome == "" {
		return nil, errors.New("Nome é obrigatorio")
	}

	item := new(Item)
	item.Nome = nome
	item.Valor = valor
	item.Quantidade = quantidade
	return item,nil
}