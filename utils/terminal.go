package utils

import (
	"os"
	"os/exec"

	"github.com/AlecAivazis/survey/v2"
)

type TerminalParams struct {
}

func NewTerminal() TerminalParams {
	return TerminalParams{}
}

func (t TerminalParams) Clear() TerminalParams {
	cexec("clear")
	return t
}

func (t TerminalParams) SelectOne(title string, items []string) string {

	pp := survey.Select{
		Message: title,
		Options: items,
	}

	var pick string
	MustCheckError(survey.AskOne(&pp, &pick))
	return pick
}

func cexec(cmds ...string) {

	cmd := exec.Command(cmds[0], cmds[1:]...)
	cmd.Stdout = os.Stdout
	MustCheckError(cmd.Run())
}
