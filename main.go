package main

import (
	"fmt"

	"github.com/ortizdavid/reflect-golang/core"
)


func main() {
	
	product := core.Product{
		Id: 1,
		Name: "Computer",
		Price: 1000,
		Quantity: 2,
	}

	fmt.Println("Object ")
	product.ShowData("xml")
	
	///--- Reflection
	var productReflect core.ProductReflect

	fmt.Println("Fields")
	productReflect.GetFields(product)

	fmt.Println("Methods")
	productReflect.GetMethods(product)

	fmt.Println("Public Methods")
	productReflect.GetMethods(product)

}