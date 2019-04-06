package cost

import "math"

//Cost interface helps to standarize the way we are going to implement
//cost functions used in nnetwork
type Cost interface {
	Run(FeedForward func(input []float64) []float64, TrainingDataSet [][]float64, DesiredOutput []float64) float64
	Derivative()
}

//MSE is short for mean squared error
type MSE struct{}

//Run ...
func (c MSE) Run(FeedForward func(input []float64) []float64, TrainingDataSet [][]float64, DesiredOutput []float64) []float64 {

	var AverageOutput = make([]float64, len(DesiredOutput))

	for index := 0; index < len(TrainingDataSet); index++ {
		aux := FeedForward(TrainingDataSet[index])

		for index, val := range aux {

			AverageOutput[index] += (math.Pow((DesiredOutput[index]-val), 2) / float64(2*(len(TrainingDataSet))))
		}
	}
	return AverageOutput
}

//Derivative ...
func (c MSE) Derivative() {

}
