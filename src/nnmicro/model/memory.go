package model

type Memory interface {
	CreatePerceptron(layers []int32)
	SaveWeights()
}
