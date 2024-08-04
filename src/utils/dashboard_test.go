package utils

import "testing"

func Test_dashboa(t *testing.T) {

	friends := map[string][]string{
		"Name": {"kayla", "joseph", "eric"},
		"Job":  {"BE", "Devopse", "FullStack"},
		"Age":  {"30", "30", "30"},
	}

	PrintDashboard(friends)
}
