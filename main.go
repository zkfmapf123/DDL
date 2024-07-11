package main

import "github.com/zkfmapf123/DDL/src/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
