package interactions

import "github.com/AlecAivazis/survey/v2"

func Prompt(msg string) string {

	res := ""
	prompt := &survey.Input{
		Message: msg,
	}

	survey.AskOne(prompt, &res)
	return res
}
