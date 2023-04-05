package core

import (
	"fmt"
	"reflect"
)

type ProductReflect struct {
}


func (ref *ProductReflect) GetFields(p Product) {
	typeProduct := reflect.TypeOf(p)
	for i := 0; i < typeProduct.NumField(); i++ {
		field := typeProduct.Field(i)
		fmt.Println(field)
	}
}

func (ref *ProductReflect) GetMethods(p Product) {
	typeProduct := reflect.TypeOf(p)
	for i := 0; i < typeProduct.NumMethod(); i++ {
		method := typeProduct.Method(i)
		fmt.Println(method.Name)
	}
}
	

func (ref *ProductReflect) GetPrivateMethods(p Product) {
	
}



