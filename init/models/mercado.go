package models

import (
	"fmt"
	"time"
)

type Mercado struct {
	Itens      []Item
	ValorTotal float64
	DataCompra time.Time
}


func NewCompra(dataCompra time.Time) *Mercado {

	return &Mercado{
		DataCompra : dataCompra,
	}
}

func (mercado *Mercado) AdicionarItem(item Item) {
	mercado.Itens = append(mercado.Itens, item)
}

func (mercado *Mercado) CalcularTotal() {
	var totalTemp float64

	for _, element := range mercado.Itens {
		totalTemp += element.calcularTotalItem()

	}

	mercado.ValorTotal = totalTemp

}

func (item Item) calcularTotalItem() float64 {
	return item.Valor * float64(item.Quantidade)
}

func (mercado *Mercado) ImprimirLista() {
	fmt.Println("Compra feita no dia", mercado.DataCompra)
	for _, m := range mercado.Itens {
		fmt.Println(m.Nome, m.Quantidade, m.Valor, m.calcularTotalItem())
	}
	
}