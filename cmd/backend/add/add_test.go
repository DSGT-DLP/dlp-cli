package add

import (
    "bytes"
    "testing"

    "github.com/stretchr/testify/assert"
    "github.com/spf13/cobra"
)

type testingTAdapter struct {
    t *testing.T
}

func (ta *testingTAdapter) Errorf(format string, args ...interface{}) {
    ta.t.Errorf(format, args...)
}

func TestAddCmd(t *testing.T) {
    // Create a buffer to capture the command's output
    var outputBuffer bytes.Buffer

    // Create a Cobra command to represent AddCmd
    cmd := &cobra.Command{}

    // Set the output of the AddCmd to our buffer
    cmd.SetOut(&outputBuffer)

    // Create a TestingT adapter
    tt := &testingTAdapter{t}

    // Create a new command for testing with the same flags
    testCmd := &cobra.Command{}
    testCmd.Flags().Bool("mamba", true, "Add package via mamba (used mainly if cross-platform compatibility is needed)")
    testCmd.Flags().String("channel", "conda-forge", "Specify conda channel to install from (only used if --mamba flag is set)")
    testCmd.Flags().String("env-name", "dlp", "Name of the environment")
    testCmd.Flags().Bool("dev", true, "Add package as dev dependency")

    //testCmd.setArgs("pip-install-test")
    // Run the AddCmd with testCmd as the parent command
    AddCmd.Run(testCmd, []string{"pip-install-test"})

    // Check if the expected message is present in the output
    output := outputBuffer.String()
    expectedMessage := "IMPORTANT"
    assert.Contains(tt, output, expectedMessage)
}
