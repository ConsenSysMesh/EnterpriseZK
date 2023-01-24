package futureValue

import (
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/std/math/bits"
)

// fv*10^(n*3)

// x := 9000000.515174083496093750
// y := 9000000515174083496093750

type Circuit struct {
	//PRIVATE INPUTS
	PresentValue frontend.Variable
	FutureValue  frontend.Variable

	//PUBLIC INPUTS
	InterestRate   frontend.Variable `gnark:",public"`
	NumberOfYears frontend.Variable `gnark:",public"`
}
/*
	PV formula, rewritten to eliminate division and avoid floating decimals
	PV = FV[1/(1+i)^N]
	FV = PV[(1+i)^N]
	FV = PV[(1000 + 1000(i))^N]	
	i – interest rate (comes in as i*1000)
	N – number of years
*/

func (circuit *Circuit) Define(api frontend.API) error {

	i := circuit.InterestRate //already x1000

	//Adds 1000 to interest
	iplus := api.Add(1000, i)

	//EXPONENT FOR (1000+I)^N--------
	const bitSize = 8
	output := frontend.Variable(1)
	bits1 := bits.ToBinary(api, circuit.NumberOfYears, bits.WithNbDigits(bitSize))

	for i := 0; i < len(bits1); i++ {
		if i != 0 {
			output = api.Mul(output, output)
		}
		multiply := api.Mul(output, iplus)
		output = api.Select(bits1[len(bits1)-1-i], multiply, output)
	}
	//-------

	fv := api.Mul(output, circuit.PresentValue)

	expo := api.Mul(3, circuit.NumberOfYears)

	/* 
		EXPONENT FOR 10^(3*N)
		This exponent loop multiplies takes expo value (3N) and multiplies 10 to that power
		This is the calculated formula to determine the length of the produced value,
		which is needed in order to properly align the end result with the comparison value
	*/
	
	modif := frontend.Variable(1)
	bits2 := bits.ToBinary(api, expo, bits.WithNbDigits(bitSize))


	for i := 0; i < len(bits2); i++ {
		if i != 0 {
			modif = api.Mul(modif, modif)
		}
		multiply := api.Mul(modif, 10)
		modif = api.Select(bits2[len(bits2)-1-i], multiply, modif)
	}
	

	fv3 := api.Mul(modif, circuit.FutureValue)

	api.AssertIsEqual(api.Cmp(api.Sub(fv3, modif), fv), -1)
	api.AssertIsEqual(api.Cmp(api.Add(fv3, modif), fv), 1)

	return nil
}
