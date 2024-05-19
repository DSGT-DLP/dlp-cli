/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package start

import (
	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/backend"
	"github.com/spf13/cobra"
)

// StartCmd represents the backend start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the training backend",
	Long:  `Starts an instance of the training backend Django app in /training in the terminal. To change the backend port, set the environment var BACKEND_PORT. `,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		backend.ExecBashCmd("docker", "compose", "up", "--build")
	},
}

func init() {
	backend.BackendCmd.AddCommand(StartCmd)
}
