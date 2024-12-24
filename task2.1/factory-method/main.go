package main

type Product interface{}

type Creator interface{
	// Factory method
	Create() Product
}

type ConcreteCreatorA struct {}
func (c *ConcreteCreatorA) Create() Product {
	return 50
}

type ConcreteCreatorB struct {}
func (c *ConcreteCreatorB) Create() Product {
	return "200"
}

func main() {
	creators := []Creator{&ConcreteCreatorA{}, &ConcreteCreatorB{}}
	products := []Product{}

	for _, creator := range creators {
		products = append(products, creator.Create())
	}
}