/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package add

import (
	"strconv"

	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/backend"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// AddCmd represents the backend add command
var AddCmd = &cobra.Command{
	Use:   "add {package}",
	Short: "Add package to conda environment, defaults to installation via poetry unless otherwise specified",
	Long:  `Add package to conda environment from /training, defaults to installation via poetry unless otherwise specified`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mamba, _ := strconv.ParseBool(cmd.Flag("mamba").Value.String())
		channel := cmd.Flag("channel").Value.String()
		env_name := cmd.Flag("env-name").Value.String()
		dev, _ := strconv.ParseBool(cmd.Flag("dev").Value.String())
		if mamba {
			backend.ExecBashCmd("mamba", "run", "--live-stream", "-n", env_name, "mamba", "install", "-c", channel, args[0])
			pterm.DefaultSection.Println("IMPORTANT")
			pterm.Info.Println("Add the following line in dependencies section in environment.yml:\n" + "  - " + channel + "::" + args[0])
			pterm.Info.Println("Add the following line at the bottom of the channels section in environment.yml above defaults:\n" + "  - " + channel)
			pterm.Info.Println("Anaconda docs also recommend reinstalling the conda environment to reduce conflicts between conda-forge and PyPI dependencies, so after adding the above line, run:\ndlp-cli backend install --force")
		} else if dev {
			backend.ExecBashCmd("mamba", "run", "--live-stream", "-n", env_name, "poetry", "add", args[0], "--group", "dev")
		} else {
			backend.ExecBashCmd("mamba", "run", "--live-stream", "-n", env_name, "poetry", "add", args[0])
		}
	},
}

func init() {
	backend.BackendCmd.AddCommand(AddCmd)
	AddCmd.Flags().BoolP("dev", "d", false, "Add package as dev dependency")
	AddCmd.Flags().BoolP("mamba", "m", false, "Add package via mamba (used mainly if cross-platform compatibility is needed)")
	AddCmd.Flags().StringP("channel", "c", "conda-forge", "Specify conda channel to install from (only used if --mamba flag is set)")
	AddCmd.MarkFlagsMutuallyExclusive("dev", "mamba")
}
