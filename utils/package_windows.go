//go:build windows

package utils

import "regexp"

var conda = "conda.exe"

var condaRegex = regexp.MustCompile(`^conda\.exe$`)
