package add

import (
	"testing"
    "fmt"
    "os"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
    "os/exec"
	"strings"
)

func TestAddCmd(t *testing.T) {
	assert := assert.New(t)

	// Check if AddCmd is *cobra.Command type
	assert.IsType(&cobra.Command{}, AddCmd)

	// Check the command use
	assert.Equal("add {package}", AddCmd.Use)

	// Check the command short description
	assert.Equal("Add package to conda environment, defaults to installation via poetry unless otherwise specified", AddCmd.Short)
}


func TestAddCmdIntegration(t *testing.T) {
	assert := assert.New(t)
    fmt.Println(os.Getenv("PATH"))
	// Create a new conda environment
	_, err := exec.Command("mamba", "create", "--name", "temp_env", "-y").Output()
	assert.NoError(err)

	// Create a new command for testing
	testCmd := &cobra.Command{
		Use:   "add {package}",
		Short: "Add package to conda environment, defaults to installation via poetry unless otherwise specified",
		Long:  `Add package to conda environment from /training, defaults to installation via poetry unless otherwise specified`,
		Args:  cobra.ExactArgs(1),
		Run:   AddCmd.Run,  // Use the same Run function as AddCmd
	}

	// Set the flags on testCmd
	testCmd.Flags().Bool("mamba", true, "")
	testCmd.Flags().String("channel", "conda-forge", "")
	testCmd.Flags().String("env-name", "temp_env", "")
	testCmd.Flags().Bool("dev", false, "")

	// Run the testCmd with the arguments
	testCmd.Run(testCmd, []string{"numpy"})

	// Check if the package was correctly added
	out, err := exec.Command("mamba", "list", "-n", "temp_env", "numpy").Output()
	assert.NoError(err)
	assert.True(strings.Contains(string(out), "numpy"))

	// Delete the temporary conda environment
	_, err = exec.Command("mamba", "env", "remove", "-n", "temp_env", "-y").Output()
	assert.NoError(err)
}