package utils

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type TwParams struct {
	header []string
	body   [][]string
}

func NewTableWriter() *TwParams {
	return &TwParams{}
}

func (t *TwParams) Setheader(header []string) *TwParams {
	t.header = header
	return t
}

func (t *TwParams) SetBody(body [][]string) *TwParams {
	t.body = body
	return t
}

func (t *TwParams) Print() {
	h, b := t.header, t.body

	tt := tablewriter.NewWriter(os.Stdout)
	tt.SetHeader(h)

	for _, v := range b {
		tt.Append(v)
	}

	tt.Render()
}
