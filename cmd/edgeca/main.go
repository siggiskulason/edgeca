package main

import (
	"github.com/edgeca-org/edgeca/internal/cli/cmd"
	"github.com/edgeca-org/edgeca/internal/cli/config"
)

func main() {
	config.InitCLIConfiguration()
	cmd.Execute()
}
