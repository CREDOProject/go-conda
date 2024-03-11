//go:build unix

package goconda

import "regexp"

var condaRegex = regexp.MustCompile(`^conda$`)
