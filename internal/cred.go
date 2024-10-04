package internal

import (
	"bufio"
	"errors"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/zkfmapf123/ddl/utils"
	"gopkg.in/yaml.v3"
)

type DDLCreds struct {
	Profile string `yaml:"profile"`
	Region  string `yaml:"region"`

	// ...
	AccessKey string `yaml:"accessKey"`
	SecretKey string `yaml:"secretKey"`
}

func GetAWSCredsProfile() (DDLCreds, error) {
	path, err := getAWSCredentialPath(utils.AWS_DDL_PATH)

	// not exists
	if err != nil {
		return DDLCreds{}, errors.New("not Exists File")
	}

	// already exists
	return getDDLCreds(path)
}

func getDDLCreds(path string) (DDLCreds, error) {
	file, err := os.Open(path)
	if err != nil {
		return DDLCreds{}, err
	}
	defer file.Close()

	fb, err := io.ReadAll(file)
	if err != nil {
		return DDLCreds{}, err
	}
	defer file.Close()

	var ddlCreds DDLCreds
	if err := yaml.Unmarshal(fb, &ddlCreds); err != nil {
		return DDLCreds{}, err
	}

	return ddlCreds, nil
}

func SetDDLCreds(path string, profile string, region string) error {

	filePath, _ := getAWSCredentialPath(path)

	creds := DDLCreds{
		Profile: profile,
		Region:  region,
	}

	yamlBytes, err := yaml.Marshal(creds)
	if err != nil {
		return err
	}

	err = os.WriteFile(filePath, yamlBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func GetAWSCredsBeforeSaveDDLCreds(path string) ([]string, error) {
	filePath, _ := getAWSCredentialPath(path)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var profiles []string
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		line := sc.Text()

		if strings.HasPrefix(line, "[") && strings.HasSuffix(line, "]") {
			profiles = append(profiles, strings.Trim(line, "[]"))
		}
	}

	return profiles, nil
}

func isExistsFile(path string) error {
	_, err := os.Stat(path)
	return err
}

func getAWSCredentialPath(path string) (string, error) {

	dir, _ := os.UserHomeDir()
	err := isExistsFile(dir)

	return filepath.Join(dir, path), err
}
