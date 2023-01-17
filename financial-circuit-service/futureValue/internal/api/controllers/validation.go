package controllers

import (
	"futurevalue/internal/api/json"
)

func validateYears(v json.Data) string {
	if v.NumberOfYears > 18 {
		return "false"
	} else {
		return "true"
	}
}
