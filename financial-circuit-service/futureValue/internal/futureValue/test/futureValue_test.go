package futureValue

import (
	"futurevalue/internal/futureValue"
	"testing"

	"github.com/consensys/gnark-crypto/ecc"
	"github.com/consensys/gnark/backend"
	"github.com/consensys/gnark/frontend"
	"github.com/consensys/gnark/test"
)

var ExampleValues1 = &futureValue.Circuit{
	FutureValue:   6000000000,
	PresentValue:  2345548629,
	InterestRate:   110,
	NumberOfYears: 9,
}

var ExampleValues2 = &futureValue.Circuit{
	FutureValue:   9000000,
	PresentValue:  5831654,
	InterestRate:   75,
	NumberOfYears: 6,
}

var ExampleValues3 = &futureValue.Circuit{
	FutureValue:   11088900,
	PresentValue:  9000000,
	InterestRate:   110,
	NumberOfYears: 2,
}

var ExampleValues4 = &futureValue.Circuit{
	FutureValue:   309611009,
	PresentValue:  85000000,
	InterestRate:   90,
	NumberOfYears: 15,
}

var ExampleValues5 = &futureValue.Circuit{
	FutureValue:   881143,
	PresentValue:  850000,
	InterestRate:   9,
	NumberOfYears: 40,
}

func TestDCFCircuit(t *testing.T) {
	assert := test.NewAssert(t)
	var DCFCircuit futureValue.Circuit

	assert.ProverSucceeded(&DCFCircuit, ExampleValues1, test.WithCurves(ecc.BN254), test.WithBackends(backend.GROTH16), test.WithCompileOpts(frontend.IgnoreUnconstrainedInputs()))
	assert.ProverSucceeded(&DCFCircuit, ExampleValues2, test.WithCurves(ecc.BN254), test.WithBackends(backend.GROTH16), test.WithCompileOpts(frontend.IgnoreUnconstrainedInputs()))
	assert.ProverSucceeded(&DCFCircuit, ExampleValues3, test.WithCurves(ecc.BN254), test.WithBackends(backend.GROTH16), test.WithCompileOpts(frontend.IgnoreUnconstrainedInputs()))
	assert.ProverSucceeded(&DCFCircuit, ExampleValues4, test.WithCurves(ecc.BN254), test.WithBackends(backend.GROTH16), test.WithCompileOpts(frontend.IgnoreUnconstrainedInputs()))
	assert.ProverFailed(&DCFCircuit, ExampleValues5, test.WithCurves(ecc.BN254), test.WithBackends(backend.GROTH16), test.WithCompileOpts(frontend.IgnoreUnconstrainedInputs()))

}
