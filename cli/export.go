package cli

import (
	"github.com/urfave/cli/v2"
)

// ExportCommands 导出给其他模块用
func ExportCommands() cli.Commands {
	return commands
}
