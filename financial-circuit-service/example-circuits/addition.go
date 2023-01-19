package futureValue

import (
	"github.com/consensys/gnark/frontend"
)

type Circuit struct {
	//PRIVATE INPUTS
	x 	frontend.Variable
	y   frontend.Variable


	//PUBLIC INPUTS
	z   frontend.Variable	`gnark:",public"`
}


func (circuit *Circuit) Define(api frontend.API) error {

	z := api.Add(circuit.x, circuit.y)
	api.AssertIsEqual(z, circuit.z)

	return nil
}
