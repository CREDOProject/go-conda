//go:build windows

package goconda

import "regexp"

var condaRegex = regexp.MustCompile(`^conda\.exe$`)
