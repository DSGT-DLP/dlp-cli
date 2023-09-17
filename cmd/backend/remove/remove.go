/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package remove

import (
	"strconv"

	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/backend"
	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/pkg"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

// RemoveCmd represents the backend remove command
var RemoveCmd = &cobra.Command{
	Use:   "remove {package}",
	Short: "Remove package from conda environment, defaults to removal via poetry unless otherwise specified",
	Long:  `Remove package from conda environment from /training, defaults to removal via poetry unless otherwise specified`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		mamba, _ := strconv.ParseBool(cmd.Flag("mamba").Value.String())
		channel := cmd.Flag("channel").Value.String()
		env_name := cmd.Flag("env-name").Value.String()
		dev, _ := strconv.ParseBool(cmd.Flag("dev").Value.String())
		if mamba {
			pkg.ExecBashCmd(backend.BackendDir, "mamba", "run", "--live-stream", "-n", env_name, "mamba", "remove", "-c", channel, args[0])
			pterm.DefaultSection.Println("IMPORTANT")
			pterm.Info.Println("Remove the following line in dependencies section in environment.yml:\n" + "  - " + channel + "::" + args[0])
		} else if dev {
			pkg.ExecBashCmd(backend.BackendDir, "mamba", "run", "--live-stream", "-n", env_name, "poetry", "remove", args[0], "--group", "dev")
		} else {
			pkg.ExecBashCmd(backend.BackendDir, "mamba", "run", "--live-stream", "-n", env_name, "poetry", "remove", args[0])
		}
	},
}

func init() {
	backend.BackendCmd.AddCommand(RemoveCmd)
	RemoveCmd.Flags().BoolP("dev", "d", false, "Remove package from dev dependencies")
	RemoveCmd.Flags().BoolP("mamba", "m", false, "Remove package via conda (used mainly if cross-platform compatibility is needed)")
	RemoveCmd.Flags().StringP("channel", "c", "conda-forge", "Specify conda channel to install from (only used if --mamba flag is set)")
	RemoveCmd.MarkFlagsMutuallyExclusive("dev", "mamba")
}
