package main

import (
	"fmt"
	model "goestudo/models"
	"time"
)



func main() {

	// fmt.Println("Starting the application...")

	// address := model.Address{
	// 	Street:  "Rua x",
	// 	City:    "Campo grande",
	// 	ZipCode: "7905120",
	// }
	// address.Street = "Rua Y"

	// person := model.Person{
	// 	Name:  "Marlom",
	// 	Address:  address,
	// 	Birthday: time.Date(2000,07,22,0,0,0,0, time.Local),
	// }
	// person.CalcAge()


	// automovelMoto :=  model.Moto{
	// 	Automovel: model.Automovel{
			
	// 			Ano: 2022,
	// 			Placa: "xxx=pto2",
	// 			Modelo: "factor2",
			
	// 	},
	// 	Cilindradas: 125,
	// }
	var mercado = model.NewCompra(time.Now())
	
	
	 tenis := model.NewItem("Tenis nike", 200, 1)

	 

	 arroz := model.NewItem("Arroz indiano",10.0, 2,);

	mercado.AdicionarItem(*tenis)
	mercado.AdicionarItem(*arroz)
	mercado.CalcularTotal()

	fmt.Println(mercado.ValorTotal)
	mercado.ImprimirLista()


}