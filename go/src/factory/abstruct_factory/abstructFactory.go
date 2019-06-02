package main

import "fmt"


type iFactory interface {
	create(factoryType string) iFoodFactory
}

type iFoodFactory interface {
	create(foodName string) iFood
}

type iFood interface {
	getName() string
}

type factoryStruct struct{}

func (fs factoryStruct) create(factoryType string) iFoodFactory {
	foodFactory := map[string]iFoodFactory{
		"drink": drinkFactory{},
	}

	return foodFactory[factoryType]
}

type drinkFactory struct {}

func (df drinkFactory) create(foodName string) iFood {
	return drink{
		food: food{
			name: foodName,
		},
	}
}

type barbecueFactory struct {}

func (df barbecueFactory) create(foodName string) iFood {
	return barbecue{
		food: food{
			name: foodName,
		},
	}
}

type stewFactory struct {}

func (df stewFactory) create(foodName string) iFood {
	return stew{
		food: food{
			name: foodName,
		},
	}
}

type food struct {
	name string
}

type drink struct {
	food
}

func (d drink) getName() string {
	return d.food.name
}

type barbecue struct {
	food
}

func (d barbecue) getName() string {
	return d.food.name
}

type stew struct {
	food
}


func (d stew) getName() string {
	return d.food.name
}

func main() {
	var factory iFactory = factoryStruct{}
	var foodFactory iFoodFactory = factory.create("drink")
	var juice iFood = foodFactory.create("juice")
	fmt.Println(juice.getName())
}