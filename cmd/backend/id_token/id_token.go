package id_token

/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/

import (
	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/cmd/backend"
	"github.com/spf13/cobra"
)

// IdTokenCmd represents the IdToken command
var IdTokenCmd = &cobra.Command{
	Use:   "id-token {email}",
	Short: "gets a user's id token by email",
	Long:  `gets a user's id token by email from the backend`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		env_name := cmd.Flag("env-name").Value.String()
		backend.ExecBashCmd("mamba", "run", "-n", env_name, "poetry", "run", "python", "cli.py", "get-id-token", args[0])
	},
}

func init() {
	backend.BackendCmd.AddCommand(IdTokenCmd)
	IdTokenCmd.PersistentFlags().StringP("env-name", "e", "dlp", "The name of the mamba environment")
}
