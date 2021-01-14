package costs_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/viejodecaldas/nuorder/costs"
	"testing"
)

func TestGetMinimumCost(t *testing.T) {

	input := []costs.FlyingCost{{10,20},{30,200},{400,50},{30,20}}

	minimumCost := costs.GetMinimumCost(input)
	assert.Equal(t, 110, minimumCost, "Something went wrong")

	input = []costs.FlyingCost{{10,10},{300,200},{40,50},{30,200}}

	minimumCost = costs.GetMinimumCost(input)
	assert.Equal(t, 280, minimumCost, "Something went wrong")

	input = []costs.FlyingCost{{10,10},{30,0},{400,401},{30,29}}

	minimumCost = costs.GetMinimumCost(input)
	assert.Equal(t, 439, minimumCost, "Something went wrong")
}