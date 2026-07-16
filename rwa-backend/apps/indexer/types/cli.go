package types

import (
	"github.com/cb00j/cbj-rwa/rwa-backend/libs/core/bootstrap"
	"github.com/urfave/cli/v3"
)

const (
	ConfigFile        = bootstrap.FlagConfigFile
	ConfigFileDefault = bootstrap.FlagConfigFileDefault
)

type Cli struct {
	ConfigFilePath string
}

func NewCli(cmd *cli.Command) *Cli {
	return &Cli{
		ConfigFilePath: cmd.String(ConfigFile),
	}
}
