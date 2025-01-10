package main

import (
	"emailn/internal/domain/campaign"

	"github.com/go-playground/validator/v10"
)

func main() {

	contact := []campaign.Contact{{Email: "22@gmail.com"}, {Email: "22"}}
	campaign := campaign.Campaign{Recipients: contact}
	validate := validator.New()
	err := validate.Struct(campaign)

	if err == nil {
		println("Nenhum erro")
	} else {
		validationErros := err.(validator.ValidationErrors)
		for _, v := range validationErros {

			switch v.Tag() {
			case "required":
				println(v.StructField() + " is required")
			case "min":
				println(v.StructField() + " is invalid")
			}

		}
	}
}
