package arguments

import (
	"flag"
	"fmt"
	"os"
)

type CmdArgs struct {
	Port    *int64
	EnvPath *string
}

var (
	usageError = `Usage: [options]
Options:
  -port     : (Required) http server port number, must be an integer.
  -env-path : (Required) path to the .env configuration file.

Example:
  %s -port 8080 -env-path /path/to/.env
`
)

func Command() (*CmdArgs, error) {
	args := CmdArgs{}
	args.Port = flag.Int64("port", -1, "user service port")
	args.EnvPath = flag.String("env-path", "", "user service path to .env")
	flag.Parse()

	if *args.EnvPath == "" || *args.Port == -1 {
		return nil, fmt.Errorf(usageError, os.Args[0])
	}
	return &args, nil
}
