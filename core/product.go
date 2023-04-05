package core

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
)

type Product struct {
	Id       int16 `json:id`
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

func (p *Product) ToJson() ([]byte, error) {
	jsonProduct, _ := json.Marshal(p)
	return jsonProduct, nil
}

func (p *Product) ToXml() ([]byte, error) {
	xmlProduct, _ := xml.Marshal(p)
	return xmlProduct, nil
}

func (p *Product) ToMap() map[string]interface{} {

	return map[string]interface{}{
		"Id": p.Id,
		"Name": p.Name,
		"Price": p.Price,
		"Quantity": p.Quantity,
	}
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

