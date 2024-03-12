package utils

import (
	"errors"

	"github.com/CREDOProject/sharedutils/files"
	"github.com/CREDOProject/sharedutils/shell"
)

var execCommander = shell.New

// Function is used to detect the path to the Conda binary on the current
// system. It returns a string representing the path to the Conda
// binary file, or an error if it fails to detect the binary.
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
