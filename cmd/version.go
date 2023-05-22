package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/zawpavel/qrcp/version"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version number and build information.",
	Run: func(c *cobra.Command, args []string) {
		fmt.Println(version.String())
	},
}
