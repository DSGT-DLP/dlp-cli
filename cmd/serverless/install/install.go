/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package install

import (
	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/serverless"
	"github.com/spf13/cobra"
)

// InstallCmd represents the serverless install command
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs packages from package.json",
	Long:  `Installs serverless packages from package.json from /serverless in node_modules`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		bash_args := []string{"install"}
		if cmd.Flag("force").Value.String() == "true" {
			bash_args = append(bash_args, "--force")
		}
		if cmd.Flag("yarn").Value.String() == "true" {
			serverless.ExecBashCmd("yarn", bash_args...)
		} else {
			serverless.ExecBashCmd("pnpm", bash_args...)
		}
	},
}

func init() {
	serverless.ServerlessCmd.AddCommand(InstallCmd)
	InstallCmd.Flags().BoolP("force", "f", false, "Force a reinstall of serverless packages")
}
