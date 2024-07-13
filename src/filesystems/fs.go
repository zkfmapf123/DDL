package filesystems

import "os"

func MustGetCurrentPath() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return dir
}

func MustGetHomePath() string {
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}

	return dir
}
