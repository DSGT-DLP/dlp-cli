/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package frontend

import (
	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd"
	"github.com/spf13/cobra"
)

const FrontendDir string = "./frontend"
const AwsRegion = "us-east-1"

// FrontendCmd represents the frontend command
var FrontendCmd = &cobra.Command{
	Use:   "frontend",
	Short: "All frontend related subcommands",
	Long:  `Contains all frontend /frontend directory related subcommands`,
	Args:  cobra.ExactArgs(0),
}

func init() {
	cmd.RootCmd.AddCommand(FrontendCmd)
}

func ExecBashCmd(name string, args ...string) string {
	return cmd.ExecBashCmd(FrontendDir, name, args...)
}
