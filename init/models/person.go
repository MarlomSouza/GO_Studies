package models

import "time"

type Person struct {
	Name     string
	Address  Address
	Birthday time.Time
	AgeNow	 int
}


func (p *Person) CalcAge()  {
	p.AgeNow = time.Now().Year() - p.Birthday.Year();
}