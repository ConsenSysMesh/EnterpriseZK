package assignment

import (
	"futurevalue/internal/api/json"
	futureValue "futurevalue/internal/futureValue"
)

func AssignFullLocationCircuit(d json.Data) *futureValue.Circuit {
	assignment := &futureValue.Circuit{
		FutureValue:   d.FutureValue,
		PresentValue:  d.PresentValue,
		NumberOfYears: d.NumberOfYears,
		IntrestRate:   int64(d.IntrestRate * 1000),
	}
	return assignment
}

func AssignPublicLocationCircuit(d json.PublicData) *futureValue.Circuit {
	assignment := &futureValue.Circuit{
		NumberOfYears: d.NumberOfYears,
		IntrestRate:   int64(d.IntrestRate * 1000),
	}
	return assignment
}
