package readers

import (
	"os"
	"path"

	"github.com/rodrigoaraujo46/assert"
)

func ReadExample(todaysPath string) string {
	const filePath = "example.txt"
	fileBytes, err := os.ReadFile(path.Join(todaysPath, filePath))
	assert.NoError(err, "Failed reading example file")

	return string(fileBytes)
}

func ReadInput(todaysPath string) string {
	const filePath = "input.txt"
	fileBytes, err := os.ReadFile(path.Join(todaysPath, filePath))
	assert.NoError(err, "Failed reading example file")

	return string(fileBytes)
}
