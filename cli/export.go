package cli

import (
	"github.com/urfave/cli"
)

// ExportCommands 导出给其他模块用
func ExportCommands() []cli.Command {
	return commands
}
