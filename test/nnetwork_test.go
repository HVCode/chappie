package test

import (
	"chappie/activation"
	"chappie/nnetwork"
	"testing"
)

//also called X
var input = [][]float64{
	{0, 0, 1},
	{0, 1, 1},
	{1, 0, 1},
	{1, 1, 1},
}

//also called y
var desiredOutput = [4]int8{0, 0, 1, 1}

func Test_StartFeedForward(t *testing.T) {
	test := new(nnetwork.NeuralNetwork)
	test.SetInputAndOutputLayers(3, 1, activation.HLT{}, activation.HLT{})
	test.SetAmountOfHideLayers(2)
	test.SetHideLayers(0, 1, activation.HLT{})
	test.SetHideLayers(1, 1, activation.HLT{})

	test.StartFeedForward(input[0])
}
