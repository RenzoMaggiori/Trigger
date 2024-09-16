package parser_test

import (
	"flag"
	"os"
	"testing"

	"trigger.com/api/src/parser"
)

func TestParseArgs(t *testing.T) {
	port := int64(8080)
	envPath := "../../../.env"
	defaultPort := int64(8000)
	tests := []struct {
		name          string
		args          []string
		expected      *parser.Arguments
		expectedError error
	}{
		{
			name: "Valid arguments",
			args: []string{
				"-env-path=../../../.env",
				"-port=8080",
			},
			expected: &parser.Arguments{
				EnvPath: &envPath,
				Port:    &port,
			},
			expectedError: nil,
		},
		{
			name:          "Missing env-path argument",
			args:          []string{"-port=8080"},
			expected:      nil,
			expectedError: flag.ErrHelp,
		},
		{
			name: "Default port value",
			args: []string{"-env-path=../../../.env"},
			expected: &parser.Arguments{
				EnvPath: &envPath,
				Port:    &defaultPort,
			},
			expectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			os.Args = append([]string{os.Args[0]}, tt.args...)

			params, err := parser.CmdArgs()

			if err != tt.expectedError {
				t.Errorf("Expected error: %v, got: %v", tt.expectedError, err)
			}
			if params == nil && tt.expected == nil {
				return
			}
			if params == tt.expected {
				t.Errorf("Expected params: %+v, got: %+v", tt.expected, params)
			}
		})
	}
}
