package main

import (
	"awesomeProject3/design_patterns/solid"
	"fmt"
)

func main() {
	apple := solid.Product{"Apple", solid.Green, solid.Small}
	tree := solid.Product{"Tree", solid.Green, solid.Large}
	house := solid.Product{"TentHouse", solid.Blue, solid.Large}

	products := []solid.Product{apple, tree, house}

	filter := solid.Filter{}
	for _, v := range *filter.FilterByColour(products, solid.Blue){

		fmt.Println(v)

	}
	fmt.Println("Done with execution")




}


