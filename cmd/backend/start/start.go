/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package start

import (
	"fmt"

	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/backend"
	"github.com/spf13/cobra"
)

// StartCmd represents the backend start command
var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the training backend",
	Long:  `Starts an instance of the training backend Django app in /training in the terminal`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		env_name := cmd.Flag("env-name").Value.String()
		backend.ExecBashCmd("mamba", "run", "--live-stream", "-n", env_name, "poetry", "run", "python", "manage.py", "runserver", fmt.Sprintf("%v", cmd.Flag("port").Value))
	},
}

func init() {
	backend.BackendCmd.AddCommand(StartCmd)
	StartCmd.PersistentFlags().IntP("port", "p", 8000, "A port to run the backend on")
	StartCmd.PersistentFlags().StringP("env-name", "e", "dlp", "The name of the mamba environment")
}
