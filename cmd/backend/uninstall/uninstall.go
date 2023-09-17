/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package install

import (
	"fmt"

	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/backend"
	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/pkg"
	"github.com/spf13/cobra"
)

// InstallCmd represents the backend install command
var InstallCmd = &cobra.Command{
	Use:   "uninstall",
	Short: "Uninstalls all training backend packages and removes conda environment for dlp",
	Long:  `Uninstalls all training backend packages and removes conda environment for dlp`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		env_name := cmd.Flag("env-name").Value.String()
		if cmd.Flag("reference").Value.String() == "true" {
			fmt.Println("Remove " + env_name + " conda environment:")
			fmt.Println("\tconda remove -n " + env_name + " -y --all")
			return
		}
		pkg.ExecBashCmd(backend.BackendDir, "conda", "remove", "-n", env_name, "-y", "--all")
	},
}

func init() {
	backend.BackendCmd.AddCommand(InstallCmd)
}
