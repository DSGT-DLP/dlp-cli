/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package install

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/backend"
	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/pkg"
	"github.com/spf13/cobra"
)

// InstallCmd represents the backend install command
var InstallCmd = &cobra.Command{
	Use:   "install",
	Short: "Installs training backend packages from pyproject.toml",
	Long:  `Installs training backend packages from pyproject.toml from /training in .venv`,
	Args:  cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		force, _ := strconv.ParseBool(cmd.Flag("force").Value.String())
		name := cmd.Flag("name").Value.String()
		py_version := cmd.Flag("python-version").Value.String()
		if cmd.Flag("reference").Value.String() == "true" {
			fmt.Println("Check if " + name + " conda environment is created:")
			fmt.Println("\tconda info --envs")
			fmt.Println("If not created, create conda environment:")
			fmt.Println("\tconda create --name " + name + " python=3.9")
			fmt.Println("Activate conda env if not already activated:")
			fmt.Println("\tconda activate dlp")
			fmt.Println("If poetry not installed in conda env, install poetry:")
			fmt.Println("\tconda install -c conda-forge poetry")
			fmt.Println("Install python packages from pyproject.toml with poetry:")
			fmt.Println("\tpoetry install")
			return
		}
		res := pkg.ExecBashCmd(backend.BackendDir, "conda", "info", "--envs")
		activated := strings.Contains(strings.ReplaceAll(res, " ", ""), name+"*")
		if strings.Contains(res, name) && force {
			pkg.ExecBashCmd(backend.BackendDir, "conda", "remove", "-n", name, "-y", "--all")
		}
		if !strings.Contains(res, name) || force {
			pkg.ExecBashCmd(backend.BackendDir, "conda", "create", "--name", name, "-y", "python="+py_version)
		}
		cmd.Println(activated)
		pkg.ExecBashCmd(backend.BackendDir, "conda", "run", "-n", name, "conda", "install", "-c", "conda-forge", "poetry")
		pkg.ExecBashCmd(backend.BackendDir, "conda", "run", "-n", name, "poetry", "install")
		//pkg.ExecBashCmd(backend.BackendDir, "bash", "-c", "eval \"$(conda shell.bash hook)\" && conda activate "+name+" && poetry install")
		/*
			pkg.ExecBashCmd(backend.BackendDir, "conda", "create", "--name", name, "--file", "requirements.txt")
			pkg.ExecBashCmd(backend.BackendDir, "pyenv", "local", "3.9")
			pkg.ExecBashCmd(backend.BackendDir, "poetry", "install")*/
	},
}

func init() {
	backend.BackendCmd.AddCommand(InstallCmd)
	InstallCmd.Flags().BoolP("force", "f", false, "Force a reinstall of backend packages")
	InstallCmd.Flags().String("name", "dlp", "Name of the conda environment you want to create")
	InstallCmd.Flags().String("python-version", "3.9", "Python version to specify when creating the conda environment")
}
