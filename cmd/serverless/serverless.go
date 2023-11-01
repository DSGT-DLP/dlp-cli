/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package serverless

import (
	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd"
	"github.com/spf13/cobra"
)

const ServerlessDir string = "./serverless"

// ServerlessCmd represents the serverless command
var ServerlessCmd = &cobra.Command{
	Use:   "serverless",
	Short: "All serverless related subcommands",
	Long:  `Contains all serverless /serverless directory related subcommands`,
	Args:  cobra.ExactArgs(0),
}

func init() {
	cmd.RootCmd.AddCommand(ServerlessCmd)
	cmd.RootCmd.PersistentFlags().Bool("yarn", false, "Uses yarn instead of pnpm")
}

func ExecBashCmd(name string, args ...string) string {
	return cmd.ExecBashCmd(ServerlessDir, name, args...)
}
