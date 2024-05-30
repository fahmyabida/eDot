package config_test

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/fahmyabida/eDot/cmd/config"

	"github.com/stretchr/testify/assert"
)

func getAbsolutePathName(filename string) string {
	_, thisFilePath, _, _ := runtime.Caller(0)
	dotEnvPath := filepath.Join(filepath.Dir(thisFilePath), "./../..", filename)
	return dotEnvPath
}

func writeToTempFile(dotEnvPath string, inputStrings []string) error {
	f, err := os.Create(dotEnvPath)
	if err != nil {
		return err
	}

	for _, s := range inputStrings {
		_, err = f.WriteString(s + "\n")
	}

	return err
}

func removeTempFile(dotEnvPath string) error {
	return os.Remove(dotEnvPath)
}

func TestInitEnv(t *testing.T) {
	filename := ".temp.env"
	dotEnvPath := getAbsolutePathName(filename)

	err := writeToTempFile(dotEnvPath, []string{
		"MOCK_KEY=MOCK_VALUE",
	})
	assert.NoError(
		t,
		err,
		"Failed in writing to file",
	)

	defer func() {
		err = removeTempFile(dotEnvPath)
		assert.NoError(
			t,
			err,
			"Failed in writing to file",
		)
	}()

	err = config.InitEnv(filename)
	assert.NoError(t, err)
	assert.Equal(
		t,
		os.Getenv("MOCK_KEY"),
		"MOCK_VALUE",
		"should successfully load env variables",
	)
}
