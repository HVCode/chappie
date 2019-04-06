package nnetwork

import (
	"github.com/HVCode/chappie/activation"
	"github.com/HVCode/chappie/cost"
	"github.com/HVCode/chappie/neuron"
)

//NeuralNetwork ...
type NeuralNetwork struct {
	InputLayer   Layer
	HideLayers   []Layer
	OutputLayer  Layer
	LearningRate float64
	trainMethod  func([]float64, func([]float64) float64) float64
	costFunction cost.Cost
}

//StartFeedForward init the think process.
//Remember to pass an input with the same size of InputLayer
func (n *NeuralNetwork) StartFeedForward(input []float64) []float64 {

	//if len(input) == len(n.InputLayer.Neurons)
	input = n.InputLayer.SendSignal(input)
	for _, element := range n.HideLayers {
		input = element.SendSignal(input)
	}
	input = n.OutputLayer.SendSignal(input)
	return input
}

//SetInputAndOutputLayers ...
func (n *NeuralNetwork) SetInputAndOutputLayers(NeuronsInInputLayer, NeuronsInOutputLayer int,
	inputActivationFunc, outputActivationFunc activation.Activation) {

	n.InputLayer.SetNeurons(NeuronsInInputLayer, neuron.InputLayer)
	n.InputLayer.ActivationFunction = inputActivationFunc

	n.OutputLayer.SetNeurons(NeuronsInOutputLayer, neuron.OutputLayer)
	n.OutputLayer.ActivationFunction = outputActivationFunc

}

//SetAmountOfHideLayers ...
func (n *NeuralNetwork) SetAmountOfHideLayers(LayersInHideLayers int) {

	n.HideLayers = make([]Layer, LayersInHideLayers)

}

//SetHideLayers ...
func (n *NeuralNetwork) SetHideLayers(index, NeuronsInHideLayer int, hideActivationFunc activation.Activation) {

	n.HideLayers[index].SetNeurons(NeuronsInHideLayer, neuron.HideLayer)
	n.HideLayers[index].ActivationFunction = hideActivationFunc
}

//SetCostFunction ...
func (n *NeuralNetwork) SetCostFunction(c cost.Cost) {
	n.costFunction = c
}

//SetTrainMethod ...
func (n *NeuralNetwork) SetTrainMethod(t func([]float64, func([]float64) float64) float64) {
	n.trainMethod = t
}

//Train //remember DesiredOutput must has the same size that OutputLayer
func (n *NeuralNetwork) Train(TrainingDataSet [][]float64, DesiredOutput []float64, Epochs int, LearningRate float32) {

	for index := 0; index < Epochs; index++ {
		//AverageOutput := n.costFunction.Run(n.StartFeedForward, TrainingDataSet, DesiredOutput)

	}
}
