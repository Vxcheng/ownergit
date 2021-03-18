package creation

import (
	"fmt"
)

type packing interface {
	pack() string
}

type item interface {
	name() string
	price() float64
	packing() packing
}

type bottle struct{}

func (b *bottle) pack() string {
	return "bottle"
}

type coldDrink struct{}

func (c *coldDrink) packing() packing {
	return new(bottle)
}

type pepsi struct {
	coldDrink
}

func (p *pepsi) name() string {
	return "pepsi"
}

func (p *pepsi) price() float64 {
	return 10.0
}

type wrapper struct{}

func (b *wrapper) pack() string {
	return "wrapper"
}

type burger struct{}

func (c *burger) packing() packing {
	return new(wrapper)
}

type chickenBuiger struct {
	burger
}

func (p *chickenBuiger) name() string {
	return "chickenBuiger"
}

func (p *chickenBuiger) price() float64 {
	return 20.0
}

type Meal struct {
	items []item
}

func (m *Meal) GetCosts() (cost float64) {
	for _, it := range m.items {
		cost += it.price()
	}
	return
}

func (m *Meal) ShowItems() {
	for _, it := range m.items {
		fmt.Printf("item: %s, packing: %s, price: %.2f\n", it.name(), it.packing().pack(), it.price())
	}

}

func (m *Meal) addItem(item item) {
	m.items = append(m.items, item)
}

type MealBuilder struct {
}

func (b *MealBuilder) PrePepsiDrink() *Meal {
	m := new(Meal)
	m.addItem(new(pepsi))
	m.addItem(new(chickenBuiger))

	return m
}
