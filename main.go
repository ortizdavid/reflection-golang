package main

import (
	"github.com/ortizdavid/reflection-golang/core"
)


func main() {
	
	product := core.Product{
		Id: 1,
		Name: "Computer",
		Price: 1000,
		Quantity: 2,
	}

	product.ShowData("string")

}