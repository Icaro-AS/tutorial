package main

import "github.com/ias-epf/tutorial/internal/entity"

type Car struct {
	Model string
	Color string
}

// metodo
func (c Car) Start() {
	println(c.Model + " is starting...")
}

func main() {

	order, err := entity.NewOrder("1", -10, 1)
	if err != nil {
		println(err.Error())
	} else {
		println(order.ID)
	}

}
