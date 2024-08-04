package utils

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type DashboardType map[string][]string

/*
* Key : Values[]
 */
func PrintDashboard(dashboardAttr DashboardType) {

	t := tablewriter.NewWriter(os.Stdout)

	headers := []string{}

	// set header
	for k := range dashboardAttr {
		headers = append(headers, k)
	}

	t.SetHeader(headers)

	// set value
	headerLen := 0
	for headerLen < len(headers) {

		values := []string{}
		for _, v := range dashboardAttr {

			values = append(values, v[headerLen])
		}

		fmt.Println(values)

		t.Append(values)
		headerLen++
	}

	t.Render()
}
