package parser

import (
	"flag"
	"fmt"
)

type Arguments struct {
	EnvPath *string
}

const Usage string = `
Usage: -env-path <path_to_env_file>
`

func CmdArgs() (*Arguments, error) {
	args := Arguments{}
	args.EnvPath = flag.String("env-path", "", "Path to .env")
	flag.Parse()

	if *args.EnvPath == "" {
		return nil, fmt.Errorf("missing required argument: -env-path is required%s", Usage)
	}
	return &args, nil
}
