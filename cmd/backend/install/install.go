/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package install

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/backend"
	"github.com/spf13/cobra"
)

// InstallCmd represents the backend install command
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs training backend packages from pyproject.toml using poetry and environment.yml using conda in a conda environment for dlp, creates conda environment if it doesn't exist",
	Long:  "Installs training backend packages from pyproject.toml using poetry and environment.yml using conda in a conda environment for dlp, creates conda environment if it doesn't exist",
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := strconv.ParseBool(cmd.Flag("force").Value.String())
		env_name := cmd.Flag("env-name").Value.String()
		if cmd.Flag("reference").Value.String() == "true" {
			fmt.Println("Check if " + env_name + " conda environment is created:")
			fmt.Println("\tmamba info --envs")
			fmt.Println("If not created, create conda environment:")
			fmt.Println("\tmamba create --name " + env_name + " python=3.9")
			fmt.Println("Activate conda env if not already activated:")
			fmt.Println("\tmamba activate dlp")
			fmt.Println("Install packages from environment.yml:")
			fmt.Println("\tmamba env update --file environment.yml --prune")
			fmt.Println("Install python packages from pyproject.toml with poetry:")
			fmt.Println("\tpoetry install")
			return
		}
		res := backend.ExecBashCmd("mamba", "info", "--envs")
		if strings.Contains(res, env_name) && force {
			backend.ExecBashCmd("mamba", "remove", "-n", env_name, "-y", "--all")
		}
		if !strings.Contains(res, env_name) || force {
			backend.ExecBashCmd("mamba", "create", "--name", env_name, "-y")
		}
		backend.ExecBashCmd("mamba", "run", "--live-stream", "-n", env_name, "mamba", "env", "update", "--file", "environment.yml", "--prune")
		backend.ExecBashCmd("mamba", "run", "--live-stream", "-n", env_name, "poetry", "install")
	},
}

func init() {
	backend.BackendCmd.AddCommand(InstallCmd)
	InstallCmd.Flags().BoolP("force", "f", false, "Force a reinstall of backend packages")
}
