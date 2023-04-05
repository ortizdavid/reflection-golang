package core

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Product struct {
	Id       int16 
	Name     string 
	Price    float32
	Quantity int16
}


func (p *Product) isNotAvalilable() bool {
	return p.Quantity == 0
}

func (p *Product) ToString() string {
	return fmt.Sprintf("%d\t%s\t%f\t%d\n", p.Id, p.Name, p.Price, p.Quantity)
}

func (p *Product) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"Id": p.Id,
		"Name": p.Name,
		"Price": p.Price,
		"Quantity": p.Quantity,
	}
}

func (p *Product) ToJson() string {
	result, _ := json.MarshalIndent(p, " ", "  ")
	return string(result)
}

func (p *Product) ToXml() string {
	result, _ := xml.MarshalIndent(p, " ", "  ")
	return string(result)
}

func (p *Product) ShowData(format string) {
	switch format {
	case "string":
		fmt.Println(p.ToString())
	case "json":
		fmt.Println(p.ToJson())
	case "map":
		fmt.Println(p.ToMap())
	case "xml":
		fmt.Println(p.ToXml())
	}
}

