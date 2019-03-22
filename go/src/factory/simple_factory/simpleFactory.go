package main

import "fmt"

type food struct {
	Name string
}
type drink struct {
	food
}

type barbecue struct {
	food
}

type stew struct {
	food
}

type foodFactory struct {
}

func (ff *foodFactory) NewDrink(name string) drink {
	drink := drink{}
	drink.Name = name
	return drink
}

func (ff *foodFactory) NewBarbecue(name string) barbecue {
	barbecue := barbecue{}
	barbecue.Name = name
	return barbecue
}

func (ff *foodFactory) NewStew(name string) stew {
	stew := stew{}
	stew.Name = name
	return stew
}

func main() {
	foodFactory := foodFactory{}
	drink := foodFactory.NewDrink("juice")
	barbecue := foodFactory.NewBarbecue("BBQ")
	stew := foodFactory.NewStew("Braised pork")
	fmt.Printf("drink: %s, type: %T\n", drink.Name, drink)
	fmt.Printf("drink: %s, type: %T\n", barbecue.Name, barbecue)
	fmt.Printf("drink: %s, type: %T\n", stew.Name, stew)
}
