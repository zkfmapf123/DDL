package utils

import (
	"fmt"
	"testing"
)

func Test_dashbaord(t *testing.T) {

	friends := map[string][]string{
		"Name": {"kayla", "joseph", "eric"},
		"Job":  {"BE", "Devopse", "FullStack"},
		"Age":  {"30", "30", "30"},
	}

	PrintDashboard(friends)
}

func Test_dashboardUseArr(t *testing.T) {

	type TestParams struct {
		Name     string
		CodeSize string
	}

	lambdas := make([]TestParams, 5)

	for i := 0; i < len(lambdas); i++ {
		lambdas[i].Name = fmt.Sprintf("Name-%d", i)
		lambdas[i].CodeSize = fmt.Sprintf("CodeSize-%d", i)
	}

	PrintDashboardUseTypeArr(lambdas)
}
