package utils

import (
	"errors"
	"os/exec"
	"testing"

	"github.com/CREDOProject/sharedutils/shell"
)

type execShim func() shell.ExecShim

type myLookPath struct {
	LookPathFunc func(string) (string, error)
}

func (m myLookPath) LookPath(name string) (string, error) {
	return m.LookPathFunc(name)
}

func (m myLookPath) Command(cmd string, args ...string) *exec.Cmd {
	return nil
}

func mockExec() execShim {
	return func() shell.ExecShim {
		shim := myLookPath{
			LookPathFunc: func(name string) (string, error) {
				return name, nil
			},
		}
		return shim
	}
}

func Test_detectPipBinary(t *testing.T) {
	prevCommander := execCommander
	defer func() { execCommander = prevCommander }()

	execCommander = mockExec()

	result, error := DetectCondaBinary()

	if error != nil {
		t.Fatalf("Logic error.")
	}

	if result != conda {
		t.Fatalf("Unexpected result. Got %s, wants %s", result, conda)
	}
}

func TestLooksLikePip(t *testing.T) {
	tests := []struct {
		filename string
		expected bool
	}{
		{"conda", true},
		{"condaexe", false},
		{"12conda", false},
		{"pypip3.10", false}, // should start with "conda"
	}

	for _, test := range tests {
		t.Run(test.filename, func(t *testing.T) {
			result := looksLikeConda(test.filename)
			if result != test.expected {
				t.Errorf("looksLikeConda(%s) returned %t, want %t", test.filename, result, test.expected)
			}
		})
	}
}

func TestPipBinaryFrom(t *testing.T) {
	tests := []struct {
		path        string
		mockResult  []string
		mockErr     error
		expected    string
		expectedErr error
	}{
		{
			path:        "/path/to/some/directory",
			mockResult:  nil,
			mockErr:     nil,
			expected:    "",
			expectedErr: errors.New("No conda found."),
		},
		{
			path:        "/path/to/some/directory",
			mockResult:  nil,
			mockErr:     errors.New("Some error occurred"),
			expected:    "",
			expectedErr: errors.New("Some error occurred"),
		},
	}

	for _, test := range tests {
		t.Run(test.path, func(t *testing.T) {
			// Call PipBinaryFrom with the mock
			result, _ := CondaBinaryFrom(test.path)

			// Check if the result matches the expectation
			if result != test.expected {
				t.Errorf("CondaBinaryFrom(%s) returned %s, want %s", test.path, result, test.expected)
			}

		})
	}
}
