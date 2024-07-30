package aws

import (
	"fmt"
	"strings"

	"github.com/spf13/viper"
	"github.com/zkfmapf123/DDL/src/utils"
)

func (aw *AWSconfig) MustBeforeHook() (ecrParams ECRParams) {
	fmt.Println("Profile :", viper.GetString("profile"))
	fmt.Println("Region :", viper.GetString("region"))

	ecrRepo := mustIspectNeedsToParameters(*aw)

	return ecrRepo
}

// SSM
// ECR
// NotExists -> Create
func mustIspectNeedsToParameters(aw AWSconfig) (ecrParams ECRParams) {
	ssmValue, ecrName, err := aw.RetrieveSSMValue()

	for {

		if err == nil {
			break
		}

		errMsg := err.Error()

		if strings.Contains(errMsg, utils.AWS_ERROR_PARAMETER_NOT_FOUND) {
			ecrName, err = aw.SetSSMValue()
			if err != nil {
				panic(err)
			}

			break
		}

		if strings.Contains(errMsg, utils.AWS_ERROR_PARAMETER_ALREADY_EXISTS) {
			fmt.Println("Already Exists (SSM)", ssmValue)
			break
		}

		panic(err)
	}

	ecrParam, _ := aw.RetrievECR(ecrName)

	if ecrParam.Name == "" {
		ecrParam, err = aw.CreateECR(ecrName)

		for {
			if err == nil {
				break
			}

			errMsg := err.Error()

			if strings.Contains(errMsg, utils.AWS_ERROR_ECR_ALREADY_EXISTS) {
				fmt.Println("Already Exists (ECR)", ecrName)
				break
			}

			panic(err)
		}

	}

	return ecrParam
}
