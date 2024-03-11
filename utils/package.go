//go:build unix

package utils

import "regexp"

var conda = "conda"

var condaRegex = regexp.MustCompile(`^conda$`)
