/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/DSGT-DLP/Deep-Learning-Playground/cli/pkg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "dlp-cli",
	Short: "DLP's CLI",
	Long:  `Welcome to DLP's CLI!`,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		viper.SetConfigName("cli-config") // name of config file (without extension)
		viper.SetConfigType("yaml")       // REQUIRED if the config file does not have the extension in the name
		viper.AddConfigPath(".")
		viper.ReadInConfig()
		dir, _ := os.Getwd()
		os.Chdir(filepath.Clean(filepath.Join(dir, viper.GetString("project-dir"))))
	},
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := RootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	RootCmd.PersistentFlags().String("project-dir", ".", "The directory of the project relative to the cli directory (cli-config.yaml project-dir overrides default value)")
	RootCmd.PersistentFlags().BoolP("no-unix", "w", false, "Execute bash commands as if on Windows (use if your os doesn't support Unix PTYs)")
	RootCmd.PersistentFlags().BoolP("reference", "r", false, "Displays the equivalent shell commands for manual usage")
	viper.BindPFlag("project-dir", RootCmd.PersistentFlags().Lookup("project-dir"))
}

func ExecBashCmd(dir string, name string, args ...string) string {
	runtime_os := runtime.GOOS
	if RootCmd.Flag("no-unix").Value.String() == "true" {
		runtime_os = "windows"
	}
	return pkg.ExecBashCmd(runtime_os, dir, name, args...)
}
