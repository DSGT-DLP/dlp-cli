/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package backend

import (
	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd"
	"github.com/spf13/cobra"
)

const BackendDir string = "./training"

// BackendCmd represents the backend command
var BackendCmd = &cobra.Command{
	Use:   "backend",
	Short: "All training backend related subcommands",
	Long:  `Contains all backend /training directory related subcommands`,
	Args:  cobra.ExactArgs(0),
}

func init() {
	cmd.RootCmd.AddCommand(BackendCmd)
	BackendCmd.PersistentFlags().String("env-name", "dlp", "Name of the conda environment you want to create")
}

func ExecBashCmd(name string, args ...string) string {
	return cmd.ExecBashCmd(BackendDir, name, args...)
}
