package main

import "fmt"

type food interface {
	GetName() string
}

type drink struct {
	name string
}

func (d *drink) GetName() string {
	return d.name
}

type barbecue struct {
	name string
}

func (b *barbecue) GetName() string {
	return b.name
}

type stew struct {
	name string
}

func (s *stew) GetName() string {
	return s.name
}

type factory interface {
	Cook(foodName string) food
}

type drinkFactory struct{}

func (df *drinkFactory) Cook(foodName string) food {
	return &drink{name: foodName}
}

type barbecueFactory struct{}

func (bf *barbecueFactory) Cook(foodName string) food {
	return &drink{name: foodName}
}

type stewFactory struct{}

func (sf *stewFactory) Cook(foodName string) food {
	return &drink{name: foodName}
}

func main() {
	var f factory
	f = &drinkFactory{}
	output(f, "Juice")
	f = &barbecueFactory{}
	output(f, "BBQ")
	f = &stewFactory{}
	output(f, "suit")
}

func output(f factory, foodName string) {
	var foodVar food = cook(f, foodName)
	fmt.Printf("drink: %s, type: %T\n", foodVar.GetName(), foodVar)
}

func cook(f factory, foodName string) food {
	return f.Cook(foodName)
}
