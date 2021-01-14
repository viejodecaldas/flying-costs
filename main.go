package main

import (
	"fmt"
	"github.com/viejodecaldas/nuorder/costs"
)

func main() {
	// This is the sample inputs provided (can be changed)
	input := []costs.FlyingCost{{10,20},{30,200},{400,50},{30,20}}

	finalResult := costs.GetMinimumCost(input)

	// print out the total
	fmt.Println("TOTAL MINIMUM COST IS: ", finalResult)
}

