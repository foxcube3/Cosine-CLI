package cosineapi

import (
	"os"
)

// openFileAppend opens a file for appending, creating it with 0600 permissions if it does not exist.
func openFileAppend(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
}