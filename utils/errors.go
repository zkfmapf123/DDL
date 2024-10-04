package utils

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func PanicError(err error) {
	fmt.Println(color.RedString("[success] %s", err.Error()))
	os.Exit(0)
}

func PanicWarn(err error) {
	fmt.Println(color.YellowString("[success] %s", err.Error()))
	os.Exit(0)

}

func PanicDebug(err error) {
	fmt.Println(color.GreenString("[success] %s", err.Error()))
	os.Exit(0)
}

func MustCheckError(err error) {
	if err != nil {
		PanicError(err)
	}
}
