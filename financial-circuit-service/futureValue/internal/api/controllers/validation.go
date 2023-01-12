package controllers

import (
	"futurevalue/internal/api/json"
	"strings"
)

// Validation
func validateProveInput(p string) string {
	m := "Missing input(s): "

	if strings.Contains(p, "Data.FutureValue") {
		m += "FutureValue, "
	}
	if strings.Contains(p, "Data.PresentValue") {
		m += "PresentValue, "
	}
	if strings.Contains(p, "Data.IntrestRate") {
		m += "IntrestRate, "
	}
	if strings.Contains(p, "Data.NumberOfYears") {
		m += "NumberOfYears, "
	}
	if m == "Missing input(s): " {
		return "true"
	}
	return m + "please provide these inputs and try again"
}

func validateVerifyInput(v string) string {
	m := "Missing input(s): "
	if strings.Contains(v, "IntrestRate") {
		m += "IntrestRate, "
	}
	if strings.Contains(v, "NumberOfYears") {
		m += "NumberOfYears, "
	}
	if strings.Contains(v, "Proof") {
		m += "Proof, "
	}
	if strings.Contains(v, "VerifyingKey") {
		m += "VerifyingKey, "
	}
	if m == "Missing input(s): " {
		return "true"
	}
	return m + "please provide these inputs and try again"
}

func validateYears(v json.Data) string {
	if v.NumberOfYears > 18 {
		return "false"
	} else {
		return "true"
	}
}
