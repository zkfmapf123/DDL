package utils

import (
	"fmt"
	"os"
	"reflect"

	"github.com/olekukonko/tablewriter"
)

const (
	ZERO_INDEX = 0
)

type DashboardType map[string][]string

func PrintDashboardUseTypeArr[T any](arr []T) {

	to := reflect.TypeOf(arr[ZERO_INDEX])

	headers := []string{}
	for i := 0; i < to.NumField(); i++ {
		headers = append(headers, to.Field(i).Name)
	}

	t := tablewriter.NewWriter(os.Stdout)
	t.SetHeader(headers)

}

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
