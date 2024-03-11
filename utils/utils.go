package utils

import (
	"errors"

	"github.com/CREDOProject/sharedutils/files"
	"github.com/CREDOProject/sharedutils/shell"
)

var execCommander = shell.New

// Function used to find the pip binary in the system.
func DetectCondaBinary() (string, error) {
	return execCommander().LookPath(conda)
}

func CondaBinaryFrom(path string) (string, error) {
	execs, err := files.ExecsInPath(path, looksLikeConda)
	if err != nil {
		return "", err
	}
	if len(execs) < 1 {
		return "", errors.New("No conda found.")
	}

	return execs[0], err

}

// looksLikeConda returns true if the given filename looks like a Python
// executable.
func looksLikeConda(name string) bool {
	return condaRegex.MatchString(name)
}
