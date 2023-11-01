/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package serverless

import (
	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/serverless"
	"github.com/spf13/cobra"
)

// StartCmd represents the serverless start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the serverless environment",
	Long:  `Starts SST's Live Lambda Development environment in the terminal`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		if cmd.Flag("yarn").Value.String() == "true" {
			serverless.ExecBashCmd("yarn", "sst", "dev")
		} else {
			serverless.ExecBashCmd("pnpm", "sst", "dev")
		}
	},
}

func init() {
	serverless.ServerlessCmd.AddCommand(StartCmd)
}
