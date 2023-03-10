package assignment

import (
	"futurevalue/internal/api/json"
	futureValue "futurevalue/internal/futureValue"
)

func AssignFullFutureValueCircuit(d json.Data) *futureValue.Circuit {
	assignment := &futureValue.Circuit{
		FutureValue:   d.FutureValue,
		PresentValue:  d.PresentValue,
		NumberOfYears: d.NumberOfYears,
		InterestRate:   int64(d.InterestRate * 1000),
	}
	return assignment
}

func AssignPublicFutureValueCircuit(d json.PublicData) *futureValue.Circuit {
	assignment := &futureValue.Circuit{
		NumberOfYears: d.NumberOfYears,
		InterestRate:   int64(d.InterestRate * 1000),
	}
	return assignment
}
